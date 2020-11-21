package model

import (
	"github.com/jinzhu/gorm"
)

const (
	READY = iota
	DOING
	DONE
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type Task struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	State   int    `json:"state" gorm:"not null"`
	OwnerId uint   `json:"owner_id" gorm:"not null"`
}
