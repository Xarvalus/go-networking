package models

import (
	"github.com/jinzhu/gorm"
)

type Children struct {
	gorm.Model
	Boolean  bool
	ParentID uint
}
