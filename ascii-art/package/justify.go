package asciiart

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type textAlign struct {
	Name string
	Func func(str []string) string
}

const (
	_TextAlignLeft    = "left"
	_TextAlignRight   = "right"
	_TextAlignCenter  = "center"
	_TextAlignJustify = "justify"
)

// Returns Function Which gets Format
func getAlignFormat(align string) func(str []string) string {
	switch strings.ToLower(align) {
	case _TextAlignLeft:
		return getAlignLeft
	case _TextAlignRight:
		return getAlignRight
	case _TextAlignCenter:
		return getAlignCenter
	case _TextAlignJustify:
		return getAlignJustify
	default:
		return getAlignLeft
	}
}

func getAlignLeft(str []string) string {
	if len(str) == 0 {
		_CloseProgram(errors.New("ERROR: Aling func len(arg) == 0"))
	}
	return str[0]
}

func getAlignRight(str []string) string {
	if len(str) == 0 {
		_CloseProgram(errors.New("Aling func len(arg) == 0"))
	}
	tempTerminalWidth := _TerminalWidth + getColorsSizeInStr(str[0])
	format := fmt.Sprintf("%%%dv", tempTerminalWidth)
	return fmt.Sprintf(format, str[0])
}

func getAlignCenter(str []string) string {
	if len(str) == 0 {
		_CloseProgram(errors.New("Aling func len(arg) == 0"))
	}
	tempTerminalWidth := _TerminalWidth + getColorsSizeInStr(str[0])
	paddRight := tempTerminalWidth - (tempTerminalWidth-uint(len(str[0])))/2
	if paddRight > tempTerminalWidth {
		paddRight = 0
	}
	format := fmt.Sprintf("%%%dv", paddRight)
	return fmt.Sprintf(format, str[0])
}

func getAlignJustify(str []string) string {
	cntWords := len(str)
	if cntWords == 0 {
		_CloseProgram(errors.New("Aling func len(arg) == 0"))
	}
	sizes := make([]uint, cntWords)
	freePaddings := _TerminalWidth
	for i, word := range str {
		sizes[i] = uint(len(word)) - getColorsSizeInStr(word)
		freePaddings = freePaddings - sizes[i]
	}
	res := ""
	for _, word := range str {
		res += word
		if cntWords > 1 {
			cntWords--
			nextPadding := freePaddings / uint(cntWords)
			paddings := createPaddingWithLength(nextPadding)
			freePaddings -= nextPadding
			res += paddings
		}
	}
	return res
}

func createPaddingWithLength(size uint) string {
	paddings := ""
	for i := uint(0); i < size; i++ {
		paddings += " "
	}
	return paddings
}

func isAlignAvailable(align string) bool {
	switch strings.ToLower(align) {
	case _TextAlignLeft:
		return true
	case _TextAlignCenter:
		return true
	case _TextAlignRight:
		return true
	case _TextAlignJustify:
		return true
	default:
		return false
	}
}

func getColorsSizeInStr(str string) uint {
	pattern := regexp.MustCompile(_ColorsPattern)
	foundColors := pattern.FindAllString(str, -1)
	size := uint(0)
	for _, value := range foundColors {
		size += uint(len(value))
	}
	return size
}
