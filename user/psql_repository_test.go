package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilianto/spy-tracker-backend/user"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStore(t *testing.T) {
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

	query := "INSERT INTO users \\(username, password, name\\) VALUES \\(\\?, \\?, \\?\\)"
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs(usr.Username, usr.Password, usr.Name).
		WillReturnResult(sqlmock.NewResult(5, 1))

	repo := user.NewPsqlRepository(db)
	lastID, err := repo.Store(usr)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), lastID)
}
