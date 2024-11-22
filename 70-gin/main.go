package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	PORT        string
	fileNameCh  chan string
	fileNameSig chan struct{}
)

func init() {
	fileNameCh = make(chan string)
	fileNameSig = make(chan struct{})
}
func main() {

	//args := os.Args
	PORT = os.Getenv("PORT") // take from environment if not then take can commandline argument
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8080", "--port=8080")
		flag.Parse()
	}
	r := gin.Default()
	//r.Use(Auth)

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		c.String(http.StatusOK, "pong")
	})

	r.POST("/upload", uploadfile)
	r.POST("/uploadname", uploadfileByName())
	r.POST("/user/:id", CreateUser)
	go SendUUID()
	if err := r.Run(":" + PORT); err != nil {
		runtime.Goexit()
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

// func PostProcess(c *gin.Context) {
// 	func() { ch <- string(uuid.New().String()) }()
// 	c.Next()
// }

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

func uploadfile(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	//log.Println(file.Filename)
	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "images/"+file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func uploadfileByName() func(*gin.Context) {
	go func() {
		fileNameSig <- struct{}{}
	}()
	return func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		//log.Println(file.Filename)
		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, "images/"+<-fileNameCh)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

func uploadfiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "/images"+file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func SendUUID() {
	for {
		select {
		case <-fileNameSig:
			fileNameCh <- uuid.New().String()
		}
	}
}
