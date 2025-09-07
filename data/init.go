package data

import (
	"Web/data/db"
	"Web/model"
	"log"

	"gorm.io/gorm"
)

func Up_1() {

	database := db.GetDb()

	err := database.AutoMigrate(
		&model.City{},
		&model.Description{},
	)
	if err != nil {
		log.Println("migrate failed")
	}

	CreateCity(database)

}

func CreateCity(database *gorm.DB) {

	var count int64
	if err := database.Model(&model.City{}).Count(&count).Error; err != nil {

		log.Printf("error counting cities: %v", err)
		return
	}

	if count == 0 {
		cities := []model.City{
			{Name: "آذربایجان شرقی"},
			{Name: "آذربایجان غربی"},
			{Name: "اردبیل"},
			{Name: "اصفهان"},
			{Name: "البرز"},
			{Name: "ایلام"},
			{Name: "بوشهر"},

		}

		if err := database.Create(&cities).Error; err != nil {

			log.Println("error Creating cities: %v", err)

		} else {

			log.Println("cities inserted successfully")
		}
	}
}
