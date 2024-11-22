package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	//gorm.Model
	Id           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("invalid email or email field not provided")
	}
	if u.Name == "" {
		return errors.New("invalid name or name field not provided")
	}

	return nil
}

func (u *User) ToBytes() ([]byte, error) {
	return json.Marshal(u)
}
