package resume

import (
	"encoding/json"
	"log"
	"net/http"

	"rwwiv_com/api"
	models "rwwiv_com/api/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	pcConn := api.PgConn{}
	db, err := pcConn.GetConnection()

	if err != nil {
		log.Fatalf("coul not do")
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	db.AutoMigrate(&models.Employment{})

	var employments []models.Employment

	db.Find(&employments)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(employments)
	w.Write(json)
}
