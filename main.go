package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("resources/views/*.html")

	router.GET("/", mainpage)
	router.GET("/hello", helloWorld)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")

}

func mainpage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
