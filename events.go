package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var set int

//OnMessageCreate gets called when a new message occurs in the Session
func OnMessageCreate(s *discordgo.Session, m discordgo.Message) {
	log.Printf("[%5s]: %5s > %s\n", m.ChannelID, m.Author.Username, m.Content)

	if strings.HasPrefix(m.Content, "!go") {
		out, err := exec.Command("go", "version").Output()
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			s.ChannelMessageSend(m.ChannelID, string(out))
		}
	}

	if strings.HasPrefix(m.Content, "!timer") {
		var second = strings.Split(m.Content, " ")
		if len(second) != 1 {
			s.ChannelMessageSend(m.ChannelID, "Ok, <@"+m.Author.ID+">, timer for "+second[1]+" minutes!")

			n, err := strconv.Atoi(second[1])
			if err != nil {
				log.Printf("%s", err)
			}

			timer1 := time.NewTimer(time.Minute * time.Duration(n))
			<-timer1.C
			fmt.Println("Timer 1 expired")
			s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">, your timer has ended!")
		}

	}

	if strings.HasPrefix(m.Content, "!rain") {
		db, err := ConnectDB()
		if err != nil {
			log.Fatal(err)
		}

		if m.ChannelID == Server.SFWChannel { //SFW Check
			s.ChannelMessageSend(m.ChannelID, "This is a SFW channel!")
			return
		}

		var Pictures = []Picture{}
		db.Find(&Pictures)
		for i := 0; i < len(Pictures); i++ {
			s.ChannelMessageSend(m.ChannelID, Pictures[i].URL)
		}

		db.Close()
	}

	if strings.HasPrefix(m.Content, "!picture") {
		db, err := ConnectDB()
		if err != nil {
			log.Fatal(err)
		}

		if m.ChannelID == Server.SFWChannel { //SFW Check
			s.ChannelMessageSend(m.ChannelID, "This is a SFW channel!")
			return
		}
		var param = strings.Split(m.Content, " ")
		var Tag = []Tag{}
		var Pictures = []Picture{}

		if len(param) == 1 { //Return random from all
			db.Find(&Pictures)
			s.ChannelMessageSend(m.ChannelID, Pictures[rand.Intn(len(Pictures))].URL)
			db.Close()
			return
		}

		if len(param) == 2 {
			db.Preload("Pictures").Find(&Tag, "name = ?", param[1])
			if len(Tag) != 0 {
				s.ChannelMessageSend(m.ChannelID, Tag[0].Pictures[rand.Intn(len(Tag[0].Pictures))].URL) //Nice
			}
			db.Close()
			return
		}

		if len(param) == 3 {
			db.Preload("Pictures").Find(&Tag, "name = ?", param[1])
			if len(Tag) != 0 {
				for i := 0; i < len(Tag[0].Pictures); i++ {
					s.ChannelMessageSend(m.ChannelID, Tag[0].Pictures[i].URL)
				}
			}
			db.Close()
		}

	}

	if strings.HasPrefix(m.Content, "!token") {
		token, message, err := NewAuthSession(m)

		if message != "" {
			s.ChannelMessageSend(m.ChannelID, message)
		}

		if err != nil {
			log.Println(err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, token)

	}

	if strings.HasPrefix(m.Content, "!tags") {
		db, err := ConnectDB()
		if err != nil {
			log.Fatal(err)
		}

		if m.ChannelID == Server.SFWChannel { //SFW Check
			s.ChannelMessageSend(m.ChannelID, "This is a SFW channel!")
			return
		}

		var Tags = []Tag{}
		db.Find(&Tags)
		for i := 0; i < len(Tags); i++ {
			s.ChannelMessageSend(m.ChannelID, Tags[i].Name)
		}

		db.Close()
	}

	if strings.HasPrefix(m.Content, "!secret") {
		if set == 1 {
			s.ChannelMessageSend(m.ChannelID, "Secret already enabled")
			return
		}
		set = 1
		s.ChannelMessageSend(m.ChannelID, "Secret enabled!")

		Now := time.Now()
		NewYear, _ := time.Parse("2006-Jan-02", "2016-Jan-01")

		//NewYear := Now.Add(1 * time.Minute)
		log.Println(Now)
		log.Println(NewYear)
		log.Println(NewYear.Sub(Now))

		timer1 := time.NewTimer(NewYear.Sub(Now))
		go func() {
			<-timer1.C
			db, err := ConnectDB()
			if err != nil {
				log.Fatal(err)
			}

			HappyNewYear(db)
			s.ChannelMessageSend(m.ChannelID, "!picture HNY all")

			for i := 0; i < 10; i++ {
				s.ChannelMessageSend(m.ChannelID, "Happy New Year, groetjes Gerard!")
			}
			log.Println("Happy New Year, Docker!")
		}()

	}

}

//OnReady gets called when a OnReady event happens in the Session
//This event is neccessary to keep a websocket connection with the Discord API
func OnReady(s *discordgo.Session, st discordgo.Ready) {
	fmt.Println("* OnReady fired.")
	// start the Heartbeat
	s.Heartbeat(st.HeartbeatInterval)
}
