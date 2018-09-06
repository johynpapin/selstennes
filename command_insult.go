package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("insult", helpInsult(), commandInsult)
}

func helpInsult() string {
	return "== Aide de la commande `insult` ==\n" +
		"`insult` affiche une insulte"
}

type insult struct {
	Insult struct {
		ID            int         `json:"id"`
		Value         string      `json:"value"`
		TotalVoteUp   int         `json:"total_vote_up"`
		TotalVoteDown int         `json:"total_vote_down"`
		TotalVote     int         `json:"total_vote"`
		CurrentVote   interface{} `json:"current_vote"`
	} `json:"insult"`
}

func commandInsult(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	insult := new(insult)

	err := getJson("https://www.insult.es/api/random", insult)
	if err != nil {
		return err
	}

	s.ChannelMessageSend(m.ChannelID, insult.Insult.Value)

	return nil
}
