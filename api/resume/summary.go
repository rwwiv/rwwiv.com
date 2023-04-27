package resume

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tech struct {
	gorm.Model
	Value string `json:"value"`
}

type Experience struct {
	gorm.Model
	Value string `json:"value"`
}

type Employment struct {
	gorm.Model
	JobTitle     string       `json:"jobTitle"`
	EmployerName string       `json:"employerName"`
	Location     string       `json:"location"`
	StartDate    time.Time    `json:"startDate"`
	EndDate      time.Time    `json:"endDate,omitempty"`
	Tech         []Tech       `json:"tech"`
	Experience   []Experience `json:"experience"`
}

type VolunteerRole struct {
	gorm.Model
	Group     string    `json:"group"`
	Role      string    `json:"role"`
	StartYear time.Time `json:"startYear"`
	EndYear   time.Time `json:"endYear,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("coul not do")
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	db.AutoMigrate(&Employment{})

	var employments []Employment

	db.Find(&employments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(employments)
	w.Write(json)
}
