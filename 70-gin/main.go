package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func main() {
	//args := os.Args
	PORT = os.Getenv("PORT") // take from environment if not then take can commandline argument
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8080", "--port=8080")
		flag.Parse()
	}
	r := gin.Default()
	r.Use(Auth)

	r.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		c.String(http.StatusOK, "pong")
	})

	r.POST("/user/:id", CreateUser, PostProcess)

	if err := r.Run(":" + PORT); err != nil {
		Fatal(err)
	}
}

func Auth(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Next()
}

func PostProcess(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	if id != "" {
		user.Id = id
	}
	if status != "" {
		user.Status = status
	} else {
		user.Status = "inactive"
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	err = WriteToFile("data.txt", bytes)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, user)
	c.Next()

}

func Fatal(elems ...any) {
	fmt.Println(elems...)
	os.Exit(1)
}

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func WriteToFile(fileName string, data []byte) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
