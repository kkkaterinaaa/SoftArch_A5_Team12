package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
	"twitter/interction-service/initializers"

	"twitter/interction-service/models"
)

func LikeMessage(c *gin.Context) {
	var input struct {
		UserID    string `json:"user_id"`
		MessageID uint   `json:"message_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Invalid input for liking message")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.UserID == "" {
		logrus.Warn("Unauthorized attempt to like a message")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must be logged in to like a message"})
		return
	}

	userId, _ := strconv.Atoi(input.UserID)
	userID := uint(userId)

	var existingLike models.Like
	if err := initializers.DB.Where("user_id = ? AND message_id = ?", userID, input.MessageID).First(&existingLike).Error; err == nil {
		if err := initializers.DB.Delete(&existingLike).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"error":      err.Error(),
				"user_id":    userID,
				"message_id": input.MessageID,
			}).Error("Failed to unlike message")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to unlike message"})
			return
		}

		logrus.WithFields(logrus.Fields{
			"user_id":    userID,
			"message_id": input.MessageID,
		}).Info("Message successfully unliked")

		c.JSON(http.StatusOK, gin.H{"message": "Unliked"})
		return
	}

	newLike := models.Like{
		UserID:    userID,
		MessageID: input.MessageID,
		CreatedAt: time.Now(),
	}

	if err := initializers.DB.Create(&newLike).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error":      err.Error(),
			"user_id":    userID,
			"message_id": input.MessageID,
		}).Error("Failed to like message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to like message"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    userID,
		"message_id": input.MessageID,
	}).Info("Message successfully liked")

	c.JSON(http.StatusCreated, gin.H{"message": "Liked"})
}

func GetMessageLikes(c *gin.Context) {
	messageID, _ := strconv.Atoi(c.Param("messageID"))

	var count int64
	if err := initializers.DB.Model(&models.Like{}).Where("message_id = ?", messageID).Count(&count).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error":      err.Error(),
			"message_id": messageID,
		}).Error("Failed to retrieve likes for message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve likes"})
		return
	}

	liked := false

	userID := c.Query("user_id")
	if userID != "" {
		userIDint, err := strconv.Atoi(userID)
		if err == nil {
			userIDu := uint(userIDint)
			var existingLike models.Like
			if err := initializers.DB.Where("user_id = ? AND message_id = ?", userIDu, messageID).First(&existingLike).Error; err == nil {
				liked = true
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"error":   err.Error(),
				"user_id": userID,
			}).Warn("Invalid user_id provided")
		}
	}

	logrus.WithFields(logrus.Fields{
		"message_id": messageID,
		"likes":      count,
		"liked":      liked,
	}).Info("Retrieved like count for message")

	c.JSON(http.StatusOK, gin.H{
		"message_id": messageID,
		"likes":      count,
		"liked":      liked,
	})
}
