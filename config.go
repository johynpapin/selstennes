package main

import (
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Rule struct {
	Regex        string
	Translations []string
}

type Config struct {
	Shortname string
	Name      string
	Rules     []Rule
}

func defineFlags() {
	flag.CommandLine.String("token", "", "Bot Token")
	flag.CommandLine.String("shortname", "s", "Short name of the BOT")
	flag.CommandLine.String("name", "selstennes", "Name of the BOT")

	flag.Parse()
}

func loadConfig() error {
	defineFlags()

	viper.BindPFlags(flag.CommandLine)

	viper.SetEnvPrefix("selstennes")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "error reading config")
	}

	return nil
}

func getConfig() (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config")
	}

	return &config, nil
}
