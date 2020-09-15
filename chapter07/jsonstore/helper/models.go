package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	Packages string // Package
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func prepareDBVars() func () (*gorm.DB, error) {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "chapter07/jsonstore/.env")
	_ = godotenv.Load(path)

	host, port := os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT")
	user, password := os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_PWD")
	dbname := os.Getenv("POSTGRESQL_DB")

	return func () (*gorm.DB, error) {
		var err error
		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
		db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

		if err != nil {
			return nil, err
		}

		db.AutoMigrate(&Shipment{}, &Package{})
		return db, nil
	}
}

var InitDB = prepareDBVars()

