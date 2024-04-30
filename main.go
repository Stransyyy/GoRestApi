package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type resp struct {
	Title string    `json:"title"`
	Msg   string    `json:"msg"`
	TS    time.Time `json:"ts"`
}

var mensajes = []resp{
	{
		Title: "Hello",
		Msg:   "Hello, World!",
		TS:    time.Now(),
	},
	{
		Title: "This is a test of a web service",
		Msg:   "This is a test of a web service",
		TS:    time.Unix(0, 0),
	},
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})

	})

	r.GET("/mensajes", GetMensajes)
	r.POST("/mensajes", postMensaje)

	r.Run("localhost:8080")

}

func GetMensajes(c *gin.Context) {
	c.JSON(http.StatusOK, mensajes)
}

func postMensaje(c *gin.Context) {
	var newMessaje resp

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMessaje); err != nil {
		return
	}

	// Add the new album to the slice.
	mensajes = append(mensajes, newMessaje)
	c.IndentedJSON(http.StatusCreated, newMessaje)
}
