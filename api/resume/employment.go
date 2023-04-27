package resume

import (
	"encoding/json"
	"net/http"
	"os"
	"rwwiv_com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.Limit(1).Preload(clause.Associations).Find(&employment)

	w.Header().Set("Content-Type", "application/json")

	output, _ := json.Marshal(employment)
	w.Write(output)
}
