package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type FacultyName string

const (
	FacultyOfEducation                             FacultyName = "คณะศึกษาศาสตร์"
	FacultyOfHumanities                            FacultyName = "คณะมนุษยศาสตร์"
	FacultyOfFineArts                              FacultyName = "คณะวิจิตรศิลป์"
	FacultyOfSocialSciences                        FacultyName = "คณะสังคมศาสตร์"
	FacultyOfScience                               FacultyName = "คณะวิทยาศาสตร์"
	FacultyOfEngineering                           FacultyName = "คณะวิศวกรรมศาสตร์"
	FacultyOfMedicine                              FacultyName = "คณะแพทยศาสตร์"
	FacultyOfAgriculture                           FacultyName = "คณะเกษตรศาสตร์"
	FacultyOfDentistry                             FacultyName = "คณะทันตแพทยศาสตร์"
	FacultyOfPharmacy                              FacultyName = "คณะเภสัชศาสตร์"
	FacultyOfMedTechnology                         FacultyName = "คณะเทคนิคการแพทย์"
	FacultyOfNursing                               FacultyName = "คณะพยาบาลศาสตร์"
	FacultyOfAgroIndustry                          FacultyName = "คณะอุตสาหกรรมเกษตร"
	FacultyOfVeterinaryMedicine                    FacultyName = "คณะสัตวแพทยศาสตร์"
	FacultyOfArchitecture                          FacultyName = "คณะสถาปัตยกรรมศาสตร์"
	BachelorOfBussinessAdministration              FacultyName = "คณะบริหารธุรกิจ"
	FacultyOfEconomics                             FacultyName = "คณะเศรษฐศาสตร์"
	FacultyOfLaw                                   FacultyName = "คณะนิติศาสตร์"
	FacultyOfMassCommunication                     FacultyName = "คณะสื่อสารมวลชน"
	FacultyOfPoliticalScience                      FacultyName = "คณะรัฐศาสตร์และรัฐประศาสนศาสตร์"
	CollegeOfArtMediaAndTechnology                 FacultyName = "วิทยาลัยศิลปะ สื่อ และเทคโนโลยี"
	FacultyOfPublicHealth                          FacultyName = "คณะสาธารณสุขศาสตร์"
	CollegeOfEducationAndSeaManagement             FacultyName = "วิทยาลัยการศึกษาและการจัดการทะเล" // ไม่มี
	InternationalCollegeOfDigitalTechnology        FacultyName = "วิทยาลัยนานาชาตินวัตกรรมดิจิทัล"
	InstituteOfPublic                              FacultyName = "สถาบันนโยบายสาธารณะ"
	BiomedicalEngineeringInstitute                 FacultyName = "สถาบันวิศวกรรมชีวการแพทย์"
	AcademicServiceCenter                          FacultyName = "สำนักบริการวิชาการ"                     // ไม่มี
	OfficeOfEducationQualityDevelopment            FacultyName = "สำนักงานพัฒนาคุณภาพการศึกษา"            // ไม่มี
	OfficeOfTheUniversity                          FacultyName = "สำนักงานมหาวิทยาลัย"                    // ไม่มี
	LannaRiceResearchCenter                        FacultyName = "ศูนย์วิจัยข้าวล้านนา"                   // ไม่มี
	LanguageInstitute                              FacultyName = "สถาบันภาษา"                             // ไม่มี
	SchoolOfLifeLongEducation                      FacultyName = "วิทยาลัยการศึกษาตลอดชีวิต"              // ไม่มี
	CenterOfSafetyOccupationalHealthAndEnvironment FacultyName = "ศูนย์บริหารจัดการความปลอดภัยฯ SHE"      // ไม่มี
	TeachingAndLearningInnovationCenter            FacultyName = "ศูนย์นวัตกรรมการสอนและการเรียนรู้ TLIC" // ไม่มี
)

