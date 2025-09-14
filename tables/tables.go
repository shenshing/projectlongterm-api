// Create an array of query to create tables.
// Create a function to loop, create tables.
package tables

var Query_to_create_table = [5]string{
	`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			firstname CHAR(255) NOT NULL,
			lastname CHAR(255),
			email CHAR(255) NOT NULL,
			login_type CHAR(255) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`,
	`CREATE TABLE IF NOT EXISTS articles (
			id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
			lang CHAR(10) NOT NULL, 
			title CHAR(255) NOT NULL,
			slug CHAR(255) NOT NULL,
			body TEXT NOT NULL,
			html TEXT NOT NULL,  
			author CHAR(255) NOT NULL, 
			read_time CHAR(255),
			author_name CHAR(255) NOT NULL,
			author_profile_url VARCHAR(1000),
			original_url VARCHAR(1000),
			thumbnail_url VARCHAR(1000),
			type CHAR(255),
			excerpt VARCHAR(1000) NOT NULL,
			avatar VARCHAR(1000) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		); 
	`,
	`
		CREATE TABLE IF NOT EXISTS article_comments (
			id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
			article_id CHAR(36) NOT NULL, 
			user_id INT NOT NULL,
			description CHAR(255) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);	
	`,
	`
		CREATE TABLE IF NOT EXISTS article_reactions (
			id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
			article_id CHAR(36) NOT NULL, 
			user_id INT NOT NULL,
			articles_reactions_types_id CHAR(36) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);	
	`,
	`
		CREATE TABLE IF NOT EXISTS article_reaction_types (
			id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
			description CHAR(255),
			emoji CHAR(255) NOT NULL,
			version VARCHAR(50) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);	
	`,
}

var Update_table_string = [1]string{
	// Alter table query

	// `
	// 	ALTER TABLE articles
	// 	ADD COLUMN type CHAR(255);
	// `,
}
