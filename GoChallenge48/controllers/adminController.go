package controllers

import (
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNewsletters handle /newsletters (GET)
func GetNewsletters(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetAdmin())
}