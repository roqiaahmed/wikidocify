package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roqiaahmed/wikidocify/initializers"
	"github.com/roqiaahmed/wikidocify/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *gin.Engine

func TestMain(m *testing.M) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	initializers.DB = db
	initializers.DB.AutoMigrate(&models.Document{})
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.POST("/documents", CreateDocument)
	router.GET("/documents", GetAllDocuments)
	router.GET("/documents/:id", GetDocument)
	router.PUT("/documents/:id", UpdateDocument)
	router.DELETE("/documents/:id", DeleteDocument)
	m.Run()
}

func TestCreateDocument(t *testing.T) {

	documentData := models.Document{
		ID:        "1",
		Title:     "Test Document",
		Author:    "Test Author",
		Content:   "This is a test document.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	body, _ := json.Marshal(documentData)
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdDocument models.Document

	json.Unmarshal(w.Body.Bytes(), &createdDocument)
	assert.Equal(t, documentData.Title, createdDocument.Title)
	assert.Equal(t, documentData.Author, createdDocument.Author)
	assert.Equal(t, documentData.Content, createdDocument.Content)
	assert.NotEmpty(t, createdDocument.ID)
	assert.NotEmpty(t, createdDocument.CreatedAt)
	assert.NotEmpty(t, createdDocument.UpdatedAt)
	// Clean up the test database
	initializers.DB.Delete(&models.Document{}, createdDocument.ID)

}

func TestGetAllDocuments(t *testing.T) {

	documents := []models.Document{
		{
			ID:        "1",
			Title:     "Test Document",
			Author:    "Test Author",
			Content:   "This is a test document.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()},
		{
			ID:        "2",
			Title:     "Test Document",
			Author:    "Test Author",
			Content:   "This is a test document.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()},
		{
			ID:        "3",
			Title:     "Test Document",
			Author:    "Test Author",
			Content:   "This is a test document.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()}}

	for _, doc := range documents {
		initializers.DB.Create(&doc)
	}
	req, _ := http.NewRequest("GET", "/documents", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var responseDocuments []models.Document
	json.Unmarshal(w.Body.Bytes(), &responseDocuments)
	assert.Equal(t, len(documents), len(responseDocuments))
	for i, doc := range documents {
		assert.Equal(t, doc.ID, responseDocuments[i].ID)
		assert.Equal(t, doc.Title, responseDocuments[i].Title)
		assert.Equal(t, doc.Author, responseDocuments[i].Author)
		assert.Equal(t, doc.Content, responseDocuments[i].Content)
	}
	// Clean up the test database
	for _, doc := range documents {
		initializers.DB.Delete(&models.Document{}, doc.ID)
	}

}

func TestGetDocument(t *testing.T) {
	document := models.Document{
		ID:        "1",
		Title:     "Test Document",
		Author:    "Test Author",
		Content:   "This is a test document.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	initializers.DB.Create(&document)
	req, _ := http.NewRequest("GET", "/documents/"+document.ID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var responseDocument models.Document
	json.Unmarshal(w.Body.Bytes(), &responseDocument)
	assert.Equal(t, document.ID, responseDocument.ID)
	assert.Equal(t, document.Title, responseDocument.Title)
	assert.Equal(t, document.Author, responseDocument.Author)
	assert.Equal(t, document.Content, responseDocument.Content)
	// Clean up the test database
	initializers.DB.Delete(&models.Document{}, document.ID)
}

func TestUpdateDocument(t *testing.T) {
	oldDocument := models.Document{
		ID:        "1",
		Title:     "Test Document",
		Author:    "Test Author",
		Content:   "This is a test document.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	initializers.DB.Create(&oldDocument)
	newDocument := models.Document{
		ID:        "1",
		Title:     "Updated Document",
		Author:    "Updated Author",
		Content:   "This is an updated test document.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	body, _ := json.Marshal(newDocument)
	req, _ := http.NewRequest("PUT", "/documents/"+oldDocument.ID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var updatedDocument models.Document
	json.Unmarshal(w.Body.Bytes(), &updatedDocument)
	assert.Equal(t, newDocument.ID, updatedDocument.ID)
	assert.Equal(t, newDocument.Title, updatedDocument.Title)
	assert.Equal(t, newDocument.Author, updatedDocument.Author)
	assert.Equal(t, newDocument.Content, updatedDocument.Content)
	// Clean up the test database
	initializers.DB.Delete(&models.Document{}, oldDocument.ID)
	initializers.DB.Delete(&models.Document{}, newDocument.ID)
}

func TestDeleteDocument(t *testing.T) {
	oldDocument := models.Document{
		ID:        "1",
		Title:     "Test Document",
		Author:    "Test Author",
		Content:   "This is a test document.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	initializers.DB.Create(&oldDocument)
	req, _ := http.NewRequest("DELETE", "/documents/"+oldDocument.ID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "success", response["status"])
}
