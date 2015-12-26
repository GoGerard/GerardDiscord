package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

//SendMeuk is a temp function that sends four messages to the ChannelID of the Message calling the function
func SendMeuk(s *discordgo.Session, m discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "Hallon, dit is GoGerard, Powered by Golang + AngularJS")
}

//ReturnRandomLine outputs a random line as a string from a provided file(string to location)
func ReturnRandomLine(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		err := "ERROR: file: '" + file + "' not found!"
		return err
	}
	lines := strings.Split(string(b), "\n")

	rand.Seed(time.Now().UnixNano())
	return lines[rand.Intn(len(lines))]
}
