package controller

import (
	"log"
	"net/http"
	"practice_api/data"
	"sort"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var list []data.Task

func sortList() {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
}

func lookTaskIntoList(s string) int {
	length := len(list)
	i := sort.Search(length, func(i int) bool { return list[i].Id >= s })
	return i
}

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

		sortList()
	} else {
		c.String(http.StatusCreated, "title is very long\n")
	}
}

func GetOne(c *gin.Context) {
	search_id := c.Param("task_id")
	i := lookTaskIntoList(search_id)

	if list[i].Id == search_id {
		c.JSON(http.StatusAccepted, list[i])
	} else {
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

	i := lookTaskIntoList(search_id)
	if list[i].Id == search_id {
		list[i].Title = p_json.Title
		list[i].Description = p_json.Description
		c.JSON(http.StatusAccepted, list[i])
	} else {
		c.String(http.StatusBadRequest, "CANNOT")
	}
}
