package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// connectDiscord open the connection with Discord by using the token stored in the configuration.
func connectDiscord() (*discordgo.Session, error) {
	// We create a new Discord session using the token (that we read from the configuration)
	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session")
	}

	// We add the main handler (to read messages).
	dg.AddHandler(handleMessageCreate)

	// Finally, we can connect to Discord.
	err = dg.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord connection")
	}

	return dg, nil
}
