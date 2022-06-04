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

func searchKey(searchkey string, searchValue string) data.Task {
	var t data.Task
	for _, t = range list {
		for k, v := range t {
			if k == searchkey && v == searchValue {
				return t
			}
		}
	}
	return nil
}

func GetAll(c *gin.Context) {
	for _, task := range list {
		c.JSON(http.StatusOK, task)
	}
}

func Create(c *gin.Context) {
	title := c.DefaultPostForm("title", "none")
	title_len := utf8.RuneCountInString(title)
	description := c.DefaultPostForm("description", "none")

	if 0 < title_len && title_len < 32 {
		id, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return
		}
		list = append(list, data.Task{"id": id.String(), "title": title, "description": description})
		for _, t := range list {
			c.JSON(200, t)
			c.String(200, "\n")
		}
		c.JSON(http.StatusCreated, "ok")
	} else {
		c.String(http.StatusCreated, "title is very long\n")
	}
}

func GetOne(c *gin.Context) {
	search_id := c.Param("task_id")
	result := searchKey("id", search_id)

	if result == nil {
		c.String(http.StatusBadRequest, "No Task")
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ModifyOne(c *gin.Context) {
	search_id := c.Param("task_id")
	title := c.PostForm("title")
	description := c.PostForm("description")

	flag := false
	for _, t := range list {
		for k, v := range t {
			if k == "id" && v == search_id {
				t["title"] = title
				t["description"] = description
				flag = true
			}
		}
	}

	if !flag {
		c.String(http.StatusBadRequest, "CANNOT")
	} else {
		result := searchKey("id", search_id)
		c.JSON(http.StatusAccepted, result)
	}
}
