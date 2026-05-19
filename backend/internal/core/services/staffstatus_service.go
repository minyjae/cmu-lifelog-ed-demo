package services

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	repoPort "github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/repositories"
)

type staffStatusService struct {
	repoSS repoPort.StaffStatusRepository
	repoLQ repoPort.ListQueueRepository
}

func NewStaffStatusServiceImpl(s repoPort.StaffStatusRepository, q repoPort.ListQueueRepository) *staffStatusService {
	return &staffStatusService{repoSS: s, repoLQ: q}
}

func (s *staffStatusService) CreateStaffStatus(status *entities.StaffStatus) (*entities.StaffStatus, error) {
	status, err := s.repoSS.Create(status)
	if err != nil {
		return nil, err
	}

	// เพิ่ม create Mapping default = 1
	return status, nil
}

func (s *staffStatusService) GetStaffStatus() (*[]entities.StaffStatus, error) {
	ss, err := s.repoSS.FindAll()
	if err != nil {
		return nil, err
	}

	return ss, nil
}

func (s *staffStatusService) GetStaffStatusByID(id uint) (*entities.StaffStatus, error) {
	ss, err := s.repoSS.FindByID(id)
	if err != nil {
		return nil, err
	}

	return ss, nil
}

func (s *staffStatusService) UpdateStaffStatusName(id uint, name string) (*entities.StaffStatus, error) {
	ss, err := s.repoSS.FindByID(id)
	if err != nil {
		return nil, err
	}
	ss.Status = name
	updatedSS, err := s.repoSS.Save(ss)
	if err != nil {
		return nil, err
	}
	return updatedSS, nil
}

func (s *staffStatusService) RemoveStaffStatus(id uint) error {
	// 1. เปลี่ยนค่า staff_status_id ใน list_queue ก่อน
	if err := s.repoLQ.ChangeStaffStatusToNone(id); err != nil {
		return err
	}

	ss, err := s.repoSS.FindByIDToSoftDelete(id)
	if err != nil {
		return err
	}

	ss.IsActive = false

	if _, err := s.repoSS.Save(ss); err != nil {
		return err
	}

	return nil
}
