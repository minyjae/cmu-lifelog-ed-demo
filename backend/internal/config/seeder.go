package config

import (
	"log"
	"os"

	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/models"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
	"gorm.io/gorm"
)

func CreateSeeds(db *gorm.DB, config *Config) error {

	// ตรวจสอบว่ามี admin user อยู่แล้วหรือไม่
	var count int64
	db.Model(&models.Users{}).Where("email = ?", os.Getenv("ADMIN_EMAIL")).Count(&count)

	if count > 0 {
		log.Println("All seed already exist, skipping seeding")
		return nil
	}

	// ตรวจสอบว่ามีการตั้งค่า admin credentials หรือไม่
	if config.AdminEmail == "" {
		log.Println("⚠️  ADMIN_EMAIL not set, skipping admin user seeding")
		log.Println("💡 To create admin user, set ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_FIRST_NAME, ADMIN_LAST_NAME in .env")
		return nil
	}

	// hash default password สำหรับ seed users
	hashedPassword, err := utils.HashPassword(config.AdminPassword)
	if err != nil {
		return err
	}
	log.Printf("🔑 Seed users default password: %s", config.AdminPassword)

	// สร้าง admin user
	staff := []models.Users{
		{
			Name:               "raweewan.p",
			Email:              "raweewan.p@cmu.ac.th",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "MS",
			PreNameTH:          "นางสาว",
			PreNameEN:          "Miss",
			FirstNameTH:        "ระวีวรรณ",
			FirstNameEN:        "RAWEEWAN",
			LastNameTH:         "พยัคฆชาติ",
			LastNameEN:         "PAYAKKACHAT",
			OrganizationCode:   "76",
			OrganizationNameTH: "วิทยาลัยการศึกษาตลอดชีวิต",
			OrganizationNameEN: "Lifelong Education College",
			ITAccountTypeID:    "MISEmpAcc",
			ITAccountTypeTH:    "บุคลากร",
			ITAccountTypeEN:    "MIS Employee",
			IsFirstTime:        false,
		},
		{
			Name:               "thatthana_sri",
			Email:              "thatthana_sri@cmu.ac.th",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "OTH",
			PreNameTH:          "",
			PreNameEN:          "",
			FirstNameTH:        "ทัตธน",
			FirstNameEN:        "THATTHANA",
			LastNameTH:         "ศรีเงิน",
			LastNameEN:         "SRINGEON",
			OrganizationCode:   "06",
			OrganizationNameTH: "คณะวิศวกรรมศาสตร์",
			OrganizationNameEN: "Faculty of Engineering",
			ITAccountTypeID:    "StdAcc",
			ITAccountTypeTH:    "นักศึกษาปัจจุบัน",
			ITAccountTypeEN:    "Student Account",
			IsFirstTime:        false,
		},
		{
			Name:               "surapa_luangpiwdet",
			Email:              "surapa_luangpiwdet@cmu.ac.th",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "OTH",
			PreNameTH:          "",
			PreNameEN:          "",
			FirstNameTH:        "สุรภา",
			FirstNameEN:        "SURAPA",
			LastNameTH:         "หลวงผิวเดช",
			LastNameEN:         "LUANGPIWDET",
			OrganizationCode:   "06",
			OrganizationNameTH: "คณะวิศวกรรมศาสตร์",
			OrganizationNameEN: "Faculty of Engineering",
			ITAccountTypeID:    "StdAcc",
			ITAccountTypeTH:    "นักศึกษาปัจจุบัน",
			ITAccountTypeEN:    "Student Account",
			IsFirstTime:        false,
		},
		{
			Name:               "nontapan_chanadee",
			Email:              "nontapan_chanadee@cmu.ac.th",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "OTH",
			PreNameTH:          "",
			PreNameEN:          "",
			FirstNameTH:        "นนทพันธุ์",
			FirstNameEN:        "NONTAPAN",
			LastNameTH:         "ชนะดี",
			LastNameEN:         "CHANADEE",
			OrganizationCode:   "06",
			OrganizationNameTH: "คณะวิศวกรรมศาสตร์",
			OrganizationNameEN: "Faculty of Engineering",
			ITAccountTypeID:    "StdAcc",
			ITAccountTypeTH:    "นักศึกษาปัจจุบัน",
			ITAccountTypeEN:    "Student Account",
			IsFirstTime:        false,
		}, {
			Name:               "jiradate_ora",
			Email:              "jiradate_ora@cmu.ac.th",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "OTH",
			PreNameTH:          "",
			PreNameEN:          "",
			FirstNameTH:        "จิรเดช",
			FirstNameEN:        "JIRADATE",
			LastNameTH:         "อรทัย",
			LastNameEN:         "ORATAI",
			OrganizationCode:   "06",
			OrganizationNameTH: "คณะวิศวกรรมศาสตร์",
			OrganizationNameEN: "Faculty of Engineering",
			ITAccountTypeID:    "StdAcc",
			ITAccountTypeTH:    "นักศึกษาปัจจุบัน",
			ITAccountTypeEN:    "Student Account",
			IsFirstTime:        false,
		},
		{
			Name:               "for_testing",
			Email:              "test@gmail.com",
			Password:           hashedPassword,
			Role:               "admin",
			PreNameID:          "MS",
			PreNameTH:          "นางสาว",
			PreNameEN:          "Miss",
			FirstNameTH:        "สาทินี",
			FirstNameEN:        "SATENEE",
			LastNameTH:         "พายัพ",
			LastNameEN:         "PAYUB",
			OrganizationCode:   "76",
			OrganizationNameTH: "วิทยาลัยการศึกษาตลอดชีวิต",
			OrganizationNameEN: "Lifelong Education College",
			ITAccountTypeID:    "MISEmpAcc",
			ITAccountTypeTH:    "บุคลากร",
			ITAccountTypeEN:    "MIS Employee",
			IsFirstTime:        false,
		},
	}

	if err := db.Create(staff).Error; err != nil {
		log.Printf("❌ Error creating admin user: %v", err)
		return err
	}

	staffStatuses := []models.StaffStatus{
		{Status: "ยังไม่ได้อ่าน", Type: "None"},
		{Status: "รอคณะแก้ไข", Type: "Processing"},
		{Status: "แก้ไขเรียบร้อยแล้ว", Type: "Processing"},
		{Status: "กำลังแก้ไข - กำลังอ่าน", Type: "Processing"},
		{Status: "กำลังทำเอกสาร final", Type: "Processing"},
		{Status: "รอ เสนอเซ็น", Type: "Processing"},
		{Status: "เอกสารเสร็จ รอ บันทึกข้อความ", Type: "Processing"},
		{Status: "เอกสารเสร็จ รอ Link ITSC", Type: "Processing"},
		{Status: "รอคณะอนุฯ พิจารณา", Type: "Processing"},
		{Status: "รอส่ง e-Doc", Type: "Processing"},
		{Status: "เอกสารเสร็จแล้ว", Type: "Processing"},
		{Status: "รอส่ง comment อนุฯ ให้คณะ", Type: "Processing"},
		{Status: "รอคณะปรับเอกสาร / ตัดต่อคลิป", Type: "Processing"},
		{Status: "กำลังคีย์ข้อมูล", Type: "Processing"},
		{Status: "รอกดแสดงผลหน้า WEB", Type: "Processing"},
		{Status: "รอเชื่อม Link MOOC & LE", Type: "Processing"},
		{Status: "คีย์ข้อมูลเสร็จแล้ว รอ AW", Type: "Processing"},
		{Status: "เสร็จหมดแล้ว รอ AW", Type: "Processing"},
		{Status: "คณะอนุฯ อนุมัติแล้ว", Type: "Processing"},
		{Status: "All Done", Type: "Done"},
		{Status: "x คณะขอ Cancel x", Type: "Cancel"},
		{Status: "x คณะขอ Pause x", Type: "Cancel"},
	}

	if err := db.Create(&staffStatuses).Error; err != nil {
		panic("failed to create default staff status")
	}

	orders := []models.Order{
		{Title: "โน้ตงาน AW Graphic"},
		{Title: "ได้รับ Artwork"},
		{Title: "ใบปะหน้า & แบบพิจารณา"},
		{Title: "เสนอเซ็น"},
		{Title: "เตรียมเอกสารส่ง e-Doc"},
		{Title: "ส่ง e-Doc เสนอคณะอนุฯ"},
		{Title: "สร้าง LMS & คีย์ลงระบบ"},
		{Title: "ทำใบสมัครผู้เรียน"},
		{Title: "จัดลำดับหลักสูตร"},
		{Title: "กด Display บน WEB"},
	}

	if err := db.Create(&orders).Error; err != nil {
		panic("failed to create default order")
	}

	faculty := []models.Faculty{
		{Code: "01", NameTH: "คณะมนุศย์ศาสตร์", NameEN: "Faculty of Humanities"},
		{Code: "02", NameTH: "คณะศึกษาศาสตร์", NameEN: "Faculty of Education"},
		{Code: "03", NameTH: "คณะวิจิตรศิลป์", NameEN: "Faculty of Fine Art"},
		{Code: "04", NameTH: "คณะสังคมศาสตร์", NameEN: "Faculty of Social Sciences"},
		{Code: "05", NameTH: "คณะวิทยาศาสตร์", NameEN: "Faculty of Science"},
		{Code: "06", NameTH: "คณะวิศวกรรมศาสตร์", NameEN: "Faculty of Engineering"},
		{Code: "07", NameTH: "คณะแพทยศาสตร์", NameEN: "Faculty of Medicine"},
		{Code: "08", NameTH: "คณะเกษตรศาสตร์", NameEN: "Faculty of Agriculture"},
		{Code: "09", NameTH: "คณะทันตแพทยศาสตร์", NameEN: "Faculty of Dentistry"},
		{Code: "10", NameTH: "คณะเภสัชศาสตร์", NameEN: "Faculty of Pharmacy"},
		{Code: "11", NameTH: "คณะเทคนิคการแพทย์", NameEN: "Faculty of Medical Technology"},
		{Code: "12", NameTH: "คณะพยาบาลศาสตร์", NameEN: "Faculty of Nursing"},
		{Code: "13", NameTH: "คณะอุตสาหกรรมเกษตร", NameEN: "Faculty of Agro-Industry"},
		{Code: "14", NameTH: "คณะสัตวแพทยศาสตร์", NameEN: "Faculty of Veterinary Medicine"},
		{Code: "15", NameTH: "คณะบริหารธุรกิจ", NameEN: "Faculty of Business Administration"},
		{Code: "16", NameTH: "คณะเศรษฐศาสตร์", NameEN: "Faculty of Economics"},
		{Code: "17", NameTH: "คณะสถาปัตยกรรมศาสตร์", NameEN: "Faculty of Architecture"},
		{Code: "18", NameTH: "คณะสื่อสารมวลชน", NameEN: "Faculty of Mass Communication"},
		{Code: "19", NameTH: "คณะรัฐศาสตร์และรัฐประศาสนศาสตร์", NameEN: "Faculty of Political Science and Public Administration"},
		{Code: "20", NameTH: "คณะนิติศาสตร์", NameEN: "Faculty of Law"},
		{Code: "21", NameTH: "วิทยาลัยศิลปะ สื่อ และเทคโนโลยี", NameEN: "Faculty of Art, Media and Technology"},
		{Code: "22", NameTH: "คณะสาธารณสุขศาสตร์", NameEN: "Faculty of Public Health"},
		{Code: "23", NameTH: "วิทยาลัยการศึกษาและการจัดการทะเล", NameEN: "Faculty of Education and Sea Management"},
		{Code: "24", NameTH: "วิทยาลัยนานาชาตินวัตกรรมดิจิทัล", NameEN: "Faculty of International Digital Innovation"},
		{Code: "25", NameTH: "สถาบันนโยบายสาธารณะ", NameEN: "Faculty of Public Policy"},
		{Code: "26", NameTH: "สถาบันวิศวกรรมชีวการแพทย์", NameEN: "Faculty of Biomedical Engineering"},
		{Code: "27", NameTH: "สถาบันวิจัยวิทยาศาสตร์สุขภาพ", NameEN: "Faculty of Health Science Research Institute"},
		{Code: "28", NameTH: "วิทยาลัยพหุวิทยาการและสหวิทยาการ", NameEN: "Faculty of Interdisciplinary and Multidisciplinary Studies"},
		{Code: "54", NameTH: "สำนักบริการวิชาการ", NameEN: "Academic Service Center"},
		{Code: "59", NameTH: "สำนักงานพัฒนาคุณภาพการศึกษา", NameEN: "Office of Education Quality Development"},
		{Code: "61", NameTH: "สำนักงานมหาวิทยาลัย", NameEN: "Office of the University"},
		{Code: "64", NameTH: "ศูนย์วิจัยข้าวล้านนา", NameEN: "Lanna Rice Research Center"},
		{Code: "65", NameTH: "สถาบันภาษา", NameEN: "Language Institute"},
		{Code: "76", NameTH: "วิทยาลัยการศึกษาตลอดชีวิต", NameEN: "School of Life Long Education"},
		{Code: "107", NameTH: "ศูนย์บริหารจัดการความปลอดภัยฯ SHE", NameEN: "Center of Safety, Occupational Health and Environment"},
		{Code: "108", NameTH: "ศูนย์นวัตกรรมการสอนและการเรียนรู้ TLIC", NameEN: "Teaching and Learning Innovation Center"},
	}

	if err := db.Create(&faculty).Error; err != nil {
		log.Printf("❌ Error creating default faculty: %v", err)
	}

	courseStatus := []models.CourseStatus{
		{Status: "Not Started", Type: "None"},
		{Status: "In Progress", Type: "Processing"},
		{Status: "Almost Complete", Type: "Processing"},
		{Status: "Completed", Type: "Done"},
		{Status: "Pause", Type: "Pause"},
		{Status: "Cancel", Type: "Cancel"},
	}

	if err := db.Create(&courseStatus).Error; err != nil {
		log.Printf("❌ Error creating default course status: %v", err)
	}

	role := []models.Role{
		{Role: "admin"},
		{Role: "staff"},
		{Role: "LE"},
		{Role: "officer"},
		{Role: "user"},
	}

	if err := db.Create(&role).Error; err != nil {
		log.Printf("❌ Error creating default roles: %v", err)
	}

	log.Println("✅ Seeder created successfully")

	return nil
}
