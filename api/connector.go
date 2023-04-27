package api

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConn struct{}

func (p *PgConn) GetConnection() (db *gorm.DB, err error) {
	dsn := os.Getenv("DB_CONN")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
