package controllers

import (
	"community-marketplace/config"
	"community-marketplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var input struct {
		ReceiverId uint   `json:"receiver_id" binding:"required"`
		ItemId     uint   `json:"item_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senderId := c.MustGet("userID").(uint)
	role := c.MustGet("role").(string)

	message := models.Message{
		SenderId:         senderId,
		ReceiverId:       input.ReceiverId,
		ItemId:           input.ItemId,
		Content:          input.Content,
		CreatedByLoginId: senderId,
		CreatedByRoleId:  role,
	}

	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, message)
}

func GetConversations(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	// Complex query to get conversations grouped by item and other user
	// For simplicity in this test, we'll just return all messages involving the user
	// and let the frontend group them if needed, or do a simple group by.

	var messages []models.Message
	// Get messages where user is either sender or receiver
	config.DB.Where("sender_id = ? OR receiver_id = ?", userId, userId).
		Preload("Sender").Preload("Receiver").Preload("Item").
		Order("created_at desc").Find(&messages)

	c.JSON(http.StatusOK, messages)
}

func GetThread(c *gin.Context) {
	userId := c.MustGet("userID").(uint)
	otherUserId := c.Param("other_user_id")
	itemId := c.Query("item_id")

	var messages []models.Message
	query := config.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userId, otherUserId, otherUserId, userId)

	if itemId != "" {
		query = query.Where("item_id = ?", itemId)
	}

	query.Preload("Sender").Preload("Receiver").Preload("Item").Order("created_at asc").Find(&messages)

	// Mark messages as read
	config.DB.Model(&models.Message{}).Where("receiver_id = ? AND sender_id = ? AND is_read = false", userId, otherUserId).Update("is_read", true)

	c.JSON(http.StatusOK, messages)
}
