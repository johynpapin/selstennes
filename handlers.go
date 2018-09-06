package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"strings"
)

// handleMessageCreate check if a message is destinated to the bot. If it is the case, it will parse the message to execute the requested action.
func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.Split(m.Content, " ")

	if msg[0] == "!"+viper.GetString("shortname") || msg[0] == "!"+viper.GetString("name") {
		err := parseCommand(s, m, msg[1:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, TranslateToSelstennes("Désolé, j’ai rencontré un problème interne."))
		}
	} else if c, _ := s.Channel(m.ChannelID); c.Name == "selstennes" {
		err := commands["translate"].Exec(s, m, msg)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, TranslateToSelstennes("Désolé, j’ai rencontré un problème interne."))
		}
	}
}
