package repositories

import (
	"fmt"
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"gorm.io/gorm"
)

type listQueueRepositoryImpl struct {
	db *gorm.DB
}

func NewListQueueRepository(db *gorm.DB) *listQueueRepositoryImpl {
	return &listQueueRepositoryImpl{db: db}
}

func (s *listQueueRepositoryImpl) Create(q *entities.ListQueue) (*entities.ListQueue, error) {
	list := models.ListQueue{}
	list.FromEntity(q)

	var count int64
	if err := s.db.Model(&models.ListQueue{}).
		Joins("JOIN staff_statuses ON list_queues.staff_status_id = staff_statuses.id").
		Where("staff_statuses.type NOT IN ?", []string{"Done", "Cancel"}).
		Count(&count).Error; err != nil {
		return nil, err
	}

	list.Priority = uint(count) + 1

	if err := s.db.Create(&list).Error; err != nil {
		return nil, err
	}

	result := list.ToEntity()
	return result, nil
}

func (s *listQueueRepositoryImpl) Save(q *entities.ListQueue) error {
	list := models.ListQueue{}
	list.FromEntity(q)

	err := s.db.Save(&list).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *listQueueRepositoryImpl) Update(q *entities.ListQueue) (*entities.ListQueue, error) {
	list := models.ListQueue{}
	list.FromEntity(q)

	if err := s.db.Model(&list).
		Where("id = ?", q.ID).
		Updates(q).Error; err != nil {
		return nil, err
	}

	return list.ToEntity(), nil
}

func (s *listQueueRepositoryImpl) FindByID(id uint) (*entities.ListQueue, error) {
	list := models.ListQueue{}

	if err := s.db.Where("id = ?", id).
		First(&list).Error; err != nil {
		return nil, err
	}

	result := list.ToEntity()

	return result, nil
}

func (s *listQueueRepositoryImpl) FindByIDWithRelations(id uint) (*entities.ListQueue, error) {
	list := models.ListQueue{}
	if err := s.db.Preload("StaffStatus").
		Preload("CourseStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("Faculty").
		First(&list, id).Error; err != nil {
		return nil, err
	}

	result := list.ToEntity()
	return result, nil
}

func (s *listQueueRepositoryImpl) FindAllWithRelations() (*[]entities.ListQueue, error) {
	var lists []models.ListQueue
	if err := s.db.
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Order("priority ASC").
		Find(&lists).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(lists))
	for i, l := range lists {
		result[i] = *l.ToEntity()
	}
	return &result, nil
}

