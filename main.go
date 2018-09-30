package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// fs := http.FileServer(http.Dir("./"))
	// http.Handle("/", fs)

	// log.Println("Listening...")
	// http.ListenAndServe(":80", nil)
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./", true)))

	// api := r.Group("/api")
	// {
	// 	api.GET("/events", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// }

	r.Run()
}
