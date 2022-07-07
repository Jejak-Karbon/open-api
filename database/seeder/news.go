package seeder

import (
	"log"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

func newsTableSeeder(conn *gorm.DB) {

	var news = []model.News{}

	for i := 0; i < 15; i++ {
		news = append(news, model.News{Title: "News", Description: "News Descriptiom", Image: "-", IsActive: 1})
	}

	if err := conn.Create(&news).Error; err != nil {
		log.Printf("cannot seed data news, with error %v\n", err)
	}
	log.Println("success seed data news")
}
