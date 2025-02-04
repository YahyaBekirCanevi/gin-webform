package main

import (
	"fmt"
	"log"
	"net/http"

	//"path/filepath"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	//versioned := r.Group("/v1/")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Success")
	})
	// Run the server
	err := r.Run(":8080")

	if err != nil {
		log.Fatalln("failed to run Gin app")
	}
}
