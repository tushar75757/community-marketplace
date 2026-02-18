package controllers

import (
	"community-marketplace/config"
	"community-marketplace/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	query := config.DB.Preload("Category").Preload("Seller")

	// Filters
	categoryID := c.Query("category_id")
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	minPrice := c.Query("min_price")
	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	maxPrice := c.Query("max_price")
	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	search := c.Query("search")
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var total int64
	query.Model(&models.Item{}).Count(&total)

	if err := query.Limit(pageSize).Offset(offset).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      items,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.Preload("Category").Preload("Seller").First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var input struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		CategoryId  uint    `json:"category_id" binding:"required"`
		Condition   string  `json:"condition" binding:"required"`
		ImageUrl    string  `json:"image_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userID").(uint)
	role := c.MustGet("role").(string)

	item := models.Item{
		Title:            input.Title,
		Description:      input.Description,
		Price:            input.Price,
		CategoryId:       input.CategoryId,
		Condition:        input.Condition,
		ImageUrl:         input.ImageUrl,
		SellerId:         userId,
		CreatedByLoginId: userId,
		CreatedByRoleId:  role,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.MustGet("userID").(uint)
	role := c.MustGet("role").(string)

	// Check ownership or admin
	if item.SellerId != userId && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this item"})
		return
	}

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		CategoryId  uint    `json:"category_id"`
		Condition   string  `json:"condition"`
		ImageUrl    string  `json:"image_url"`
		Status      string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.MustGet("userID").(uint)
	role := c.MustGet("role").(string)

	if item.SellerId != userId && role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this item"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func GetMyListings(c *gin.Context) {
	userId := c.MustGet("userID").(uint)
	var items []models.Item
	config.DB.Where("seller_id = ?", userId).Preload("Category").Find(&items)
	c.JSON(http.StatusOK, items)
}
