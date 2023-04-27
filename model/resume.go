package model

import (
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type Tech struct {
	gorm.Model   `json:"-"`
	Value        string `json:"value"`
	EmploymentID uint   `json:"-"`
}

func (Tech) TableName() string {
	return "techs"
}

type Experience struct {
	gorm.Model   `json:"-"`
	Value        string `json:"value"`
	EmploymentID uint   `json:"-"`
}

type Employment struct {
	gorm.Model   `json:"-"`
	JobTitle     string       `json:"jobTitle"`
	EmployerName string       `json:"employerName"`
	Location     string       `json:"location"`
	StartDate    time.Time    `json:"startDate"`
	EndDate      time.Time    `json:"endDate,omitempty"`
	Tech         []Tech       `json:"tech"`
	Experience   []Experience `json:"experience"`
}
