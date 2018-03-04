package user_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilianto/spy-tracker-backend/user"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStore_Success(t *testing.T) {
	usr := &user.User{
		Username: "wilianto",
		Password: "hash_password",
		Name:     "Wilianto Indrawan",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when creating sql mock: %s", err.Error())
	}
	defer db.Close()

	query := "INSERT INTO users \\(username, password, name\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id"
	rows := sqlmock.NewRows([]string{"id"}).AddRow("5")
	mock.ExpectQuery(query).
		WithArgs(usr.Username, usr.Password, usr.Name).
		WillReturnRows(rows)

	repo := user.NewPsqlRepository(db)
	lastID, err := repo.Store(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), lastID)
}

func TestStore_Error(t *testing.T) {
	usr := &user.User{
		Username: "wilianto",
		Password: "hash_password",
		Name:     "Wilianto Indrawan",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when creating sql mock: %s", err.Error())
	}
	defer db.Close()

	query := "INSERT INTO users \\(username, password, name\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id"
	mock.ExpectQuery(query).
		WithArgs(usr.Username, usr.Password, usr.Name).
		WillReturnError(errors.New("Insert error"))

	repo := user.NewPsqlRepository(db)
	lastID, err := repo.Store(usr)
	assert.Equal(t, "Insert error", err.Error())
	assert.Equal(t, int64(0), lastID)
}
