package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"

	"shopCart/model"
	"shopCart/module/user/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadDatabase(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	// mock sql db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}
	if db == nil {
		t.Error("mock db is null")
	}
	if mock == nil {
		t.Error(" sql moock is null")
	}

	// open gorm db
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}
	if gdb == nil {
		t.Error("gorm db is null")
	}

	return db, gdb, mock
}

func TestGetUserList(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	mockUsers := []*model.User{
		&model.User{
			ID:       uint(1),
			Username: "123",
			Password: "123",
		},
		&model.User{
			ID:       uint(2),
			Username: "123",
			Password: "123",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Password).
		AddRow(mockUsers[1].ID, mockUsers[1].Username, mockUsers[1].Password)

	query := `SELECT * FROM "users"`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	userRepo := repository.NewUserRepository(gdb)
	users, err := userRepo.GetUserList(make(map[string]interface{}))
	if err != nil {
		t.Errorf("Failed to get user list, got error: %v", err)
	}

	assert.Equal(t, mockUsers, users)
}

func TestGetUser(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).
		AddRow(mockUser.ID, mockUser.Username, mockUser.Password)

	query := `SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	userRepo := repository.NewUserRepository(gdb)

	queryUser := new(model.User)
	queryUser.ID = 1
	user, err := userRepo.GetUser(queryUser)
	if err != nil {
		t.Errorf("Failed to get a user with id, got error: %v", err)
	}

	assert.Equal(t, mockUser, user)
}

func TestCreateUser(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	// mock creating a user into a virtual db
	sqlInsert := `INSERT INTO "users" ("username","password","id") VALUES ($1,$2,$3) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(mockUser.Username, mockUser.Password, mockUser.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	userRepo := repository.NewUserRepository(gdb)

	// test createUser function
	newUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}
	if _, err := userRepo.CreateUser(newUser); err != nil {
		t.Errorf("error was not expected while create a user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	mockUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}

	// mock creating a user into a virtual db
	mock.ExpectBegin()
	sqlUpdate := `UPDATE "users" SET "username"=$1,"password"=$2 WHERE "id" = $3`
	mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).
		WithArgs("456", "456", mockUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepo := repository.NewUserRepository(gdb)

	newUser := &model.User{
		ID:       uint(1),
		Username: "456",
		Password: "456",
	}
	if _, err := userRepo.UpdateUser(newUser); err != nil {
		t.Errorf("error was not expected while updated the user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestModifyUser(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	// mock creating a user into a virtual db
	mock.ExpectBegin()
	sqlModify := `UPDATE "users" SET "password"=$1,"username"=$2 WHERE "id" = $3`
	mock.ExpectExec(regexp.QuoteMeta(sqlModify)).
		WithArgs("123", "123", uint(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepo := repository.NewUserRepository(gdb)

	newUser := map[string]interface{}{
		"username": "123",
		"password": "123",
	}
	if _, err := userRepo.ModifyUser(&model.User{ID: 1}, newUser); err != nil {
		t.Errorf("error was not expected while modified the user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, gdb, mock := loadDatabase(t)
	defer db.Close()

	// mock creating a user into a virtual db
	mock.ExpectBegin()
	sqlDelete := `DELETE FROM "users" WHERE "users"."id" = $1`
	mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
		WithArgs(uint(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepo := repository.NewUserRepository(gdb)

	newUser := &model.User{
		ID:       uint(1),
		Username: "123",
		Password: "123",
	}
	if err := userRepo.DeleteUser(newUser); err != nil {
		t.Errorf("error was not expected while deleted the user: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
