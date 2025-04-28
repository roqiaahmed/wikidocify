package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/roqiaahmed/wikidocify/pkg/documents"
)

func main() {
  r := gin.Default()

  // r.GET("/documents", DocumentHandler.GetAllDocuments)
  // r.GET("/documents/:id", DocumentHandler.GetDocument)
  r.POST("/documents", DocumentHandler.CreateDocument)
  // r.PUT("/documents/:id", DocumentHandler.UpdateDocument)
  // r.DELETE("/documents/:id", DocumentHandler.DeleteDocument)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type DocumentHandler struct {
  store DocumentStore
}

type DocumentStore interface {
  Create(doc documents.Document) error
  // Get(id string) (doc documents.Document, error) 
  // GetAll() (map[string]documents.Document, error)
  // Update(id string, doc documents.Document) error
  // Delete(id string) error
}


func (h *DocumentHandler) CreateDocument(c *gin.Context){
  var doc documents.Document
  if err := c.ShouldBindJSON(&doc); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if err := h.store.Create(doc); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  
  h.store.Create(doc)
  c.JSON(http.StatusCreated, doc)
}

// func (h *DocumentHandler) GetAllDocuments(c *gin.Context){}

// func (h *DocumentHandler) GetDocument(c *gin.Context){}

// func (h *DocumentHandler) UpdateDocument(c *gin.Context){}

// func (h *DocumentHandler) DeleteDocument(c *gin.Context){}

