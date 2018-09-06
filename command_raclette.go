package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("raclette", helpRaclette(), commandRaclette)
}

func helpRaclette() string {
	return "== Aide de la commande `raclette` ==\n" +
		"`raclette` fait de la raclette"
}

func commandRaclette(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: "https://previews.123rf.com/images/margouillat/margouillat1611/margouillat161100158/65248462-fromage-%C3%A0-raclette-fondu.jpg",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

	return nil
}
