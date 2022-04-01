package asciiart

import (
	"fmt"
	"regexp"
	"strings"
)

type figure struct {
	Text       []*Symbol
	MODE       int
	IsColoured bool
	FileName   string
	Font       *font
	Align      *textAlign
}

const (
	MODE_DEFAULT = iota
	MODE_OUTPUT
	MODE_REVERSE
)

func (fig *figure) SetMode(mode int) {
	if fig.MODE == MODE_DEFAULT {
		fig.MODE = mode
	} else {
		err := fmt.Errorf("You cant combine --reverse & --output!")
		_CloseProgram(err)
	}
}

func (fig *figure) Constructor(args *[]string) {
	fig.InitReverse(args)
	if fig.MODE != MODE_REVERSE {
		fig.InitText(args)
	}
	fig.InitFont(args)
	fig.InitAlign(args)
	fig.InitOutput(args)
}

func (fig *figure) InitFont(args *[]string) {
	fig.Font = &font{}
	*args, fig.Font.Name = getFontNameByParam(*args)
	fig.Font.InitFont()
}

func (fig *figure) InitAlign(args *[]string) {
	_InitTerminal()
	cntArgs := len(*args)
	fig.Align = &textAlign{}
	pattern := regexp.MustCompile(`^--align=([A-Za-z]{1,})?$`)
	for i := 0; i < cntArgs; i++ {
		parametr := (*args)[i]

		isMatched := pattern.MatchString(parametr)
		if isMatched {
			align := pattern.FindStringSubmatch(parametr)[1]
			align = strings.ToLower(align)
			if isAlignAvailable(align) {
				fig.Align.Name = align
				fig.Align.Func = getAlignFormat(align)
			} else {
				err := fmt.Errorf("Incorrect '--align=' value: '%v'!", align)
				_CloseProgram(err)
			}
			tempArgs := append(make([]string, 0), (*args)[:i]...)
			*args = append(tempArgs, (*args)[i+1:]...)
			cntArgs, i = cntArgs-1, i-1
		}
	}
	if fig.Align.Name == "" {
		fig.Align.Name = _TextAlignLeft
		fig.Align.Func = getAlignLeft
	}
}

func (fig *figure) InitText(args *[]string) {
	InitText(args, fig)
	*args = fig.InitColors(*args)
}

func (fig *figure) Slice(withColour bool) (rows []string) {
	rows = make([]string, 0)
	line := []*Symbol{}
	for i := 0; i < len(fig.Text); i++ {
		char := fig.Text[i]
		if char.Rune == '\n' {
			rows = append(rows, getSliceRowsByLine(fig, line, withColour)...)
			line = []*Symbol{}
		} else {
			line = append(line, fig.Text[i])
		}
	}
	rows = append(rows, getSliceRowsByLine(fig, line, withColour)...)
	return rows
}

func getSliceRowsByLine(fig *figure, line []*Symbol, withColour bool) (rows []string) {
	if len(line) == 0 {
		return []string{""}
	}
	rows = make([]string, 0)
	for r := 0; r < fig.Font.Height; r++ {
		words := getTextForAlign(fig, line, r, withColour)
		if len(words) == 0 {
			return []string{""}
		}
		printRow := fig.Align.Func(words)
		rows = append(rows, printRow)
	}
	return rows
}

func getTextForAlign(fig *figure, line []*Symbol, row int, withColour bool) []string {
	if row < 0 || 7 < row {
		err := fmt.Errorf("getTextForAlign can take only (0 <= row <= 7) But we take: %v", row)
		_CloseProgram(err)
	}
	result := make([]string, 0)
	if fig.Align.Name == _TextAlignJustify { // Justify
		word := ""
		for _, char := range line {
			if char.Rune == ' ' && word != "" {
				result = append(result, word)
				word = ""
			} else if char.Rune != ' ' {
				word += fig.GetViewLine(char, row, withColour)
			}
		}
		if word != "" {
			result = append(result, word)
		}
	} else { // Center | Right | Left
		words := ""
		for _, char := range line {
			words += fig.GetViewLine(char, row, withColour)
		}
		result = append(result, words)
	}
	return result
}

func (fig *figure) GetViewLine(symbol *Symbol, row int, withColour bool) string {
	if row < 0 || 7 < row {
		err := fmt.Errorf("Index Out Of Range size = 8 but View[%v]", row)
		_CloseProgram(err)
	}
	if withColour {
		return getColorCode(symbol.Color) + fig.Font.Banners[symbol.BannerIdx].View[row] + colors["default"]
	} else {
		return fig.Font.Banners[symbol.BannerIdx].View[row]
	}
}

func (fig *figure) GetAllSubStrings(substr []rune) [][]*Symbol {
	textLen := len(fig.Text)
	substrLen := len(substr)
	result := make([][]*Symbol, 0)
	if substrLen > textLen || substrLen == 0 {
		return result
	}
	//Finding
	for i := 0; i <= textLen-substrLen; i++ {
		if fig.Text[i].Rune == substr[0] {
			appender := make([]*Symbol, substrLen)
			appender[0] = fig.Text[i]
			isEqual := true
			for j := 1; j < substrLen; j++ {
				if fig.Text[i+j].Rune == substr[j] {
					appender[j] = fig.Text[i+j]
				} else {
					isEqual = false
					break
				}
			}
			if isEqual {
				i = i + substrLen - 1
				result = append(result, appender)
			}
		}
	}
	return result
}
