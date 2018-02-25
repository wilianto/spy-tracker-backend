package user

import (
	"database/sql"
)

type psqlRepository struct {
	Conn *sql.DB
}

func (repo *psqlRepository) Store(u *User) (ID int64, err error) {
	query := "INSERT INTO users (username, password, name) VALUES ($1, $2, $3) RETURNING id"
	err = repo.Conn.QueryRow(query, u.Username, u.Password, u.Name).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

//NewPsqlRepository constructor to create a mysql user repository struct
func NewPsqlRepository(conn *sql.DB) Repository {
	return &psqlRepository{conn}
}
