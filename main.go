package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language] string {
	"el": "ηελλο ςορλδ",
	"en": "Hello World",
	"fr": "Bonjour le monde",
	"yr": "Eku Ojumo",
}

func greet(lang language) string {

	greeting, ok := phrasebook[lang]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", lang)
	}

	return greeting


/*	switch lang {
	case "en":
		return "Hello World!"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
*/
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}
