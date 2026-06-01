package model
import (
  "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name, School, adres string
	age , phone int
	maried bool
}
