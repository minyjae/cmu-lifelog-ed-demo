// @title Life Long Education API
// @version 1.0.0
// @description API สำหรับระบบ Life Long Education ที่พัฒนาด้วย Go Fiber
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT token สำหรับการยืนยันตัวตน ให้ใส่ token ในรูปแบบ: Bearer <token>

package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/minyjae/cmu-life-long-ed-api/docs"
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/http/handlers"
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/http/routes"
	"github.com/minyjae/cmu-life-long-ed-api/internal/adapters/presistence/repositories"
	"github.com/minyjae/cmu-life-long-ed-api/internal/config"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/services"
	"github.com/redis/rueidis"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db := config.SetupDatabase(cfg)

	redisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{cfg.RedisAddr},
		Password:    cfg.RedisPassword,
	})
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer redisClient.Close()

	listQueueRepo := repositories.NewListQueueRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	orderMappingRepo := repositories.NewOrderMappingRepository(db)
	// staffRepo := repositories.NewStaffRepository(db)
	staffStatusRepo := repositories.NewStaffStatusRepository(db)
	// userRepo := repositories.NewUserRepository(db)
	facultyRepo := repositories.NewFacultyRepository(db)
	courseStatusRepo := repositories.NewCourseStatusRepository(db)
	usersRepo := repositories.NewUsersRepository(db)
	roleRepo := repositories.NewRoleRepository(db)

	listQueueService := services.NewListQueueServiceImpl(listQueueRepo, orderMappingRepo, facultyRepo, staffStatusRepo, redisClient)
	orderService := services.NewOrderServiceImpl(orderRepo, orderMappingRepo, listQueueRepo)
	// staffService := services.NewStaffServiceImpl(staffRepo)
	staffStatusService := services.NewStaffStatusServiceImpl(staffStatusRepo, listQueueRepo, redisClient)
	// userService := services.NewUserServiceImpl(userRepo)
	requireRoleService := services.NewRequireRoleService(usersRepo)
	facultyService := services.NewFacultyServiceImpl(facultyRepo, redisClient)
	courseStatusService := services.NewCourseStatusServiceImpl(courseStatusRepo, redisClient)
	usersService := services.NewUsersServiceImpl(usersRepo, facultyRepo, redisClient)
	roleService := services.NewRoleServiceImpl(roleRepo, redisClient)

	signHandler := handlers.NewSigninHandler(usersService, facultyService)
	listQueueHandler := handlers.NewListQueueHandler(listQueueService, staffStatusService, orderService, usersService)
	orderHandler := handlers.NewOrderHandler(orderService)
	// staffHandler := handlers.NewStaffHandler(usersService)
	staffStatusHandler := handlers.NewStaffStatusHandler(staffStatusService)
	// userHandler := handlers.NewUserHandler(usersService)
	facultyHandler := handlers.NewFacultyHandler(facultyService)
	courseStatusHandler := handlers.NewCourseStatusHandler(courseStatusService)
	usersHandler := handlers.NewUsersHandler(usersService, facultyService)
	allUserHandler := handlers.NewAllUserHandler(usersService)
	roleHandler := handlers.NewRoleHandler(roleService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CorsAllows, // หรือ "*" ถ้าไม่จำกัด
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true, // ถ้าคุณใช้ cookie หรือ auth header
	}))

	// health check อยู่นอก prefix เสมอ (สำหรับ Docker/load balancer)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})

	prefix := cfg.AppPrefix // e.g. "/queue-doc-api"
	var r fiber.Router = app
	if prefix != "" && prefix != "/" {
		r = app.Group(prefix) // <-- ทุก route ต่อจากนี้จะอยู่ใต้ /queue-doc-api
	}

	routes.SetupRoute(r, listQueueHandler, orderHandler, staffStatusHandler, usersHandler, signHandler, requireRoleService, facultyHandler, courseStatusHandler, allUserHandler, roleHandler)
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
