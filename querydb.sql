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