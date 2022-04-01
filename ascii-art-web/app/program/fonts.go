package asciiartweb

import (
	"asciiartweb/app/asciiart"
	"log"
	"os"
)

var Fonts map[string]bool

func InitFonts() {
	fonts, err := asciiart.GetAllFonts()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	Fonts = fonts
}
