package test

import (
	"belajar-resfull-api/app"
	"belajar-resfull-api/controller"
	"belajar-resfull-api/helper"
	"belajar-resfull-api/middleware"
	"belajar-resfull-api/model/web"
	"belajar-resfull-api/repository"
	"belajar-resfull-api/service"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar__golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func RequestBody() map[string]interface{} {
	return map[string]interface{}{
		"name": "SmartPhone",
	}
}

func toEncode(data interface{}) string {
	byte, _ := json.Marshal(data)
	return string(byte)
}

func ExpectResult(data web.WebResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":   data.Code,
		"status": data.Status,
		"data":   data.Data,
	}
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	exampleRequest := RequestBody()
	requestBody := strings.NewReader(toEncode(exampleRequest))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var resultCreateCategory map[string]interface{}
	decode := json.NewDecoder(response.Body)
	_ = decode.Decode(&resultCreateCategory)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, exampleRequest["name"], resultCreateCategory["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	exampleRequest := RequestBody()
	exampleRequest["name"] = ""
	requestBody := strings.NewReader(toEncode(exampleRequest))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var resultCreateCategory map[string]interface{}
	decode := json.NewDecoder(response.Body)
	_ = decode.Decode(&resultCreateCategory)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "Bad Request", resultCreateCategory["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)
	exampleRequest := RequestBody()
	requestBody := strings.NewReader(toEncode(exampleRequest))
	// create categories
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	exampleRequest["name"] = "Laptop"
	requestBody = strings.NewReader(toEncode(exampleRequest))
	//categoryId := 1

	request = httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	var expectResult map[string]interface{}
	json.Unmarshal(body, &expectResult)

	assert.Equal(t, 200, int(expectResult["code"].(float64)))
	assert.Equal(t, "ok", expectResult["status"])
	assert.Equal(t, "Laptop", expectResult["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)
	exampleRequest := RequestBody()
	requestBody := strings.NewReader(toEncode(exampleRequest))
	// create categories
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	exampleRequest["name"] = ""
	requestBody = strings.NewReader(toEncode(exampleRequest))
	//categoryId := 1

	request = httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	var expectResult map[string]interface{}
	json.Unmarshal(body, &expectResult)

	assert.Equal(t, 400, int(expectResult["code"].(float64)))
	assert.Equal(t, "Bad Request", expectResult["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	exampleRequest := RequestBody()
	requestBody := strings.NewReader(toEncode(exampleRequest))
	// create categories
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	request = httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response = recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var expectResult map[string]interface{}
	json.Unmarshal(body, &expectResult)

	assert.Equal(t, 200, int(expectResult["code"].(float64)))
	assert.Equal(t, "oke", expectResult["status"])
	assert.Equal(t, "SmartPhone", expectResult["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var expectResult map[string]interface{}
	json.Unmarshal(body, &expectResult)

	assert.Equal(t, http.StatusNotFound, int(expectResult["code"].(float64)))
	assert.Equal(t, "Not Found", expectResult["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)
	exampleRequest := RequestBody()
	requestBody := strings.NewReader(toEncode(exampleRequest))

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	request = httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "Rahasia")
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response = recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestListCategorySuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/1", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var expectResult map[string]interface{}
	json.Unmarshal(body, &expectResult)

	assert.Equal(t, http.StatusUnauthorized, int(expectResult["code"].(float64)))
	assert.Equal(t, "Unauthorized", expectResult["status"])
}
