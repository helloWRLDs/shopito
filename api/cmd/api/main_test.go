package main_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	productsdelivery "shopito/api/internal/delivery/http/products"
	"shopito/api/pkg/datastore/postgres"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// Connect to test database
	testDB, err := postgres.Open("admin", "admin", "mydb")
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}

	// Ensure the database is clean before running tests
	_, err = testDB.Exec("TRUNCATE TABLE products RESTART IDENTITY CASCADE;")
	if err != nil {
		log.Fatalf("Could not truncate tables: %v", err)
	}

	code := m.Run()

	// Clean up after tests
	testDB.Close()

	os.Exit(code)
}

func setupRouter() chi.Router {
	router := chi.NewRouter()
	router.Mount("/products", productsdelivery.New(testDB).Routes())
	return router
}

func TestCreateAndGetProduct(t *testing.T) {
	router := setupRouter()

	product := map[string]interface{}{
		"name":        "Test Product",
		"img_url":     "http://example.com/image.png",
		"price":       100,
		"stock":       50,
		"category_id": 1,
	}
	productJSON, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdProduct map[string]interface{}
	json.NewDecoder(rr.Body).Decode(&createdProduct)
	assert.Equal(t, product["name"], createdProduct["name"])
	assert.Equal(t, product["img_url"], createdProduct["img_url"])
	assert.Equal(t, product["price"], createdProduct["price"])
	assert.Equal(t, product["stock"], createdProduct["stock"])
	assert.Equal(t, product["category_id"], createdProduct["category_id"])

	// Test get product
	productID := createdProduct["id"].(int)
	req, _ = http.NewRequest("GET", "/products/"+strconv.Itoa(int(productID)), nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fetchedProduct map[string]interface{}
	json.NewDecoder(rr.Body).Decode(&fetchedProduct)
	assert.Equal(t, product["name"], fetchedProduct["name"])
	assert.Equal(t, product["img_url"], fetchedProduct["img_url"])
	assert.Equal(t, product["price"], fetchedProduct["price"])
	assert.Equal(t, product["stock"], fetchedProduct["stock"])
	assert.Equal(t, product["category_id"], fetchedProduct["category_id"])
}
