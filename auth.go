package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

//GetToken return pseudo-random authentication token
func GetToken() (string, error) {
	Hash := sha256.New()

	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	io.WriteString(Hash, string(b))

	// The slice should now contain random bytes instead of only zeroes.
	return base64.StdEncoding.EncodeToString(Hash.Sum(nil)), nil
}

//NewAuthSession creates new or replaces Session
func NewAuthSession(m discordgo.Message) (string, string, error) {
	//Connect DB
	var Message string

	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//Check for old Session and destroy
	DBSession := new(APISession)
	db.Where("user_id = ?", m.Author.ID).Find(&DBSession)
	log.Println(DBSession)
	if DBSession.UserID != "" {
		db.Delete(&DBSession)
		Message = "Found old token, removed!"
	}

	//Generate New Token
	token, err := GetToken()
	if err != nil {
		Message = "Error, getting token!"
		return "", Message, err
	}

	//Store New Token
	AuthSession := new(APISession)
	AuthSession.Token = token
	AuthSession.UserID = m.Author.ID
	AuthSession.Time = time.Now()

	db.Create(&AuthSession)

	db.Close()

	return AuthSession.Token, Message, nil

}
