package PGConfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateDatabase(dbName string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to connect .env")
		return "", err
	}
	db_dsn := os.Getenv("PG_DB_DSN")
	db_dsn_con := fmt.Sprintf("%s?sslmode=disable", db_dsn)
	db, err := sql.Open("postgres", db_dsn_con)

	if err != nil {
		log.Fatal(err)
	}
	createDatabaseQuery := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", dbName)

	var existingDatabaseName string
	if err = db.QueryRow(createDatabaseQuery).Scan(&existingDatabaseName); err != nil {
		if err == sql.ErrNoRows {
			createDatabaseQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			_, err := db.Exec(createDatabaseQuery)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	new_db_dsn := fmt.Sprintf("%s%s?sslmode=disable", db_dsn, dbName)
	log.Println("Database Created!")
	return new_db_dsn, nil
}
func CreateTable(db *sql.DB) error {
	// sqlFile := "querydb.sql"
	// sqlBytes, err := os.ReadFile(sqlFile)
	// if err != nil {
	// 	return err
	// }
	// sqlScript := string(sqlBytes)

	// fmt.Println(sqlScript)
	sqlScript := `
	CREATE TABLE IF NOT EXISTS users (
		user_id VARCHAR(255) PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	
	-- Master Department
	CREATE TABLE IF NOT EXISTS MasterDepartment (
		Department_id VARCHAR(255) PRIMARY KEY,
		Department_name VARCHAR(255) NOT NULL,
		Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Created_by VARCHAR(255),
		Updated_at TIMESTAMP,
		Updated_by VARCHAR(255),
		Deleted_at TIMESTAMP
	);
	
	-- Master Position
	CREATE TABLE IF NOT EXISTS MasterPosition (
		Position_id VARCHAR(255) PRIMARY KEY,
		Department_id VARCHAR(255),
		Position_name VARCHAR(255) NOT NULL,
		Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Created_by VARCHAR(255),
		Updated_at TIMESTAMP,
		Updated_by VARCHAR(255),
		Deleted_at TIMESTAMP
	);
	
	-- Master Location
	CREATE TABLE IF NOT EXISTS MasterLocation (
		Location_id VARCHAR(255) PRIMARY KEY,
		Location_name VARCHAR(255) NOT NULL,
		Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Created_by VARCHAR(255),
		Updated_at TIMESTAMP,
		Updated_by VARCHAR(255),
		Deleted_at TIMESTAMP
	);
	
	-- Employee
	CREATE TABLE IF NOT EXISTS Employee (
		Employee_id VARCHAR(255) PRIMARY KEY,
		Employee_code VARCHAR(10) UNIQUE NOT NULL,
		Employee_name VARCHAR(255) NOT NULL,
		Password VARCHAR(255) NOT NULL,
		Department_id VARCHAR(255),
		Position_id VARCHAR(255),
		Superior VARCHAR(255),
		Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Created_by VARCHAR(255),
		Updated_at TIMESTAMP,
		Updated_by VARCHAR(255),
		Deleted_at TIMESTAMP
	);
	
	-- Attendance
	CREATE TABLE IF NOT EXISTS Attendance (
		Attendance_id VARCHAR(255) PRIMARY KEY,
		Employee_id VARCHAR(255),
		Location_id VARCHAR(255),
		Absent_in TIMESTAMP,
		Absent_out TIMESTAMP,
		Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Created_by VARCHAR(255),
		Updated_at TIMESTAMP,
		Updated_by VARCHAR(255),
		Deleted_at TIMESTAMP
	);
	
	`
	// absen_in := time.Now().Format("2006-01-02")
	// absen_out := time.Now().Format("2006-01-02")

	// data := []string{
	// 	fmt.Sprintf(`INSERT INTO attendance (attendance_id, employee_id, location_id, absent_in, absent_out, created_at, created_by, updated_at)
	// VALUES ('d23ecbd6-0fbb-4779-b0d9-10e77c22b102', '2f2b44cc-15c0-4996-88cc-361b78cebe5f', '1a842c66-2ef6-48a9-9b33-79a6f2d8f29d', '%s 08:00:00', '%s 17:00:00', '2024-03-12 22:21:39.180284', 'ido', '2024-03-12 22:21:39.180285');
	// `, absen_in, absen_out),
	// 	`INSERT INTO employee (employee_id, employee_code, employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at)
	// VALUES ('2f2b44cc-15c0-4996-88cc-361b78cebe5f', 'EMP002', 'john dho', 'secretpassword', 'fbc7fabd-ac95-455a-b464-9d1535f04b1b', '73ba0301-fe62-499e-a97f-e2f3ba15ef59', '456', '2024-03-12 22:18:28.556977', 'AdminUser', '2024-03-12 22:27:30.122736');
	// `,
	// 	`INSERT INTO masterdepartment (department_id, department_name, created_at, created_by, updated_at)
	// VALUES ('fbc7fabd-ac95-455a-b464-9d1535f04b1b', 'Human Resources', '2024-03-12 22:19:06.046099', 'AdminUser', '2024-03-12 22:19:06.046099');
	// `,
	// 	`INSERT INTO masterlocation (location_id, location_name, created_at, created_by, updated_at)
	// VALUES ('1a842c66-2ef6-48a9-9b33-79a6f2d8f29d', 'Office A', '2024-03-12 22:19:47.469693', 'AdminUser', '2024-03-12 22:19:47.469694');
	// `,
	// 	`INSERT INTO masterposition (position_id, department_id, position_name, created_at, created_by, updated_at)
	// VALUES ('73ba0301-fe62-499e-a97f-e2f3ba15ef59', 'fbc7fabd-ac95-455a-b464-9d1535f04b1b', 'Manager', '2024-03-12 22:20:45.042085', 'AdminUser', '2024-03-12 22:20:45.042086');
	// `,
	// }
	_, err := db.Exec(sqlScript)
	if err != nil {
		if err.Error()[len(err.Error())-14:len(err.Error())] == "already exists" {
			log.Println("table already exists")
			// for _, sqlScript := range data {
			// 	_, err := db.Exec(sqlScript)
			// 	if err != nil {
			// 		fmt.Println("Error executing SQL script:", err)
			// 	}
			// }
			log.Println("create dummy succesfully")

			return nil
		}
	}
	log.Println("Table Created!")
	// for _, sqlScript := range data {
	// 	_, err := db.Exec(sqlScript)
	// 	if err != nil {
	// 		fmt.Println("Error executing SQL script:", err)
	// 	}
	// }
	log.Println("create dummy succesfully")
	return nil
}
