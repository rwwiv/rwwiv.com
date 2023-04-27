package model

import "gorm.io/gorm"

type Skill struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
}
