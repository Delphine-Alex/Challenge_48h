package public

import (
	"GoChallenge48/Go/repository"
	"GoChallenge48/Go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetEvents())
}

func PostEvent(c *gin.Context) {
	// Validate input
	var input models.CreateEvent
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(repository.GetEventByName(input.Title)) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "This username is already taken by another user"})
		return
	}

	// Create Wishlist to create a user
	repository.PostEvent(input)

	c.JSON(http.StatusCreated, input)
}
