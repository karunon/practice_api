package controller

import (
	"log"
	"net/http"
	"practice_api/data"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var list []data.Task

func GetAll(c *gin.Context) {
	c.JSON(http.StatusAccepted, list)
}

func Create(c *gin.Context) {
	var p_json data.PostJsonRequest
	if err := c.ShouldBindJSON((&p_json)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	title_len := utf8.RuneCountInString(p_json.Title)

	if 0 < title_len && title_len < 32 {
		id, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return
		}
		task := data.Task{
			Id:          id.String(),
			Title:       p_json.Title,
			Description: p_json.Description,
		}
		list = append(list, task)

		c.JSON(http.StatusAccepted, task)
	} else {
		c.String(http.StatusCreated, "title is very long\n")
	}
}

func GetOne(c *gin.Context) {
	search_id := c.Param("task_id")

	flag := false
	for _, t := range list {
		if t.Id == search_id {
			c.JSON(http.StatusOK, t)
			flag = true
			return
		}
	}
	if !flag {
		c.String(http.StatusBadRequest, "cannot find")
	}
}

func ModifyOne(c *gin.Context) {
	search_id := c.Param("task_id")

	var p_json data.PostJsonRequest
	if err := c.ShouldBindJSON((&p_json)); err != nil {
		c.String(http.StatusBadRequest, "shouldBindJSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag := false
	for num, t := range list {
		if t.Id == search_id {
			list[num].Title = p_json.Title
			list[num].Description = p_json.Description
			flag = true
			c.JSON(http.StatusAccepted, list[num])
			return
		}
	}

	if !flag {
		c.String(http.StatusBadRequest, "CANNOT")
	}
}
