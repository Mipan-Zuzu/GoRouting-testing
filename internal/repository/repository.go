package repository
import (
	"gorm.io/gorm"
	"goauth/internal/model"
)

func GetUser () {
	
}

func InsertUser (db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

