package main

import (
	"github.com/gin-gonic/gin"
)

type Coffee struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/", get)
	router.POST("/", post)
	router.PUT("/", put)
	router.DELETE("/", delete)
	router.Run(":8000")
}
