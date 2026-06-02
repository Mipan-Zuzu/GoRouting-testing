package model
import (
  "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name, School, Adres string
	Age , Phone int
	Maried bool
}
