package main

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// We give a seed to the pseudo random number generator.
	rand.Seed(time.Now().Unix())

	// Then, we load the configuration.
	err := loadConfig()
	if err != nil {
		log.WithField("error", err).Fatal("error loading config")
	}

	// Now that we have the configuration, we can connect to Discord (the token is in the configuration).
	dg, err := connectDiscord()
	if err != nil {
		log.WithField("error", err).Fatal("error connecting to discord")
	}

	log.Info("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	log.Info("Good bye!")
}
