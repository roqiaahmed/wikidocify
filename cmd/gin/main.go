package main

import (
  // "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/roqiaahmed/wikidocify/pkg/documents"
)

func main() {
  r := gin.Default()
  store := documents.NewStore()      
  handler := &DocumentHandler{store: store}
  r.POST("/documents", handler.CreateDocument)
  r.GET("/documents", handler.GetAllDocuments)
  r.GET("/documents/:id", handler.GetDocument)
  r.PUT("/documents/:id", handler.UpdateDocument)
  r.DELETE("/documents/:id", handler.DeleteDocument)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type DocumentHandler struct {
  store DocumentStore
}

type DocumentStore interface {
  Create(doc documents.Document) error
  GetAll() ([]documents.Document, error)
  Get(id string) (documents.Document, error) 
  Update(id string, doc documents.Document) error
  Delete(id string) error
}


func (h *DocumentHandler) CreateDocument(c *gin.Context){
  var doc documents.Document
  if err := c.ShouldBindJSON(&doc); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  h.store.Create(doc)
  c.JSON(http.StatusCreated, doc)
}

func (h *DocumentHandler) GetAllDocuments(c *gin.Context){
  documents, err := h.store.GetAll()
  if err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, documents)
}

func (h *DocumentHandler) GetDocument(c *gin.Context){
  id := c.Param("id")
  document, error := h.store.Get(id)
  if error != nil{
    
    c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
    return
  }
  c.JSON(200, document)
}

func (h *DocumentHandler) UpdateDocument(c *gin.Context){
  var doc documents.Document
  
  if err := c.ShouldBindJSON(&doc); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  id := c.Param("id")
  error := h.store.Update(id, doc)

  if error != nil{
    c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context){
  id := c.Param("id")
  err := h.store.Delete(id)
  if err != nil{
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200,"deleted")
}

