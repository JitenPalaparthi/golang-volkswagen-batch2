package database

import (
	"demo/models"
	"errors"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (u *UserDB) CreateUser(user *models.User) (*models.User, error) {
	err := u.DB.AutoMigrate(models.User{})
	if err != nil {
		return nil, err
	}
	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) GetUserBy(id string) (*models.User, error) {
	user := new(models.User)
	tx := u.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) GetUsers(limit, offset int) ([]models.User, error) {
	users := make([]models.User, 0)
	tx := u.DB.Limit(limit).Offset(offset).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// sdb, err := sql.Open("e", "")
	// id := ""
	//sdb.Query("select user.name, orders.name join orders on order.id=user.id where users.id=?", id)
	//	u.DB.Table("users").Select("users.name,users.email,orders.id,orders.name").Joins("orders on orders.userid=users.id").Where("").Scan(&users)
	return users, nil
}

func (u *UserDB) DeleteUserBy(id string) (int64, error) {
	tx := u.DB.Delete(models.User{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("no record to delete")
	}
	return tx.RowsAffected, nil
}
