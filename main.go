package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

//Session contains events and points to function
var Session = discordgo.Session{
	OnMessageCreate: OnMessageCreate,
	OnReady:         OnReady,
}

//Server contains default GoGerard data
var Server = Config{}

var err error

func main() {
	log.Println("* Parsing Config file")

	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalln("conf.json not found")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Server)
	if err != nil {
		log.Fatalln("Something wrong with JSON file")
	}

	log.Println("* Initialize DB")

	err = InitDB()
	if err != nil {
		log.Println(err)
	}
	//Running Server
	RunServer()
}
