package handlers

import (
	"demo/interfaces"
	"demo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	interfaces.IUser // promoted field
}

func (u *Userhandler) Create(c *gin.Context) {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		log.Println("error:", err.Error())
		c.String(http.StatusBadRequest, "invalid input")
		return
	}

	user.LastModified = time.Now().Unix()
	if user.Status == "" {
		user.Status = "active"
	}
	err = user.Validate()
	if err != nil {
		log.Println("error:", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err = u.CreateUser(user)
	if err != nil {
		log.Println("error:", err.Error())
		c.String(http.StatusBadRequest, "failed to create user")
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (u *Userhandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("error:", "no id param value found")
		c.String(http.StatusBadRequest, "no id param value found")
		return
	}

	user, err := u.GetUserBy(id)
	if err != nil {
		log.Println("error:", err.Error())
		c.String(http.StatusBadRequest, "failed to fetch user")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *Userhandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("error:", "no id param value found")
		c.String(http.StatusBadRequest, "no id param value found")
		return
	}
	r, err := u.DeleteUserBy(id)
	if err != nil {
		log.Println("error:", err.Error())
		c.String(http.StatusBadRequest, "failed to delete user or no record to delete")
		return
	}
	c.JSON(http.StatusOK, r)
}