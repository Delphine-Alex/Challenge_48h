package public

import (
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrganizer(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetOrganizer())
}