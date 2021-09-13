package rajaongkir

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	URLCheckOngkir  = "https://pro.rajaongkir.com/api/cost"
	URLCheckWaybill = "https://pro.rajaongkir.com/api/waybill"
)

var (
	APIKey string
)

func init() {
	godotenv.Load()

	APIKey = os.Getenv("RAJAONGKIR_API_KEY")
}

func InitDB() (db *sql.DB) {
	godotenv.Load()

	hostDB := os.Getenv("DB_HOST")
	userDB := os.Getenv("DB_USER")
	passDB := os.Getenv("DB_PASSWORD")
	schemaDB := os.Getenv("DB_SCHEMA")

	db, err := sql.Open("mysql", userDB+":"+passDB+"@tcp("+hostDB+")/"+schemaDB+"?parseTime=true")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	return
}
