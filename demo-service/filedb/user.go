package filedb

import (
	"demo/models"
	"os"
)

type UserDB struct {
	FileName string
}

func (u *UserDB) CreateUser(user *models.User) (*models.User, error) {
	file, err := os.OpenFile(u.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := user.ToBytes()
	if err != nil {
		return nil, err
	}

	_, err = file.Write([]byte(string(bytes) + "\n"))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) GetUserBy(id string) (*models.User, error) {
	user := new(models.User)
	return user, nil
}

func (u *UserDB) DeleteUserBy(id string) (int64, error) {

	return 0, nil
}
