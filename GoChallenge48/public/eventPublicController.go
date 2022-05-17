package public

import (
	//"GoChallenge48/Go/models"
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetEvents())
}