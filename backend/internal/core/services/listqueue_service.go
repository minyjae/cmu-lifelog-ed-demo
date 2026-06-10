package services

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"github.com/redis/rueidis"
)

// คีย์ของ cache สำหรับ list view ต่าง ๆ — ทุก key ขึ้นต้นด้วย cacheKeyListPrefix
// เพื่อให้ invalidate ทีเดียวได้ทั้งหมดด้วย SCAN
const (
	cacheKeyListPrefix       = "list_queue:"
	cacheKeyListQueueAll     = "list_queue:all"
	cacheKeyListNotYet       = "list_queue:not_yet"
	cacheKeyListStaffStatus  = "list_queue:staff_status"
	cacheKeyListCourseStatus = "list_queue:course_status"
	cacheKeyListFaculty      = "list_queue:faculty"
	cacheKeyListOwner        = "list_queue:owner"
)

type listQueueService struct {
	repo            repoPort.ListQueueRepository
	orderMapRepo    repoPort.OrderMappingRepository
	facultyRepo     repoPort.FacultyRepository
	staffStatusRepo repoPort.StaffStatusRepository
	cache           *utils.Cache
}

func NewListQueueServiceImpl(r repoPort.ListQueueRepository, om repoPort.OrderMappingRepository, f repoPort.FacultyRepository, ss repoPort.StaffStatusRepository, redis rueidis.Client) *listQueueService {
	return &listQueueService{repo: r, orderMapRepo: om, facultyRepo: f, staffStatusRepo: ss, cache: utils.NewCache(redis)}
}

// invalidateListCache ลบ cache ของ list view ทั้งหมด เมื่อข้อมูลใน list_queue เปลี่ยน
// ลบทุก key ใต้ prefix เพราะทุก view (all/not_yet/faculty/owner/status) ดึงจากตารางเดียวกัน
func (s *listQueueService) invalidateListCache() {
	s.cache.InvalidatePrefix(context.Background(), cacheKeyListPrefix)
}

// getOrLoadList: wrapper เฉพาะของ listQueue — ใช้ cache-aside กลางจาก utils
// แล้ว decorate DateLeft ให้สดทุกครั้งก่อนคืน (ทำกับทั้ง cache hit และผลจาก DB)
func (s *listQueueService) getOrLoadList(key string, fetch func() (*[]entities.ListQueue, error)) (*[]entities.ListQueue, error) {
	q, err := utils.GetOrLoad(context.Background(), s.cache, key, fetch)
	if err != nil {
		return nil, err
	}

	for i := range *q {
		s.DecorateDateLeft(&(*q)[i])
	}
	return q, nil
}

// สร้าง ListQueue ขึ้นมา
func (s *listQueueService) CreateListQueue(req *entities.ListQueue) (*entities.ListQueue, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	req.DurationOfOperationFromTheMemo = s.DaysBetweenCeil(req.DateInfoSubmit, req.OnWeb, loc)
	req.DurationOfWorkFromWordFile = s.DaysBetweenCeil(req.DateWordFileSubmit, req.OnWeb, loc)
	now := time.Now().In(loc)
	req.DateLeft = s.DaysBetweenCeil(now, req.OnWeb, loc)

	q, err := s.repo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create queue: %w", err)
	}

	// Updated
	defaultOrderIDs := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mappings := make([]entities.OrderMapping, 0, len(defaultOrderIDs))
	for _, oid := range defaultOrderIDs {
		mappings = append(mappings, entities.OrderMapping{
			ListQueueID: q.ID,
			OrderID:     oid,
			Checked:     false,
		})
	}
	if err := s.orderMapRepo.Create(&mappings); err != nil {
		return nil, fmt.Errorf("seed order mappings: %w", err)
	}

	r, err := s.repo.FindByIDWithRelations(q.ID)
	if err != nil {
		return nil, err
	}

	s.invalidateListCache()
	return r, nil
}

// Staff : เอาไว้เป็น log ของการทำงานทั้งหมด ให้ staff ได้ดู
func (s *listQueueService) GetListQueue() (*[]entities.ListQueue, error) {
	return s.getOrLoadList(cacheKeyListQueueAll, s.repo.FindAllWithRelations)
}

// Staff : เอาไว้โชว์ ListQueue ที่ยังไม่เสร็จ ไว้เป็นหน้า default สำหรับ staff มาหน้าแรก
func (s *listQueueService) GetListQueueNotYet() (*[]entities.ListQueue, error) {
	return s.getOrLoadList(cacheKeyListNotYet, s.repo.FindNotYetWithRelation)
}

