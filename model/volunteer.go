package model

import (
	"time"

	"gorm.io/gorm"
)

type VolunteerRole struct {
	gorm.Model `json:"-"`
	Group      string    `json:"group"`
	Role       string    `json:"role"`
	StartYear  time.Time `json:"startYear"`
	EndYear    time.Time `json:"endYear,omitempty"`
}
