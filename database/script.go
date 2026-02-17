package database

var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id VARCHAR(50) NOT NULL,
	username VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role ENUM('ADMIN','USER') NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
`

var migrationScripts = []string{
	createTableUsers,
}
