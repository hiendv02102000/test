package entity

// Users struct
type Users struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (i *Users) QueryCreateTable() string {
	query := `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		password varchar(255) NOT NULL,
		PRIMARY KEY (id)
		)
	`

	return query
}
