package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// command represents a command of the bot
type command struct {
	Name string
	Help string

	Exec func(*discordgo.Session, *discordgo.MessageCreate, []string) error
}

var commands = make(map[string]command)

// parseCommand read a message and call the appropriate command (if applicable)
func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, TranslateToSelstennes("Que dois-je faire ? Pour avoir la liste des commandes, utilisez la commande `help`."))
		return nil
	}

	// If the command exist, we exec it.
	if command, exist := commands[strings.ToLower(args[0])]; exist {
		command.Exec(s, m, args[1:])
	} else {
		s.ChannelMessageSend(m.ChannelID, TranslateToSelstennes(fmt.Sprintf("Je n’ai pas compris la commande `%s`, désolé. Pour avoir la liste des commandes, utilisez la commande `help`.", args[0])))
	}

	return nil
}

// newCommand register a command into the parser
func newCommand(name string, help string, exec func(*discordgo.Session, *discordgo.MessageCreate, []string) error) command {
	commands[name] = command{
		Name: name,
		Help: help,
		Exec: exec,
	}

	return commands[name]
}
