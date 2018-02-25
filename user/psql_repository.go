package user

import (
	"database/sql"
)

type psqlRepository struct {
	Conn *sql.DB
}

func (repo *psqlRepository) Store(u *User) (ID int64, err error) {
	query := `INSERT INTO users (username, password, name) VALUES (?, ?, ?)`
	stmt, err := repo.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.Name)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

//NewPsqlRepository constructor to create a mysql user repository struct
func NewPsqlRepository(conn *sql.DB) Repository {
	return &psqlRepository{conn}
}
