package main

import (
	"practice_api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	{
		r.GET("/tasks", controller.GetAll)
		r.POST("tasks", controller.Create)

		r.GET("/tasks/:task_id", controller.GetOne)
		r.PATCH("tasks/:task_id", controller.ModifyOne)
	}

	r.Run(":8080")
}
