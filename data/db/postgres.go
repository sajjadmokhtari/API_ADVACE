package db  

import (
	"log" 

	"gorm.io/driver/postgres"  
	"gorm.io/gorm"              
)

var DB *gorm.DB 

func InitDb() error {  
	
	dsn := "host=app_postgres user=postgres password=admin dbname=app_db port=5432 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		
		log.Println("❌ failed to connect to database:", err)
		return err
	}

	
	log.Println("✅ connected to database successfully")

	
	DB = db
	return nil
}


func GetDb() *gorm.DB {
	return DB
}
