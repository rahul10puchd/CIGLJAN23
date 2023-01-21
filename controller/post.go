package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//acting as db for as
var post = map[string]string{
	"rahul": "Present",
}

func GetPost(c *gin.Context) {
	res := gin.H{
		"message": "post",
	}
	c.JSON(http.StatusUnauthorized, res)
}
func DeletePost(c *gin.Context) {

	key := c.Params.ByName("id")
	//key := c.Query("id")
	//to delete the key value pair from a map via key
	delete(post, key)
	res := gin.H{
		"message": key + " Deleted",
	}
	c.JSON(http.StatusOK, res)
}

type PostPayload struct {
	Name string `json:"name"`
}

func CreatePost(c *gin.Context) {
	//log.Println("*********", c, "*********")
	var p1 PostPayload
	err := c.ShouldBindJSON(&p1)
	if err != nil {
		log.Print("got error", err)
	}
	post[p1.Name] = "Present"
	res := gin.H{
		"message": p1,
	}
	c.JSON(http.StatusOK, res)
}
