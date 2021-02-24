package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/playground",content)

	router.LoadHTMLGlob("./templates/*")
	router.Static("/img","./img")

	router.Run(":8080")
}
