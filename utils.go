package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type TranslateOperation struct {
	From        int
	To          int
	Translation string
}

// TranslateToSelstennes translate a text (message) to the Selstennes (perfect (yellow)) language.
func TranslateToSelstennes(message string) string {
	msg := strings.Split(message, "\"")

	var result []string

	for index, section := range msg {
		if index%2 == 0 {
			result = append(result, TranslateSectionToSelstennes(section))
		} else {
			result = append(result, section)
		}
	}

	return strings.Join(result, "\"")
}

func TranslateSectionToSelstennes(message string) string {
	config, err := getConfig()
	if err != nil {
		log.WithField("error", err).Fatal("error getting the config")
	}

	var operations []TranslateOperation

	for _, rule := range config.Rules {
		re := regexp.MustCompile(rule.Regex)
		submatches := re.FindAllStringSubmatchIndex(message, -1)
		if submatches != nil {
			for _, submatch := range submatches {
				operations = append(operations, TranslateOperation{
					From:        submatch[0],
					To:          submatch[1],
					Translation: rule.Translations[rand.Intn(len(rule.Translations))],
				})
			}
		}
	}

	m := strings.Split(message, "")

	for i, j, w := 0, 0, 0; i < len(message); i, j = i+w, j+1 {
		_, w = utf8.DecodeRuneInString(message[i:])

		for _, op := range operations {
			if i == op.From {
				m[j] = op.Translation
			} else if i > op.From && i < op.To {
				m[j] = ""
			}
		}
	}

	return strings.Join(m, "")
}
