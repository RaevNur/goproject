package asciiart

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	CSI                 = "\033["
	FgRGB256            = "38"
	BgRGB256            = "48"
	FgDefault           = "39"
	BgDefault           = "49"
	_ColorsPattern      = `\[([\d]{1,2};)?([\d]{1,2};)?([\d]{1,3};)?([\d]{1,3};)?([\d]{1,3})?m`
	_ColorsParamPattern = `^--color=(RGB\([\d]{1,3},[ ]?[\d]{1,3},[ ]?[\d]{1,3}\)|\w{1,})?( symbols=\{[!-\~]+\}| substr=\{[!-\~]+\})?( in=\{[\d]+\-?[\d]{0,}\})?$`
)

//
var colors map[string]string

func _InitColors() {
	colors = make(map[string]string, 12)
	colors["default"] = "\033[0m"
	colors["red"] = "\033[31m"
	colors["orange"] = "\033[38;2;255;165;0m"
	colors["green"] = "\033[32m"
	colors["yellow"] = "\033[33m"
	colors["blue"] = "\033[34m"
	colors["purple"] = "\033[35m"
	colors["cyan"] = "\033[36m"
	colors["gray"] = "\033[37m"
	colors["white"] = "\033[97m"
}

func (fig *figure) InitColors(args []string) []string {
	cntArgs := len(args)
	for i := 0; i < cntArgs; i++ {
		parametr := args[i]
		if len(parametr) > 7 && parametr[:8] == "--color=" {
			setColorForText(parametr, fig)
			fig.IsColoured = true

			tempArgs := append(make([]string, 0), args[:i]...)
			args = append(tempArgs, args[i+1:]...)
			cntArgs, i = cntArgs-1, i-1
		}
	}
	return args
}

// Кеширует цвет Если это новый цвет. И возвращяет код цвета
func getColorCode(colour string) string {
	pattern := regexp.MustCompile(`^RGB\(([\d]{1,3}),[ ]?([\d]{1,3}),[ ]?([\d]{1,3})\)$`)

	if code, isIn := colors[colour]; isIn {
		return code
	} else if isMatched := pattern.MatchString(colour); !isMatched {
		err := fmt.Errorf("Incorrect color value: %q", colour)
		_CloseProgram(err)
	}
	subMatchs := pattern.FindStringSubmatch(colour)
	red, green, blue := subMatchs[1], subMatchs[2], subMatchs[3]
	checkRGBValues(red, green, blue)

	code := fmt.Sprintf("%v%v;2;%s;%s;%sm", CSI, FgRGB256, red, green, blue)
	colors[colour] = code
	return code
}

func checkRGBValues(red, green, blue string) {
	isWrongRed := len(red) == 3 && red > "255"
	isWronBlue := len(blue) == 3 && blue > "255"
	isWronGreen := len(green) == 3 && green > "255"

	if isWrongRed || isWronGreen || isWronBlue {
		msg := "Incorrect value of RGB(red, green, blue)"
		if isWrongRed {
			msg += fmt.Sprintf(" red = %v!", red)
		}
		if isWronGreen {
			msg += fmt.Sprintf(" green = %v!", green)
		}
		if isWronBlue {
			msg += fmt.Sprintf(" blue = %v!", blue)
		}
		err := errors.New(msg)
		_CloseProgram(err)
	}
}

func getFromToInColorParam(parametr string) (int, int) {
	patternFromTo := regexp.MustCompile(`in=\{([\d]+)-?([\d]+)?\}`)
	if isIn := patternFromTo.MatchString(parametr); !isIn {
		return 0, 0
	}

	err, from, to := error(nil), 0, 0
	submatchGroup := patternFromTo.FindStringSubmatch(parametr)

	from, err = strconv.Atoi(submatchGroup[1])
	if err != nil {
		_CloseProgram(err)
	}

	if submatchGroup[2] == "" {
		return from, from
	} else {
		to, err = strconv.Atoi(submatchGroup[2])
		if err != nil {
			_CloseProgram(err)
		}
		if from > to {
			return to, from
		}
	}
	return from, to
}

func setColorForText(parametr string, fig *figure) {
	//--color=RGB(1,1,1) symbols={!-~} in={0-10}
	patternParams := regexp.MustCompile(_ColorsParamPattern)
	isMatched := patternParams.MatchString(parametr)
	if !isMatched {
		tmp := "--color=RGB(255,255,255) [symbols={Characters}/substr={SubStrings}] [in={from-to/idx of element}]"
		err := fmt.Errorf("Incorrect value on: '%v'!\nTemplate: %v", parametr, tmp)
		_CloseProgram(err)
	}
	//Разделяем на под параметры
	paramsGroup := patternParams.FindStringSubmatch(parametr)

	color := paramsGroup[1]
	getColorCode(color)

	from, to := getFromToInColorParam(paramsGroup[3])
	patternSymbols := regexp.MustCompile(`symbols=\{([!-\~]+)\}`)
	patternSubstr := regexp.MustCompile(`substr=\{([!-\~]+)\}`)

	if patternSymbols.MatchString(paramsGroup[2]) {
		symbols := patternSymbols.FindStringSubmatch(paramsGroup[2])[1]
		symbols = getFormattedString(symbols)
		setColorForInputSymbols(fig, []rune(symbols), color, from, to)
	} else if patternSubstr.MatchString(paramsGroup[2]) {
		substr := patternSubstr.FindStringSubmatch(paramsGroup[2])[1]
		substr = getFormattedString(substr)
		setColorForInputSubstrings(fig, []rune(substr), color, from, to)
	} else {
		setColorForInputAllSymbols(fig, color, from, to)
	}
}

func setColorForInputSymbol(fig *figure, symbol rune, color string, from, to int) {
	inputSize := len(fig.Text)
	if inputSize <= to {
		to = inputSize
	}
	if from == 0 && to == 0 {
		for i := 0; i < inputSize; i++ {
			if fig.Text[i].Rune == symbol {
				fig.Text[i].Color = color
			}
		}
	} else {
		cnt := 1
		for i := 0; i < inputSize; i++ {
			if fig.Text[i].Rune == symbol {
				if from <= cnt && cnt <= to {
					fig.Text[i].Color = color
				}
				cnt++
			} else if to < cnt {
				break
			}
		}
	}
}

func setColorForInputSymbols(fig *figure, symbols []rune, color string, from, to int) {
	for _, symbol := range symbols {
		setColorForInputSymbol(fig, symbol, color, from, to)
	}
}

func setColorForInputAllSymbols(fig *figure, color string, from, to int) {
	inputSize := len(fig.Text)
	if from == 0 {
		from++
	}
	if inputSize <= to || to == 0 {
		to = inputSize
	}
	for i := from - 1; i < to; i++ {
		fig.Text[i].Color = color
	}
}

func setColorForInputSubstrings(fig *figure, substr []rune, color string, from, to int) {
	subStrs := fig.GetAllSubStrings(substr)

	if from == 0 && to == 0 {
		for _, symbols := range subStrs {
			for _, symbol := range symbols {
				symbol.Color = color
			}
		}
	} else {
		if from != 0 {
			from--
		}
		for i, symbols := range subStrs {
			if from <= i && i < to {
				for _, symbol := range symbols {
					symbol.Color = color
				}
			}
		}
	}
}
