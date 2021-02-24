package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", hello)

	router.LoadHTMLGlob("./templates/*")

	router.Run(":8080")
}
