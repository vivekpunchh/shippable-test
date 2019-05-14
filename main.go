package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

type ResponseData struct {
	Dataname string `json:"dataname"`
}

type Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

type Genre struct {
	Country string `json:"country"`
	Rock    string `json:"rock"`
}

type Music struct {
	Genre Genre `json:"genre"`
}

func main() {

	app = gin.Default()

	app.GET("/", func(c *gin.Context) {
		//data := ResponseData{Dataname: "vinoth"}

		mydata := Response{Name: "vivek", Age: 99, Sex: "data"}

		data, err := json.Marshal(mydata)

		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(200, string(data))

	})

	// Start serving the application
	app.Run(":5000")
}
