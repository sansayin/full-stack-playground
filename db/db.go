package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Adapter struct {
	db *gorm.DB
}

func NewPostSqlDBAdaptor(dataSourceName string) (*Adapter, error) {
	// connect
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}
	return &Adapter{db: db}, nil
}

func NewSqlitDBAdaptor(dataSourceName string) (*Adapter, error) {
	// connect
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) Get() *gorm.DB {
	return da.db
}

func (da Adapter) CloseDbConnection() {
	log.Printf("db close not required")
}
