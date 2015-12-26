package main

import (
	"time"
)

//Picture is a default DB structure for GORM
type Picture struct {
	ID         int       `json:"id"`
	Title      string    `json:"title" binding:"required"`
	URL        string    `json:"url" binding:"required"`
	CreateTime time.Time `json:"ctime"`

	Tags []Tag `gorm:"many2many:picture_tags;" json:"tags"`
}

//Tag is a default DB structure for GORM
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`

	Pictures []Picture `gorm:"many2many:picture_tags;" json:"-"`
}

//APISession handles authentication
type APISession struct {
	ID     int       `json:"id"`
	Token  string    `json:"token" binding:"required"`
	UserID string    `json:"user"`
	Time   time.Time `json:"time"`
}

//Config struct that contains needed information for GoGerard
type Config struct {
	Username     string
	Password     string
	ServerID     string
	TestServerID string
	VoiceID      string
	SFWChannel   string
	Admin        string
}
