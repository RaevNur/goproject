package asciiart

import (
	"errors"
	"fmt"
	"os"
)

var _Figure figure

func Program(args []string) {
	_Constructor(&args)
	_RunProgram(&_Figure)
}

func _Constructor(args *[]string) {
	_CheckForCntArgs(*args)
	_InitColors()
	_Figure.Constructor(args)
}

//Печатает если надо печатать
//Читает с файла если надо читать
//Сохраняет в файл если надо
func _RunProgram(fig *figure) {
	if fig.MODE == MODE_OUTPUT {
		fig.DoOutput()
	} else if fig.MODE == MODE_REVERSE {
		fig.DoReverse()
	} else {
		fig.DoDefault()
	}
}

func (fig *figure) DoDefault() {
	for _, row := range fig.Slice(fig.IsColoured) {
		fmt.Println(row)
	}
}

func _CheckForCntArgs(args []string) {
	if len(args) == 0 {
		msg := "You didn't enter any arguments!"
		_CloseProgram(errors.New(msg))
	}
}

// Closing Program With Exit Status if Have Error
func _CloseProgram(err error) {
	if err != nil {
		fmt.Printf("%sERROR: %s%s\n", "\u001b[31;1m", err.Error(), "\033[0m")
		os.Exit(1)
	}
	os.Exit(0)
}

func _ShowWarningMessage(msg string) {
	fmt.Printf("%sWarning: %s%s\n", "\u001b[33;1m", msg, "\033[0m")
}

func _ShowSuccessfulMessage(msg string) {
	fmt.Printf("%sSuccess: %s%s\n", "\u001b[32;1m", msg, "\033[0m")
}
