package handler

import (
	"fmt"
	"log"
	"net/http"
	"newproject/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")

	fmt.Println("///////////", userId)

	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		log.Println("could not convert the string")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "could not process the request",
		})
		return
	}
	userData, err := models.FetchUser(uid)
	if err != nil {
		log.Println("no data found in the map")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Error": "no user data found",
		})
		return
	}
	c.JSON(http.StatusOK, userData)

}
