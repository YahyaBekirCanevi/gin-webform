package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type Submission struct {
	FirstName     string
	LastName      string
	PhotoFilename string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/uploads", "./uploads")
	// Route to display the form

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	// Route to handle form submission
	r.POST("/submit", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		// Handling file upload
		file, err := c.FormFile("photo")

		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		// Save the file to the uploads directory
		filename := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		submission := Submission{
			FirstName:     firstName,
			LastName:      lastName,
			PhotoFilename: filename,
		}

		c.HTML(http.StatusOK, "result.html", submission)
	})
	// Run the server
	r.Run(":8080")
}
