package entities

import (
	"time"

	"github.com/lib/pq"
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
	FacultyOfPoliticalScience                      FacultyName = "คณะรัฐศาสตร์"
	CollegeOfArtMediaAndTechnology                 FacultyName = "วิทยาลัยศิลปะ สื่อ และเทคโนโลยี"
	FacultyOfPublicHealth                          FacultyName = "คณะสาธารณสุขศาสตร์"
	CollegeOfEducationAndSeaManagement             FacultyName = "วิทยาลัยการศึกษาและการจัดการทะเล"
	InternationalCollegeOfDigitalTechnology        FacultyName = "วิทยาลัยนานาชาติเทคโนโลยีดิจิทัล"
	InstituteOfPublic                              FacultyName = "สถาบันนโยบายสาธารณะ"
	BiomedicalEngineeringInstitute                 FacultyName = "สถาบันวิศวกรรมชีวการแพทย์"
	AcademicServiceCenter                          FacultyName = "สำนักบริการวิชาการ"
	OfficeOfEducationQualityDevelopment            FacultyName = "สำนักงานพัฒนาคุณภาพการศึกษา"
	OfficeOfTheUniversity                          FacultyName = "สำนักงานมหาวิทยาลัย"
	LannaRiceResearchCenter                        FacultyName = "ศูนย์วิจัยข้าวล้านนา"
	LanguageInstitute                              FacultyName = "สถาบันภาษา"
	SchoolOfLifeLongEducation                      FacultyName = "วิทยาลัยการศึกษาตลอดชีวิต"
	CenterOfSafetyOccupationalHealthAndEnvironment FacultyName = "ศูนย์บริหารจัดการความปลอดภัยฯ SHE"
	TeachingAndLearningInnovationCenter            FacultyName = "ศูนย์นวัตกรรมการสอนและการเรียนรู้ TLIC"
)

type ListQueue struct {
	ID                             uint           `json:"id"`
	Priority                       uint           `json:"priority"`
	Title                          string         `json:"title"`
	StaffID                        uint           `json:"staff_id"`
	Staff                          Users          `json:"staff"`
	FacultyID                      uint           `json:"faculty_id"`
	Faculty                        Faculty        `json:"faculty"`
	StaffStatusID                  uint           `json:"staff_status_id"`
	StaffStatus                    StaffStatus    `json:"staff_status"`
	DateWordFileSubmit             time.Time      `json:"wordfile_submit"`
	DateInfoSubmit                 time.Time      `json:"info_submit"`
	DateInfoSubmit14Days           time.Time      `json:"info_submit_14days"`
	DateRegister                   time.Time      `json:"time_register"`
	DateLeft                       int            `json:"date_left"` // ใช้ pointer เพื่อรองรับ null
	OnWeb                          time.Time      `json:"on_web"`
	AppointmentDateAW              time.Time      `json:"appointment_date_aw"`
	OrderMappings                  []OrderMapping `json:"order_mappings"`
	CourseStatusID                 uint           `json:"course_status_id"`
	CourseStatus                   CourseStatus   `json:"course_status"`
	DurationOfOperationFromTheMemo int            `json:"duration_of_operation_from_the_memo"`
	DurationOfWorkFromWordFile     int            `json:"duration_of_work_from_word_file"`
	Note                           string         `json:"note"`
	Owner                          pq.StringArray `json:"owner"`
	CreatedAt                      time.Time      `json:"created_at"`
	UpdatedAt                      time.Time      `json:"updated_at"`
}

type Faculty struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	NameTH    string    `json:"nameTH"`
	NameEN    string    `json:"nameEN"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (f FacultyName) IsValid() bool {
// 	switch f {
// 	case
// 		FacultyOfEducation,
// 		FacultyOfHumanities,
// 		FacultyOfFineArts,
// 		FacultyOfSocialSciences,
// 		FacultyOfScience,
// 		FacultyOfEngineering,
// 		FacultyOfMedicine,
// 		FacultyOfAgriculture,
// 		FacultyOfDentistry,
// 		FacultyOfPharmacy,
// 		FacultyOfMedTechnology,
// 		FacultyOfNursing,
// 		FacultyOfAgroIndustry,
// 		FacultyOfVeterinaryMedicine,
// 		FacultyOfArchitecture,
// 		BachelorOfBussinessAdministration,
// 		FacultyOfEconomics,
// 		FacultyOfLaw,
// 		FacultyOfMassCommunication,
// 		FacultyOfPoliticalScience,
// 		CollegeOfArtMediaAndTechnology,
// 		FacultyOfPublicHealth,
// 		CollegeOfEducationAndSeaManagement,
// 		InternationalCollegeOfDigitalTechnology,
// 		InstituteOfPublic,
// 		BiomedicalEngineeringInstitute,
// 		AcademicServiceCenter,
// 		OfficeOfEducationQualityDevelopment,
// 		OfficeOfTheUniversity,
// 		LannaRiceResearchCenter,
// 		LanguageInstitute,
// 		SchoolOfLifeLongEducation,
// 		CenterOfSafetyOccupationalHealthAndEnvironment,
// 		TeachingAndLearningInnovationCenter:
// 		return true
// 	default:
// 		return false
// 	}
// }
