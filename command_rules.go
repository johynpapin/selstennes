package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func init() {
	newCommand("rules", helpRules(), commandRules)
}

func helpRules() string {
	return "== Aide de la commande `rules` ==\n" +
		"`rulese` affiche les règles de traduction de la commande `translate`"
}

func commandRules(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	s.ChannelMessageSend(m.ChannelID, "- Les sections entre \" sont ignorées.")

	config, _ := getConfig()

	var rules []string

	for _, rule := range config.Rules {
		rules = append(rules, "`"+rule.Regex+" -> "+strings.Join(rule.Translations, " ou ")+"`")
	}

	s.ChannelMessageSend(m.ChannelID, "- "+strings.Join(rules, ", "))

	return nil
}
