package database

import (
	"demo/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, error) {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		return nil, err
	}
	tx := db.Create(user)
	if tx.Error != nil {
		return nil, err
	}
	return user, nil
}

// func (u *UserDB) GetUserBy(id string) (*models.User, error) {
// 	user := new(models.User)
// 	tx := u.DB.First(user, id)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return user, nil
// }

// func (u *UserDB) DeleteUserBy(id string) (int64, error) {
// 	tx := u.DB.Delete(models.User{}, id)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return 0, errors.New("no record to delete")
// 	}
// 	return tx.RowsAffected, nil
// }
