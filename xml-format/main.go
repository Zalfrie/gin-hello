package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstname,attr"`
	LastName  string   `xml:"lastname,attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "zalfrie",
		LastName:  "Januardi",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":4141")
}
