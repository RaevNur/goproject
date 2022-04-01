package asciiart

import (
	"fmt"
	"log"
)

const (
	ascii_offset = 32
	first_ascii  = ' '
	last_ascii   = '~'
)

type symbol struct {
	Rune      rune
	BannerIdx int
}

func checkText(text string) error {
	for _, symbol := range text {
		if !isAvailableSymbol(symbol) {
			msg := fmt.Sprintf(errUnknownCharacter, string(symbol))
			log.Printf("checkText: %v", msg)
			return fmt.Errorf(msg)
		}
	}
	return nil
}

func isAvailableSymbol(symbol rune) bool {
	if symbol != '\n' && symbol < first_ascii || last_ascii < symbol {
		return false
	}
	return true
}
