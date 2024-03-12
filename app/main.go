package main

import (
	"log"

	"github.com/cocoasterr/kalbe_test/app/controllers"
	PGConfig "github.com/cocoasterr/kalbe_test/infra/db/postgres"
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/middleware"
	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db_dsn, _ := PGConfig.CreateDatabase("assesment_kalbe")
	db, err := PGConfig.ConnectPG(db_dsn)
	if err != nil {
		log.Fatal("Failed to Connect DB")
	}
	if err := PGConfig.CreateTable(db); err != nil {
		log.Fatal("Failed to Create Table!")
	}
	app := fiber.New()

	// init controller
	AttendanceRepo := PGRepository.NewAttendanceRepository(db)
	AttendanceService := services.NewAttendaceService(AttendanceRepo.Repository)
	AttendanceController := controllers.NewAttendanceController(*AttendanceService)

	EmployeeRepo := PGRepository.NewEmployeeRepository(db)
	EmployeeService := services.NewEmployeeService(EmployeeRepo.Repository)
	EmployeeController := controllers.NewEmployeeController(*EmployeeService)

	MasterDepartmentRepo := PGRepository.NewMasterDepartmentRepository(db)
	MasterDepartmentService := services.NewMasterDepartmentService(MasterDepartmentRepo.Repository)
	MasterDepartmentController := controllers.NewMasterDepartmentController(*MasterDepartmentService)

	MasterLocationRepo := PGRepository.NewMasterLocationRepository(db)
	MasterLocationService := services.NewMasterLocationtService(MasterLocationRepo.Repository)
	MasterLocationController := controllers.NewMasterLocationController(*MasterLocationService)

	MasterPositionRepo := PGRepository.NewMasterPositionRepository(db)
	MasterPositionService := services.NewMasterPositionService(MasterPositionRepo.Repository)
	MasterPositionController := controllers.NewMasterPositionController(*MasterPositionService)

	UserRepo := PGRepository.NewUserRepository(db)
	UserService := services.NewUserService(UserRepo.Repository)
	userController := controllers.NewUserController(*UserService)

	// api

	// attendance
	attendanceGroup := app.Group("/attendance")

	attendanceGroup.Post("/", middleware.DeserializeUser, AttendanceController.CreateAttendanceController)
	attendanceGroup.Get("/", middleware.DeserializeUser, AttendanceController.IndexAttendanceController)
	attendanceGroup.Get("/:id", middleware.DeserializeUser, AttendanceController.FindByAttendanceController)
	attendanceGroup.Put("/:id", middleware.DeserializeUser, AttendanceController.UpdateAttendanceController)
	attendanceGroup.Put("/delete/:id", middleware.DeserializeUser, AttendanceController.DeleteAttendanceController)
	app.Get("/report", middleware.DeserializeUser, AttendanceController.Reporting)

	// employee
	employeeGroup := app.Group("/employee")

	employeeGroup.Post("/", middleware.DeserializeUser, EmployeeController.CreateController)
	employeeGroup.Get("/", middleware.DeserializeUser, EmployeeController.IndexController)
	employeeGroup.Get("/:id", middleware.DeserializeUser, EmployeeController.FindByController)
	employeeGroup.Put("/:id", middleware.DeserializeUser, EmployeeController.UpdateController)
	employeeGroup.Put("/delete/:id", middleware.DeserializeUser, EmployeeController.DeleteController)

	// master department
	masterDepartmentGroup := app.Group("/masterDepartment")

	masterDepartmentGroup.Post("/", middleware.DeserializeUser, MasterDepartmentController.CreateController)
	masterDepartmentGroup.Get("/", middleware.DeserializeUser, MasterDepartmentController.IndexController)
	masterDepartmentGroup.Get("/:id", middleware.DeserializeUser, MasterDepartmentController.FindByController)
	masterDepartmentGroup.Put("/:id", middleware.DeserializeUser, MasterDepartmentController.UpdateController)
	masterDepartmentGroup.Put("/delete/:id", middleware.DeserializeUser, MasterDepartmentController.DeleteController)

	// master location
	masterLocationGroup := app.Group("/masterLocation")

	masterLocationGroup.Post("/", middleware.DeserializeUser, MasterLocationController.CreateController)
	masterLocationGroup.Get("/", middleware.DeserializeUser, MasterLocationController.IndexController)
	masterLocationGroup.Get("/:id", middleware.DeserializeUser, MasterLocationController.FindByController)
	masterLocationGroup.Put("/:id", middleware.DeserializeUser, MasterLocationController.UpdateController)
	masterLocationGroup.Put("/delete/:id", middleware.DeserializeUser, MasterLocationController.DeleteController)

	// master position
	masterPositionGroup := app.Group("/masterPosition")

	masterPositionGroup.Post("/", middleware.DeserializeUser, MasterPositionController.CreateController)
	masterPositionGroup.Get("/", middleware.DeserializeUser, MasterPositionController.IndexController)
	masterPositionGroup.Get("/:id", middleware.DeserializeUser, MasterPositionController.FindByController)
	masterPositionGroup.Put("/:id", middleware.DeserializeUser, MasterPositionController.UpdateController)
	masterPositionGroup.Put("/delete/:id", middleware.DeserializeUser, MasterPositionController.DeleteController)

	authGroup := app.Group("/auth")

	authGroup.Post("/register", userController.RegisterUser)
	authGroup.Post("/login", userController.Login)
	// authGroup.Post("/logout", userController.Logout)

	log.Fatal(app.Listen(":3000"))
}

/*
SELECT
    a.Created_at AS Date,
    e.Employee_code,
    e.Employee_name,
    d.Department_name,
    p.Position_name,
    l.Location_name,
    a.Absent_in,
    a.Absent_out
FROM
    Attendance a
JOIN
    Employee e ON a.Employee_id = e.Employee_id
JOIN
    MasterDepartment d ON e.Department_id = d.Department_id
JOIN
    MasterPosition p ON e.Position_id = p.Position_id
JOIN
    MasterLocation l ON a.Location_id = l.Location_id
WHERE
    a.Absent_in BETWEEN '12-04-2024' AND '12-04-2024'
ORDER BY
    a.Created_at;



	attendance_id = d23ecbd6-0fbb-4779-b0d9-10e77c22b102
	employee_id = 2f2b44cc-15c0-4996-88cc-361b78cebe5f
	location_id = 1a842c66-2ef6-48a9-9b33-79a6f2d8f29d

	department_id = fbc7fabd-ac95-455a-b464-9d1535f04b1b
	employee_id = 2f2b44cc-15c0-4996-88cc-361b78cebe5f
	position_id = 73ba0301-fe62-499e-a97f-e2f3ba15ef59

	department_id = fbc7fabd-ac95-455a-b464-9d1535f04b1b
	location_id = 1a842c66-2ef6-48a9-9b33-79a6f2d8f29d

	position_id = 73ba0301-fe62-499e-a97f-e2f3ba15ef59

*/
