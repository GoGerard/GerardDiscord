package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//InitDB inits the database
func InitDB() (err error) {
	dbstring := fmt.Sprintf("host=postgres user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD"))
	db, err := gorm.Open("postgres", dbstring)

	db.DB()
	db.LogMode(true)
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

	db.AutoMigrate(&Picture{}, &Tag{}, &APISession{})

	var count int
	db.Table("picture").Count(&count)

	if count == 0 {
		log.Println("No Pictures found, adding MockData")
		MockData(&db)
	}

	return err
}

//ConnectDB is a handle that retrieves *gorm.DB
func ConnectDB() (dbout *gorm.DB, err error) {
	dbstring := fmt.Sprintf("host=postgres user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD"))
	db, err := gorm.Open("postgres", dbstring)

	db.DB()
	db.LogMode(true)
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)

	return &db, err
}

//MockData creates two instants of pictures with 4 tags
func MockData(db *gorm.DB) {
	baby := Tag{
		Name: "Baby",
	}

	cat := Tag{
		Name: "Cat",
	}

	gif := Tag{
		Name: "Gif",
	}

	shit := Tag{
		Name: "Shit",
	}

	picture := Picture{
		Title:      "Dancing Baby Gif",
		URL:        "https://45.media.tumblr.com/07106512835c07d2a237a4d12479c7f8/tumblr_mpg5arui7E1s1clzao1_250.gif",
		CreateTime: time.Now(),
		Tags:       []Tag{},
	}

	picture2 := Picture{
		Title:      "Dancing Cat Gif",
		URL:        "http://rs143.pbsrc.com/albums/r146/sconti1369/Funny_Pictures_Animated_Dancing_Cat.gif~c200",
		CreateTime: time.Now(),
		Tags:       []Tag{},
	}

	db.Create(&picture)
	db.Create(&picture2)
	db.Create(&baby)
	db.Create(&cat)
	db.Create(&shit)
	db.Create(&gif)

	db.Model(&picture).Association("Tags").Append(&baby, &gif, &shit)
	db.Model(&picture2).Association("Tags").Append(&cat, &gif, &shit)

}
