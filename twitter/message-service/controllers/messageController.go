package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
	"twitter/message-service/initializers"
	"twitter/message-service/models"
)

func GetFeed(c *gin.Context) {
	var messages []models.Message

	logrus.Info("Fetching the last 10 messages")

	if err := initializers.DB.Order("created_at desc").Limit(10).Find(&messages).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to retrieve messages from the database")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve messages"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"message_count": len(messages),
		"messages":      messages,
	}).Info("Successfully retrieved messages")

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func CreateMessage(c *gin.Context) {

	var input struct {
		Content string `json:"content" binding:"required,max=400"`
		UserID  string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Invalid input for creating message")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.UserID == "" {
		logrus.Warn("Attempt to post a message without being logged in")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to post a message"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"UserID":  input.UserID,
		"content": input.Content,
	}).Info("Attempting to create a new message")

	userId, _ := strconv.Atoi(input.UserID)
	message := models.Message{
		UserID:    uint(userId),
		Content:   input.Content,
		CreatedAt: time.Now(),
	}

	if err := initializers.DB.Create(&message).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create message in the database")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to post message"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"message_id": message.ID,
		"user_id":    message.UserID,
	}).Info("Message successfully created")

	c.JSON(http.StatusCreated, gin.H{"message": message})
}
