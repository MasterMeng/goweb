package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.GET("/favicon.ico", handleFavicon)
	router.GET("/playground", content)
	router.POST("",handlePost)

	router.LoadHTMLGlob("./templates/*")
	router.Static("/img", "./img")
	router.Static("/static", "./static")

	router.Run(":8080")
}