// ยังไม่รู้ว่าจะเอาไปทำอะไร
func (s *listQueueService) GetListQueueByID(id uint) (*entities.ListQueue, error) {
	q, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	s.DecorateDateLeft(q)

	return q, nil
}

// Staff : ให้ staff เลือก status ที่ต้องการดูได้
func (s *listQueueService) GetListQueueByStaffStatus(ids []uint) (*[]entities.ListQueue, error) {
	key := utils.CacheKeyForIDs(cacheKeyListStaffStatus, ids)
	return s.getOrLoadList(key, func() (*[]entities.ListQueue, error) {
		return s.repo.FindByStaffStatusWithRelation(ids)
	})
}

// User : จะคืน listqueue ที่เป็น faculty นั้น
func (s *listQueueService) GetListQueueByFaculty(f string, ids []uint) (*[]entities.ListQueue, error) {
	faculty, err := s.facultyRepo.FindByArg(f)
	if err != nil {
		return nil, err
	}

	key := utils.CacheKeyForIDs(fmt.Sprintf("%s:%d", cacheKeyListFaculty, faculty.ID), ids)
	return s.getOrLoadList(key, func() (*[]entities.ListQueue, error) {
		return s.repo.FindByFacultyWithRelation(faculty.ID, ids)
	})
}

func (s *listQueueService) GetListQueueByCourseStatus(ids []uint) (*[]entities.ListQueue, error) {
	key := utils.CacheKeyForIDs(cacheKeyListCourseStatus, ids)
	return s.getOrLoadList(key, func() (*[]entities.ListQueue, error) {
		return s.repo.FindByCourseStatusWithRelation(ids)
	})
}

func (s *listQueueService) GetListQueueByOwner(email string, ids []uint) (*[]entities.ListQueue, error) {
	key := utils.CacheKeyForIDs(cacheKeyListOwner+":"+email, ids)
	return s.getOrLoadList(key, func() (*[]entities.ListQueue, error) {
		return s.repo.FindByOwnerEmailWithRelation(email, ids)
	})
}

// อัพเดท ListQueue นั้นๆโดยรับข้อมูลมา แล้ว update ทับโดยรับ priority มา **เราจะแยก update ระหว่างข้อมูลใน ListQueue และ Status
func (s *listQueueService) UpdateListQueue(q *entities.ListQueue) (*entities.ListQueue, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	// 1) ตรวจว่ามีการเปลี่ยนสำคัญไหม
	changed, err := s.repo.HasSignificantChanges(q, q.ID, loc)
	if err != nil {
		return nil, fmt.Errorf("check if date changed: %w", err)
	}

	// 2) ถ้ามีการเปลี่ยนแปลงวันที่ → คำนวณใหม่
	updates := map[string]interface{}{
		"title":                   q.Title,
		"staff_id":                q.StaffID,
		"staff_status_id":         q.StaffStatusID,
		"faculty_id":              q.FacultyID,
		"course_status_id":        q.CourseStatusID,
		"date_register":           q.DateRegister,
		"date_info_submit":        q.DateInfoSubmit,
		"date_info_submit14_days": q.DateInfoSubmit14Days,
		"date_word_file_submit":   q.DateWordFileSubmit,
		"on_web":                  q.OnWeb,
		"note":                    q.Note,
		"owner":                   q.Owner,
		// owner ถ้าแก้ไขด้วย ให้ใส่ด้วยตามชนิดที่ใช้ (pq.StringArray เป็นต้น)
	}

	if changed {
		updates["duration_of_operation_from_the_memo"] = s.DaysBetweenCeil(q.DateInfoSubmit, q.OnWeb, loc)
		updates["duration_of_work_from_word_file"] = s.DaysBetweenCeil(q.DateWordFileSubmit, q.OnWeb, loc)
		// DateLeft = now → onWeb (dynamic ณ ตอนอัปเดต)
		updates["date_left"] = s.DaysBetweenCeil(time.Now(), q.OnWeb, loc)
		s.UpdatePriority(q.ID, 1)
	}

	// 1) เซฟฟิลด์ทั้งหมดก่อน (รวม status_id)
	if err := s.repo.UpdateFields(q.ID, updates); err != nil {
		return nil, fmt.Errorf("update queue: %w", err)
	}

	// 2) โหลดกลับพร้อม relations เพื่อตรวจ Type จริง
	r, err := s.repo.FindByIDWithRelations(q.ID)
	if err != nil {
		return nil, err
	}

	// 3) ถ้า Done/Cancel → ตั้ง priority = 0 และ shift รายการด้านล่างขึ้น
	isDone := r.StaffStatus.Type == "Done" || r.CourseStatus.Type == "Done"

	if isDone && r.Priority != 0 {
		old := r.Priority
		if err := s.repo.UpdateFields(r.ID, map[string]interface{}{"priority": 0}); err != nil {
			return nil, fmt.Errorf("set priority zero: %w", err)
		}
		if err := s.repo.ShiftDownFrom(old); err != nil {
			return nil, fmt.Errorf("shift down after done: %w", err)
		}
		// โหลดซ้ำ (optional)
		r, _ = s.repo.FindByIDWithRelations(q.ID)
	}

	s.invalidateListCache()
	return r, nil
}

