package public

import (
	//"GoChallenge48/Go/models"
	"GoChallenge48/Go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetCustomers())
}