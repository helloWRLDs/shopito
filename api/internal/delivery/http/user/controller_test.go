package userdelivery_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	userdelivery "shopito/api/internal/delivery/http/user"
	userdomain "shopito/api/internal/domain/user"
	"shopito/api/pkg/datastore/postgres"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDB *sql.DB

func setup() {
	var err error
	testDB, err = postgres.Open("admin", "admin", "mydb")
	if err != nil {
		panic(err)
	}
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestGetUserController(t *testing.T) {
	delivery := userdelivery.New(testDB)

	userID := 0
	req, err := http.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(userID), nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userId", strconv.Itoa(userID))

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(delivery.GetUserController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user userdomain.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	require.NoError(t, err)
	assert.Equal(t, userID, user.ID)
}
