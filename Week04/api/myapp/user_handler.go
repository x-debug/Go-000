package myapp

import (
	"github.com/gin-gonic/gin"
	"github.com/x-debug/Go-000/Week04/internal"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userCase := internal.GetUserCase()
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Id is required")
	}
	user, err := userCase.FindById(uid)
	if err != nil {
		c.String(http.StatusNotFound, "Row is not found")
	}

	c.JSON(http.StatusOK, user)
}