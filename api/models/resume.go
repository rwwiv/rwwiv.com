package model

import (
	"time"

	"gorm.io/gorm"
)

type Employment struct {
	gorm.Model
	JobTitle     string    `json:"jobTitle"`
	EmployerName string    `json:"employerName"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate,omitempty"`
	Tech         []string  `json:"tech"`
	Experience   []string  `json:"experience"`
}

type VolunteerRole struct {
	gorm.Model
	Group     string    `json:"group"`
	Role      string    `json:"role"`
	StartYear time.Time `json:"startYear"`
	EndYear   time.Time `json:"endYear,omitempty"`
}
