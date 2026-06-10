package routes

import (
"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/http/handlers"
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/http/middleware"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

func SetupRoute(r fiber.Router, listQueueHandler *handlers.ListQueueHandler, orderHandler *handlers.OrderHandler, staffStatusHandler *handlers.StaffStatusHandler, usersHandler *handlers.UsersHandler, signInHandler *handlers.SignInHandler, requireRoleService services.RequireRoleService, facultyHandler *handlers.FacultyHandler, courseStatusHandler *handlers.CourseStatusHandler, allUserHandler *handlers.AllUserHandler, roleHandler *handlers.RoleHandler) {
	r.Get("/swagger/*", swagger.HandlerDefault)

	r.Post("/api/auth", func(c *fiber.Ctx) error {
		return signInHandler.SignIn(c)
	})
	r.Post("/api/auth/register", func(c *fiber.Ctx) error {
		return signInHandler.Register(c)
	})

	api := r.Group("/api", middleware.JWTProtected())

	// route ที่ให้ user เห็น
	//========================================================================================================
	sharedListQueue := api.Group("/", middleware.RequireRole(requireRoleService, "admin", "staff", "LE", "officer", "user"))
	// sharedListQueue role ที่สามารถใช้งานได้ "admin", "staff", "LE", "officer", "user"
	sharedListQueue.Get("/course/status", courseStatusHandler.GetCourseStatus)     // get courseStatus สำหรับโชว์หน้าสร้าง listqueue
	sharedListQueue.Post("/listqueue/owner", listQueueHandler.GetListQueueByOwner) // get listqueue ของอาจารย์คนๆนั้นที่เปิดหลักสูตร

	sharedListQueue.Get("/user/me", allUserHandler.GetCurrentUser) // get ข้อมูลจาก database ของคนที่ signin ขณะนั้น

	// route ที่ให้ officer เห็น
	//========================================================================================================
	officer := api.Group("/", middleware.RequireRole(requireRoleService, "admin", "staff", "LE", "officer"))
	// officer role ที่สามารถใช้งานได้ "admin", "staff", "LE", "officer"
	officer.Post("/listqueue/faculty", listQueueHandler.GetListQueueByFaculty) // get listqueue ตาม faculty ของ เจ้าหน้าที่คณะ

	// route ที่ให้ le เห็น
	//========================================================================================================
	le := api.Group("/", middleware.RequireRole(requireRoleService, "admin", "staff", "LE"))
	// officer role ที่สามารถใช้งานได้ "admin", "staff", "LE"
	le.Get("/listqueue", listQueueHandler.GetListQueue)                     // get listqueue ทั้งหมด (ไม่ว่าจะเสร็จหรือไม่เสร็จโชว์หมด)
	le.Get("/listqueue/status/notyet", listQueueHandler.GetListQueueNotYet) // get listqueue ที่ยังไม่เสร็จ (staffStatus != Done)

	le.Post("/listqueue/staffstatus", listQueueHandler.GetListQueueByStaffStatus) // get listqueue ด้วย staffStatus

	le.Post("/listqueue/coursestatus", listQueueHandler.GetListQueueByCourseStatus) // get listqueue ด้วย courseStatus

	// route ที่ให้ staff เห็น
	//========================================================================================================
	staff := api.Group("/", middleware.RequireRole(requireRoleService, "admin", "staff"))
	// officer role ที่สามารถใช้งานได้ "admin", "staff"
	staff.Post("/listqueue", listQueueHandler.CreateRequest)               // post สร้าง listqueue
	staff.Put("/listqueue", listQueueHandler.UpdateListQueue)              // put อัพเดท listqueue
	staff.Delete("/listqueue/:id", listQueueHandler.RemoveListQueueForDev) // delete listqueue สำหรับ developer เท่านั้น deploy จริงไม่มีการลบ สำหรับฉุกเฉิน !!!!!!!

	staff.Get("/order/:id", orderHandler.GetOrderFromListQueueID) // get order ของ listqueue นั้นๆ (ดูสถานะ order ปัจจุบันของ listqueue นั้นๆ)
	staff.Put("/order", orderHandler.UpdateOrder)                 // put อัพเดท order

	staff.Get("/staffstatus", staffStatusHandler.GetStaffStatus)                            // get staffStatus สำหรับโชว์หน้าสร้าง listqueue
	staff.Put("/listqueue/:id/status/:staff_status_id", listQueueHandler.UpdateStaffStatus) // put อัพเดท staffStatus ของ listqueue id นั้นๆ
	staff.Put("/listqueue/:id/priority/:priority", listQueueHandler.UpdatePriority)
	// staff.Get("/user/:faculty", usersHandler.GetUserByFaculty) //pass

	staff.Get("/staff", usersHandler.GetStaff) // get ข้อมูล staff สำหรับนำไปโชว์ตอนสร้าง listqueue ว่าใครเป็นคนดูแลหลักสูตรนี้

	staff.Get("/faculty", facultyHandler.GetAllFaculty) // get faculty ทั้งหมด สำหรับโชว์หน้าสร้าง listqueue
	// staff.Get("/userstatus", userStatusHandler.GetAllUserStatus)

	// route ที่ให้ admin เห็น
	//========================================================================================================
	admin := api.Group("/", middleware.RequireRole(requireRoleService, "admin"))
	// officer role ที่สามารถใช้งานได้ "admin"
	admin.Get("/user/all", usersHandler.GetAllUsers)          // get user ทุกคน เป็นหน้าสำหรับ manage คนสำหรับ admin
	admin.Post("/user/:email/:role", usersHandler.CreateUser) // เพิ่ม user ต่างๆด้วย role ที่ส่งเข้ามา
	admin.Delete("/user/:id", usersHandler.RemoveUser)        // ลบ user คนๆนั้นออก
	admin.Put("user/updateinfo/:email", usersHandler.UpdateUserInfo)

	admin.Post("/order", orderHandler.CreateOrder)         // สร้าง order อันใหม่มา
	admin.Put("/order/name", orderHandler.UpdateOrderName) // อัพเดทชื่อ order
	admin.Delete("/order/:id", orderHandler.RemoveOrder)   // ลบ order ทิ้ง

	admin.Post("/staffstatus", staffStatusHandler.CreateStaffStatus)         // สร้าง staffStatus
	admin.Put("/staffstatus/name", staffStatusHandler.UpdateStaffStatusName) // อัพเดทชื่อ staffStatus
	admin.Delete("/staffstatus/:id", staffStatusHandler.RemoveStaffStatus)   // ลบ staffStatus

	admin.Get("/role", roleHandler.GetRole) // get role สำหรับโชว์หน้า เพิ่ม user ใหม่ขึ้นมา
}