func (s *listQueueService) UpdatePriority(id uint, newPriority uint) (*entities.ListQueue, error) {
	q, err := s.repo.FindByIDWithRelations(id)
	if err != nil {
		return nil, err
	}

	// กันคิว Done/Cancel และคิว priority=0
	if q.Priority == 0 || q.StaffStatus.Type == "Done" || q.StaffStatus.Type == "Cancel" {
		return nil, fmt.Errorf("failed to update priority of listqueue: this listqueue is completed or cancel")
	}

	oldPriority := q.Priority
	if oldPriority == newPriority {
		return q, nil
	}

	if err := s.repo.ShiftPriorityAndMove(id, oldPriority, newPriority); err != nil {
		return nil, err
	}

	// บันทึกค่าใหม่ (เผื่อ fields อื่นใน entity)
	q.Priority = newPriority
	if err := s.repo.Save(q); err != nil {
		return nil, fmt.Errorf("failed to update priority of listqueue: %w", err)
	}

	r, err := s.repo.FindByIDWithRelations(q.ID)
	if err != nil {
		return nil, err
	}

	s.invalidateListCache()
	return r, nil
}

// อัพเดท StaffStatus โดยที่จะ update UserStatus ให้ด้วยให้สอดคล้องกัน และเมื่อ userStatus map ตรงกับ staffStatus == All done จะ set priority = 0 ทันที
func (s *listQueueService) UpdateStaffStatus(queueID, staffStatusID uint) (*entities.ListQueue, error) {
	q, err := s.repo.FindByID(queueID)
	if err != nil {
		return nil, err
	}

	staffStatus, err := s.staffStatusRepo.FindByID(staffStatusID)
	if err != nil {
		return nil, err
	}

	q.StaffStatusID = staffStatusID
	q.StaffStatus = *staffStatus

	if staffStatus.Type == "Done" || staffStatus.Type == "Cancel" {
		oldPriority := q.Priority
		q.Priority = 0
		if err := s.repo.ShiftDownFrom(oldPriority); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Save(q); err != nil {
		return nil, err
	}

	r, err := s.repo.FindByIDWithRelations(q.ID)
	if err != nil {
		return nil, err
	}

	s.invalidateListCache()
	return r, nil
}

func (s *listQueueService) RemoveListQueueForDev(id uint) error {
	l, err := s.orderMapRepo.FindByListQueueID(id)
	if err != nil {
		return err
	}

	err = s.orderMapRepo.DeleteOrder(l)
	if err != nil {
		return err
	}

	err = s.repo.DeleteByID(id)
	if err != nil {
		return err
	}

	s.invalidateListCache()
	return nil
}

func (s *listQueueService) DaysBetweenCeil(a, b time.Time, loc *time.Location) int {
	// นอร์มัลไลซ์เป็น “เที่ยงคืน” ของโซนเวลาเดียวกันก่อน
	a = a.In(loc)
	b = b.In(loc)
	a = time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, loc)
	b = time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, loc)

	d := b.Sub(a)
	days := int(math.Ceil(d.Hours() / 24.0))

	// ไม่ให้ติดลบ (ถ้าต้องการ)
	if days < 0 {
		days = 0
	}
	return days
}

func (s *listQueueService) DecorateDateLeft(lq *entities.ListQueue) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	lq.DateLeft = s.DaysBetweenCeil(time.Now(), lq.OnWeb, loc)
}
