package main

import (
	"demo/filedb"
	"demo/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// DSN := os.Getenv("DB_CONN")
	// if DSN == "" {
	// 	DSN = `admin:admin@tcp(127.0.0.1:3306)/demodb?charset=utf8mb4&parseTime=True&loc=Local`
	// }

	// db, err := database.GetConnection(DSN)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r := gin.Default()

	//uhandler := &handlers.Userhandler{IUser: &database.UserDB{DB: db}}
	uhandler := &handlers.Userhandler{IUser: &filedb.UserDB{FileName: "user.db"}}

	r.POST("/user", uhandler.Create)
	r.GET("/user/:id", uhandler.GetUserByID)
	r.DELETE("/user/:id", uhandler.DeleteUserByID)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}

	// log.Println(db)

	// user := &models.User{Name: "Jiten", Email: "jitenp@outlook.com", Status: "active", LastModified: time.Now().Unix()}

	// udb := new(database.UserDB)
	// udb.DB = db

	// var iuser interfaces.IUser = udb
	// user, err = iuser.CreateUser(user)
	// if err != nil {
	// 	println(err.Error())
	// }
	// fmt.Println(user)

}
