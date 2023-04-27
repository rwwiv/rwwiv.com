package resume

import (
	"encoding/json"
	"net/http"
	"os"
	"rwwiv_com/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	// naive, connection pooling should be investigated
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&model.Tech{})
	db.AutoMigrate(&model.Experience{})
	db.AutoMigrate(&model.Employment{})

	var employment model.Employment
	res := db.Limit(1).Find(&employment)
	if res.RowsAffected == 0 {
		employment = model.Employment{
			JobTitle:     "Engineer",
			EmployerName: "Lightmatter",
			Location:     "Remote",
			StartDate:    time.Now(),
			Tech: []model.Tech{
				{Value: "JS"},
			},
			Experience: []model.Experience{
				{Value: "Did a thing"},
			},
		}

		db.Create(&employment)
		db.Save(&employment)
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Set("Content-Type", "application/json")

	output, _ := json.Marshal(employment)
	w.Write(output)
}
