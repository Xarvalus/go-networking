package models

import (
	"github.com/jinzhu/gorm"
)

type Enum string
const (
	One Enum = "ONE"
	Two Enum = "TWO"
	Three Enum = "THREE"
)

type Parent struct {
	gorm.Model
	Number   uint
	Text     string
	Real     float64
	Enum     Enum
	Children []Children
}
