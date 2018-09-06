package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func init() {
	newCommand("translate", helpTranslate(), commandTranslate)
}

func helpTranslate() string {
	return TranslateToSelstennes("== Aide de la commande `translate` ==\n" +
		"`translate` traduit une phrase dans la langue parfaite (jaune) Selstennes")
}

func commandTranslate(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, TranslateToSelstennes("Il faut me donner un message à traduire (y a que les 2A qui sont assez idiots pour ne pas y penser."))
		return nil
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)

	// The message to translate is splited in the args, so we should join all the parts together.
	message := strings.Join(args, " ")

	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+TranslateToSelstennes(message))

	return nil
}
