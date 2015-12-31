package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

//HappyNewYear brings joy
func HappyNewYear(db *gorm.DB) {
	HNY := Tag{
		Name: "HNY",
	}

	picture := Picture{
		Title:      "FireWorks1",
		URL:        "https://media.giphy.com/media/IjmMzurYulKEw/giphy.gif",
		CreateTime: time.Now(),
		Tags:       []Tag{},
	}

	picture2 := Picture{
		Title:      "FireWorks2",
		URL:        "http://bestanimations.com/Holidays/Fireworks/fireworks-animation-46.gif",
		CreateTime: time.Now(),
		Tags:       []Tag{},
	}

	picture3 := Picture{
		Title:      "FireWorks3",
		URL:        "http://media0.giphy.com/media/Qh5dZDCFqr1dK/giphy.gif",
		CreateTime: time.Now(),
		Tags:       []Tag{},
	}

	db.Create(&picture)
	db.Create(&picture2)
	db.Create(&picture3)
	db.Create(&HNY)

	db.Model(&picture).Association("Tags").Append(&HNY)
	db.Model(&picture2).Association("Tags").Append(&HNY)
	db.Model(&picture3).Association("Tags").Append(&HNY)

}
