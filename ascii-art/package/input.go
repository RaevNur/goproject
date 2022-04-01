package asciiart

import (
	"errors"
	"fmt"
	"strings"
)

const ascii_offset = 32
const first_ascii = ' '
const last_ascii = '~'

type Symbol struct {
	Rune      rune
	BannerIdx int
	Color     string
}

func InitText(args *[]string, fig *figure) {
	if len(*args) == 0 {
		_CloseProgram(errors.New("Input func len(args) == 0"))
	}
	checkInput((*args)[0])
	input := getFormattedString((*args)[0])
	*args = (*args)[1:]
	for _, symbol := range input {
		if !isAvailableSymbol(symbol) {
			msg := fmt.Sprintf("Unknown symbol: '%v'!", string(symbol))
			_CloseProgram(errors.New(msg))
		}
		newSymbol := &Symbol{Rune: symbol, BannerIdx: int(symbol - ascii_offset), Color: "default"}
		fig.Text = append(fig.Text, newSymbol)
	}
}

func isAvailableSymbol(symbol rune) bool {
	if symbol != '\n' && symbol < first_ascii || last_ascii < symbol {
		return false
	}
	return true
}

func getFormattedString(input string) string {
	result := strings.ReplaceAll(input, "\\n", "\n")
	result = strings.ReplaceAll(result, "\\!", "!")
	return result
}

func checkInput(input string) {
	for _, symbol := range input {
		if symbol != '\n' && symbol < ' ' || '~' < symbol {
			err := fmt.Errorf("Unknown character '%q'", string(symbol))
			_CloseProgram(err)
		}
	}
}
