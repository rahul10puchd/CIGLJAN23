package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostPayload struct {
	Name string `json:"name"`
}

func main() {
	// post := map[string]string{
	// 	"rahul": "jkghvkjhf",
	// }
	// r := gin.Default() //application will be configured on port 8080
	// pingFunc := func(c *gin.Context) {

	// 	key := map[string]any{
	// 		"ping": "pingu",
	// 	}
	// 	c.Keys = key
	// 	res := gin.H{
	// 		"message":      "pong",
	// 		"contextvalue": c.Keys["ping"],
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// pongFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message":      "pong",
	// 		"contextvalue": c.Keys["ping"],
	// 	}
	// 	log.Println(c.Keys["ping"])
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// dingFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message": "pong",
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	// dongFunc := func(c *gin.Context) {
	// 	res := gin.H{
	// 		"message": "pong",
	// 	}
	// 	c.JSON(http.StatusUnauthorized, res)
	// }
	post := map[string]string{
		"rahul": "Present",
	}
	r := gin.Default()
	getPost := func(c *gin.Context) {
		res := gin.H{
			"message": post,
		}
		c.JSON(http.StatusUnauthorized, res)
	}

	// r.GET("/ping", pingFunc) //http://localhost:8080/ping
	// r.GET("/pong", pongFunc)
	// r.GET("/ding", dingFunc)
	// r.GET("/dong", dongFunc)
	r.GET("/post", getPost)
	createPost := func(c *gin.Context) {
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
	r.POST("/post", createPost)
	deletePost := func(c *gin.Context) {

		//key := c.Params.ByName("id")
		key := c.Query("id")
		//to delete the key value pair from a map via key
		delete(post, key)
		res := gin.H{
			"message": key + " Deleted",
		}
		c.JSON(http.StatusOK, res)
	}
	//r.DELETE("/post/:id", deletePost)
	r.DELETE("/post", deletePost)
	//8:17
	// r.PUT("/post", updatePost)       //data will be sent via payload and query
	// r.PATCH("/post", updateNamebyId) //data will be sent via query
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}

/// delete post by id:-
// url: http://localhost:8080/post?id=rahul
// url: http://localhost:8080/post/rahul
// request method: delete

// get post by id:-
// url: http://localhost:8080/post?id=1
// request method: get

//http://localhost:8080/post
