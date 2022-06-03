package controller

import (
	"net/http"
	"practice_api/data"

	"github.com/gin-gonic/gin"
)

//var list = []data.Task{{"1": "a", "3": "4"}, {"2": "b"}}
var list []data.Task

func GetAll(c *gin.Context) {
	for _, task := range list {
		c.JSON(http.StatusOK, task)
	}
}

func Create(c *gin.Context) {
}

func GetOne(c *gin.Context) {
}

func ModifyOne(c *gin.Context) {
}