func (s *listQueueRepositoryImpl) FindAllWithFacultyAndRelations() (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}
	if err := s.db.Order("priority ASC").
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *listQueueRepositoryImpl) FindNotYetWithRelation() (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	if err := s.db.Where("priority != ?", 0).
		Order("priority ASC").
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *listQueueRepositoryImpl) FindByStaffStatusWithRelation(ids []uint) (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	if err := s.db.Where("staff_status_id IN ?", ids).
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Order("priority ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *listQueueRepositoryImpl) FindStaffStatusInListQueue() (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	if err := s.db.Joins("JOIN staff_statuses ON list_queues.staff_status_id = staff_statuses.id").
		Where("staff_statuses.type NOT IN ?", []string{"Done", "Cancel"}).
		Find(&list).Error; err != nil {
		return nil, err
	}

	fmt.Println("list model when where user status in 1,2", list)

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *listQueueRepositoryImpl) FindByFacultyWithRelation(facultyID uint, courseIDs []uint) (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	// เริ่มด้วย base query
	q := s.db.Model(&models.ListQueue{}).
		Where("faculty_id = ?", facultyID)

	// มี courseIDs ค่อยเติมเงื่อนไข IN
	if len(courseIDs) > 0 {
		q = q.Where("course_status_id IN ?", courseIDs)
	}

	if err := q.
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Order("priority ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}
	return &result, nil
}

func (s *listQueueRepositoryImpl) FindByCourseStatusWithRelation(ids []uint) (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	if err := s.db.Where("course_status_id IN ?", ids).
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Order("priority ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}

	return &result, nil
}

func (s *listQueueRepositoryImpl) FindByOwnerEmailWithRelation(email string, ids []uint) (*[]entities.ListQueue, error) {
	list := []models.ListQueue{}

	// base query: owner เป็น text[] (Postgres) ใช้ ? = ANY(owner)
	q := s.db.Model(&models.ListQueue{}).
		Where("? = ANY(owner)", email).
		Preload("StaffStatus").
		Preload("OrderMappings").
		Preload("OrderMappings.Order").
		Preload("Staff").
		Preload("CourseStatus").
		Preload("Faculty").
		Order("priority ASC")

	// มี course_status_id กรองเพิ่ม
	if len(ids) > 0 {
		q = q.Where("course_status_id IN ?", ids)
	}

	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.ListQueue, len(list))
	for i, m := range list {
		result[i] = *m.ToEntity()
	}
	return &result, nil
}

func (s *listQueueRepositoryImpl) DeleteByID(id uint) error {
	list := models.ListQueue{}

	if err := s.db.Where("id = ?", id).
		Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

func (s *listQueueRepositoryImpl) ChangeStaffStatusToNone(statusID uint) error {
	if err := s.db.Model(&models.ListQueue{}).Where("staff_status_id = ?", statusID).Update("staff_status_id", 1).Error; err != nil {
		return err
	}

	return nil
}

func (s *listQueueRepositoryImpl) HasSignificantChanges(updated *entities.ListQueue, id uint, loc *time.Location) (bool, error) {
	var original models.ListQueue

	// เลือกเฉพาะฟิลด์ที่ต้องใช้เปรียบเทียบ
	if err := s.db.
		Select("id", "priority", "title", "staff_id", "staff_status_id", "faculty_id", "course_status_id",
			"date_register", "date_info_submit", "date_info_submit14_days", "date_word_file_submit", "on_web").
		First(&original, id).Error; err != nil {
		return false, fmt.Errorf("find list by id: %w", err)
	}

	// เทียบฟิลด์เวลาแบบ normalize วัน (ลดปัญหา TZ)
	if !utils.EqualDate(original.DateInfoSubmit, updated.DateInfoSubmit, loc) ||
		!utils.EqualDate(original.DateWordFileSubmit, updated.DateWordFileSubmit, loc) ||
		!utils.EqualDate(original.OnWeb, updated.OnWeb, loc) ||
		!utils.EqualDate(original.DateRegister, updated.DateRegister, loc) ||
		!utils.EqualDate(original.DateInfoSubmit14Days, updated.DateInfoSubmit14Days, loc) {
		return true, nil
	}

	// เทียบฟิลด์ non-time
	if original.Title != updated.Title ||
		original.StaffID != updated.StaffID ||
		original.StaffStatusID != updated.StaffStatusID ||
		original.FacultyID != updated.FacultyID || // ← ใช้ FacultyID แทน Faculty.ID
		original.CourseStatusID != updated.CourseStatusID ||
		original.Priority != updated.Priority {
		return true, nil
	}

	return false, nil
}

func (s *listQueueRepositoryImpl) ShiftDownFrom(from uint) error {
	return s.db.Model(&models.ListQueue{}).
		Where("priority > ? AND priority != 0", from).
		Update("priority", gorm.Expr("priority - 1")).Error
}

func (s *listQueueRepositoryImpl) ShiftPriorityAndMove(id, from, to uint) error {
	// CASE: ย้ายลง (เช่น 7 → 10)
	if from < to {
		if err := s.db.Model(&models.ListQueue{}).
			Where("priority > ? AND priority <= ? AND priority != 0", from, to).
			Update("priority", gorm.Expr("priority - 1")).Error; err != nil {
			return err
		}
	}

	// CASE: ย้ายขึ้น (เช่น 10 → 7)
	if from > to {
		if err := s.db.Model(&models.ListQueue{}).
			Where("priority >= ? AND priority < ? AND priority != 0", to, from).
			Update("priority", gorm.Expr("priority + 1")).Error; err != nil {
			return err
		}
	}

	// Update task ที่ถูกย้าย
	if err := s.db.Model(&models.ListQueue{}).
		Where("id = ?", id).
		Update("priority", to).Error; err != nil {
		return err
	}

	return nil
}

func (s *listQueueRepositoryImpl) UpdateFields(id uint, updates map[string]interface{}) error {
	return s.db.Model(&models.ListQueue{}).
		Where("id = ?", id).
		Updates(updates).Error
}
