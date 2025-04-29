package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/roqiaahmed/wikidocify/pkg/models"
	"github.com/roqiaahmed/wikidocify/initializers"
)


func CreateDocument(c *gin.Context){
	var doc models.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	initializers.DB.Create(&doc)
	c.JSON(http.StatusCreated, doc)
}

func GetAllDocuments(c *gin.Context){
	var documents []models.Document
	if err := initializers.DB.Find(&documents).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func GetDocument(c *gin.Context){
	id := c.Param("id")
	var document models.Document
	if err := initializers.DB.First(&document, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, document)
}

func UpdateDocument(c *gin.Context){
	var doc models.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := initializers.DB.Model(&doc).Where("id = ?", id).Updates(doc).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

func DeleteDocument(c *gin.Context){
	id := c.Param("id")
	if err := initializers.DB.Delete(&models.Document{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
