package models

import "github.com/jinzhu/gorm"

type BaseModel gorm.Model

var ErrRecordNotFound = gorm.ErrRecordNotFound