// ListQueue table
type ListQueue struct {
	ID                             uint           `gorm:"primaryKey" json:"id"`
	Priority                       uint           `gorm:"not null" json:"priority"`
	Title                          string         `gorm:"not null" json:"title"`
	StaffID                        uint           `gorm:"not null;index" json:"staff_id"`
	Staff                          Users          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"staff"`
	FacultyID                      uint           `gorm:"not null" json:"faculty_id"`
	Faculty                        Faculty        `gorm:"not null" json:"faculty"`
	StaffStatusID                  uint           `gorm:"not null;index" json:"staff_status_id"`
	StaffStatus                    StaffStatus    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"staff_status"`
	DateWordFileSubmit             time.Time      `json:"wordfile_submit"`
	DateInfoSubmit                 time.Time      `json:"info_submit"`
	DateInfoSubmit14Days           time.Time      `json:"info_submit_14days"`
	DateRegister                   time.Time      `json:"time_register"`
	DateLeft                       int            `json:"date_left"` // ใช้ pointer เพื่อรองรับ null
	OnWeb                          time.Time      `json:"on_web"`
	AppointmentDateAW              time.Time      `json:"appointment_date_aw"`
	OrderMappings                  []OrderMapping `gorm:"foreignKey:ListQueueID" json:"order_mappings"`
	CourseStatusID                 uint           `gorm:"not null;index" json:"course_status_id"`
	CourseStatus                   CourseStatus   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"course_status"`
	DurationOfOperationFromTheMemo int            `json:"duration_of_operation_from_the_memo"`
	DurationOfWorkFromWordFile     int            `json:"duration_of_work_from_word_file"`
	Note                           string         `json:"note"`
	Owner                          pq.StringArray `gorm:"type:text[];not null;default:'{}'" json:"owner"`
	CreatedAt                      time.Time      `json:"created_at"`
	UpdatedAt                      time.Time      `json:"updated_at"`
}

// แปลงโมเดล ListQueue เป็น entity
func (l *ListQueue) ToEntity() *entities.ListQueue {
	orderMappings := make([]entities.OrderMapping, len(l.OrderMappings))
	for i, om := range l.OrderMappings {
		orderMappings[i] = *om.ToEntity()
	}
	return &entities.ListQueue{
		ID:                             l.ID,
		Priority:                       l.Priority,
		Title:                          l.Title,
		StaffID:                        l.StaffID,
		Staff:                          *l.Staff.ToEntity(),
		FacultyID:                      l.FacultyID,
		Faculty:                        *l.Faculty.ToEntity(),
		StaffStatusID:                  l.StaffStatusID,
		StaffStatus:                    *l.StaffStatus.ToEntity(),
		DateWordFileSubmit:             l.DateWordFileSubmit,
		DateInfoSubmit:                 l.DateInfoSubmit,
		DateInfoSubmit14Days:           l.DateInfoSubmit14Days,
		DateRegister:                   l.DateRegister,
		DateLeft:                       l.DateLeft,
		OnWeb:                          l.OnWeb,
		AppointmentDateAW:              l.AppointmentDateAW,
		OrderMappings:                  orderMappings,
		CourseStatusID:                 l.CourseStatusID,
		CourseStatus:                   *l.CourseStatus.ToEntity(),
		DurationOfOperationFromTheMemo: l.DurationOfOperationFromTheMemo,
		DurationOfWorkFromWordFile:     l.DurationOfWorkFromWordFile,
		Note:                           l.Note,
		Owner:                          []string(l.Owner),
		CreatedAt:                      l.CreatedAt,
		UpdatedAt:                      l.UpdatedAt,
	}
}

// แปลง Entity เป็นโมเดล ListQueue
func (l *ListQueue) FromEntity(entity *entities.ListQueue) {
	l.ID = entity.ID
	l.Priority = entity.Priority
	l.Title = entity.Title
	l.StaffID = entity.StaffID
	l.Staff = Users{}
	l.FacultyID = entity.FacultyID
	l.Faculty = Faculty{}
	l.StaffStatusID = entity.StaffStatusID
	l.StaffStatus = StaffStatus{}
	l.DateWordFileSubmit = entity.DateWordFileSubmit
	l.DateInfoSubmit = entity.DateInfoSubmit
	l.DateInfoSubmit14Days = entity.DateInfoSubmit14Days
	l.DateRegister = entity.DateRegister
	l.DateLeft = entity.DateLeft
	l.OnWeb = entity.OnWeb
	l.AppointmentDateAW = entity.AppointmentDateAW
	orderMappings := make([]OrderMapping, len(entity.OrderMappings))
	for i, om := range entity.OrderMappings {
		orderMapping := OrderMapping{}
		orderMapping.FromEntity(&om)
		orderMappings[i] = orderMapping
	}
	l.OrderMappings = orderMappings
	l.CourseStatusID = entity.CourseStatusID
	l.CourseStatus = CourseStatus{}
	l.DurationOfOperationFromTheMemo = entity.DurationOfOperationFromTheMemo
	l.DurationOfWorkFromWordFile = entity.DurationOfWorkFromWordFile
	l.Note = entity.Note
	l.Owner = pq.StringArray(entity.Owner)
	l.CreatedAt = entity.CreatedAt
	l.UpdatedAt = entity.UpdatedAt
}
