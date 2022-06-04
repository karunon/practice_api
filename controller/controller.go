package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"practice_api/data"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// var list []data.Task
var list []data.Task

func Initial() {
	// list_1 := data.Task{
	// 	Id:          "22-22",
	// 	Title:       "SSS",
	// 	Description: "xxx",
	// }

	// list_2 := data.Task{
	// 	Id:          "33-33",
	// 	Title:       "AAA",
	// 	Description: "YYY",
	// }

	// list = append(list, list_1, list_2)
}

// func searchKey(searchValue string) *data.Task {
// 	var t data.Task
// 	for _, t = range list {
// 		// for k, v := range t {
// 		// 	if k == searchkey && v == searchValue {
// 		// 		return &t
// 		// 	}
// 		// }
// 		if t.Id == searchValue {
// 			return &t
// 		}
// 	}
// 	var st data.Task
// 	return &st
// }

func GetAll(c *gin.Context) {
	jsonData, _ := json.Marshal(list)
	c.JSON(http.StatusAccepted, string(jsonData))
}

func Create(c *gin.Context) {
	// title := c.DefaultPostForm("title", "none")
	// title_len := utf8.RuneCountInString(title)
	// description := c.DefaultPostForm("description", "none")

	var p_json data.PostJsonRequest
	if err := c.ShouldBindJSON((&p_json)); err != nil {
		c.String(http.StatusBadRequest, "shouldBindJSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	title_len := utf8.RuneCountInString(p_json.Title)
	c.String(200, p_json.Title)

	if 0 < title_len && title_len < 32 {
		id, err := uuid.NewRandom()
		if err != nil {
			c.String(http.StatusBadRequest, "UUIDerror")
			log.Println(err)
			return
		}
		task := data.Task{
			Id:          id.String(),
			Title:       p_json.Title,
			Description: p_json.Description,
		}
		list = append(list, task)

		json_data, er := json.Marshal(task)
		if er != nil {
			c.String(http.StatusBadRequest, "Maeshal error")
			log.Println(er)
			return
		} else {
			c.JSON(http.StatusCreated, string(json_data))
		}
	} else {
		c.String(http.StatusCreated, "title is very long\n")
	}
}

func GetOne(c *gin.Context) {
	search_id := c.PostForm("task_id")
	// result := searchKey(search_id)

	flag := false
	for _, t := range list {
		c.String(200, "%s this is a ID", t.Id)
		if t.Id == search_id {
			c.String(200, "check")
			c.JSON(http.StatusOK, t)
			flag = true
			break
		}
	}
	if !flag {
		c.String(http.StatusBadRequest, "cannot find")
	}

	// if result == nil {
	// 	c.String(http.StatusBadRequest, "No Task")
	// } else {
	// 	jsonData, err := json.Marshal((result))
	// 	if err != nil {
	// 		c.JSON(http.StatusOK, jsonData)
	// 	}
	// }
}

func ModifyOne(c *gin.Context) {
	// search_id := c.Param("task_id")
	// title := c.PostForm("title")
	// description := c.PostForm("description")

	// flag := false
	// for _, t := range list {
	// 	for k, v := range t {
	// 		if k == "id" && v == search_id {
	// 			t["title"] = title
	// 			t["description"] = description
	// 			flag = true
	// 		}
	// 	}
	// }

	// if !flag {
	// 	c.String(http.StatusBadRequest, "CANNOT")
	// } else {
	// 	result := searchKey("id", search_id)
	// 	c.JSON(http.StatusAccepted, result)
	// }
}
