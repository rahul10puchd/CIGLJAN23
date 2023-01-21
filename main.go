package main

import (
	"gin/basic/db"
	"gin/basic/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_, err := db.GetPostgresDBConnection()
	if err != nil {
		log.Fatal("Unable to establish DB connection", err)
	} else {
		log.Panicln("DB connection establsihed")
	}

	router.PostRouter(r)
	router.UserRouter(r)
	r.Run()

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

	// r := gin.Default()

	// router.Router(context.Background(), r)
	// r.Run()
	// r.GET("/ping", pingFunc) //http://localhost:8080/ping
	// r.GET("/pong", pongFunc)
	// r.GET("/ding", dingFunc)
	// r.GET("/dong", dongFunc)
	//r.GET("/post", getPost)

	// r.POST("/post", createPost)

	//r.DELETE("/post/:id", deletePost)
	// r.DELETE("/post", deletePost)
	//8:17
	// r.PUT("/post", updatePost)       //data will be sent via payload and query
	// r.PATCH("/post", updateNamebyId) //data will be sent via query
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

// backend app: language, docker, kubes, git, database, cloud, ci/cd(github workflow)
// adv: libs for diffrent(store pass-hashicorp vault)
/// delete post by id:-
// url: http://localhost:8080/post?id=rahul
// url: http://localhost:8080/post/rahul
// request method: delete

// get post by id:-
// url: http://localhost:8080/post?id=1
// request method: get

//http://localhost:8080/post
