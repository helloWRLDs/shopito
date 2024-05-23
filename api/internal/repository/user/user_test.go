package userrepository

import (
	"testing"
	"time"

	userdomain "shopito/api/internal/domain/user"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetByEmail(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected query and result
	email := "test@example.com"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin", "is_verified", "created_at", "updated_at"}).
		AddRow(1, "Test User", email, "hashed_password", false, true, time.Now(), time.Now())

	mock.ExpectQuery(`SELECT \* FROM users WHERE email=\$1`).WithArgs(email).WillReturnRows(rows)

	repo := New(db)
	user, err := repo.GetByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUserRepository_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	newUser := &userdomain.User{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: "new_password",
	}
	mock.ExpectQuery(`INSERT INTO users\(name, email, password\) VALUES\(\$1, \$2, \$3\) RETURNING id`).
		WithArgs(newUser.Name, newUser.Email, newUser.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := New(db)
	id, err := repo.Insert(newUser)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestUserRepository_ExistById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userId := 1
	mock.ExpectQuery(`SELECT EXISTS\(SELECT TRUE FROM users WHERE id=\$1\)`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	repo := New(db)
	exist := repo.ExistById(userId)

	assert.True(t, exist)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_IsVerified(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userId := 1
	mock.ExpectQuery(`SELECT is_verified FROM users WHERE id=\$1`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"is_verified"}).AddRow(true))

	repo := New(db)
	isVerified := repo.IsVerified(userId)

	assert.True(t, isVerified)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the input and expected query
	userId := 1
	updatedUser := &userdomain.User{
		Name:       "Updated User",
		Email:      "updateduser@example.com",
		Password:   "updated_password",
		IsVerified: true,
		IsAdmin:    false,
	}

	mock.ExpectExec(`UPDATE users SET name=\$1, email=\$2, password=\$3, is_verified=\$4, is_admin=\$5, updated_at=\$6 WHERE id=\$7`).
		WithArgs(updatedUser.Name, updatedUser.Email, updatedUser.Password, updatedUser.IsVerified, updatedUser.IsAdmin, sqlmock.AnyArg(), userId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := New(db)
	err = repo.Update(userId, updatedUser)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userId := 1

	mock.ExpectExec(`DELETE FROM users WHERE id=\$1`).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := New(db)
	err = repo.Delete(userId)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetById(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the input and expected result
	userId := 1
	expectedUser := &userdomain.User{
		ID:         userId,
		Name:       "Test User",
		Email:      "test@example.com",
		Password:   "hashed_password",
		IsAdmin:    false,
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin", "is_verified", "created_at", "updated_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.Password, expectedUser.IsAdmin, expectedUser.IsVerified, expectedUser.CreatedAt, expectedUser.UpdatedAt)

	mock.ExpectQuery(`SELECT \* FROM users WHERE id=\$1`).WithArgs(userId).WillReturnRows(rows)

	repo := New(db)
	user, err := repo.GetById(userId)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
	assert.Equal(t, expectedUser.IsAdmin, user.IsAdmin)
	assert.Equal(t, expectedUser.IsVerified, user.IsVerified)
	assert.Equal(t, expectedUser.CreatedAt, user.CreatedAt)
	assert.Equal(t, expectedUser.UpdatedAt, user.UpdatedAt)
}

func TestUserRepository_GetAll(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the expected query and result
	expectedUsers := []*userdomain.User{
		{ID: 1, Name: "User 1", Email: "user1@example.com", Password: "password1", IsAdmin: false, IsVerified: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "User 2", Email: "user2@example.com", Password: "password2", IsAdmin: true, IsVerified: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "is_admin", "is_verified", "created_at", "updated_at"})
	for _, user := range expectedUsers {
		rows.AddRow(user.ID, user.Name, user.Email, user.Password, user.IsAdmin, user.IsVerified, user.CreatedAt, user.UpdatedAt)
	}

	mock.ExpectQuery(`SELECT \* FROM users`).WillReturnRows(rows)

	repo := New(db)
	users, err := repo.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, len(expectedUsers), len(*users))
	for i, expectedUser := range expectedUsers {
		assert.Equal(t, *expectedUser, (*users)[i])
	}
}

func TestUserRepository_ExistByEmail(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Define the input email and expected result
	email := "test@example.com"
	expectedExist := true

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(expectedExist)
	mock.ExpectQuery(`SELECT EXISTS\(SELECT TRUE FROM users WHERE email=\$1\)`).WithArgs(email).WillReturnRows(rows)

	repo := New(db)
	exist := repo.ExistByEmail(email)

	assert.Equal(t, expectedExist, exist)
}
