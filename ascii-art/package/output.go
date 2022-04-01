package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

func (fig *figure) InitOutput(args *[]string) {
	cntArgs := len(*args)
	isFileInited := false
	for i := 0; i < cntArgs; i++ {
		parametr := (*args)[i]
		if len(parametr) > 8 && parametr[:9] == "--output=" {
			fileName := parametr[9:]
			if isFileInited {
				err := fmt.Errorf("You can have only 1 --output")
				_CloseProgram(err)
			} else if fileName == "" {
				err := fmt.Errorf("Incorrect FileName for param --output")
				_CloseProgram(err)
			} else {
				fig.FileName = fileName
				fig.SetMode(MODE_OUTPUT)
			}
			tempArgs := append(make([]string, 0), (*args)[:i]...)
			*args = append(tempArgs, (*args)[i+1:]...)
			cntArgs, i = cntArgs-1, i-1
			isFileInited = true
		}
	}
}

func (fig *figure) DoOutput() {
	checkParamsForOutput(fig)
	if _, err := os.Stat(fig.FileName); err == nil {
		err := fmt.Errorf("File %q is Exists!", fig.FileName)
		_CloseProgram(err)
	}
	file, err := os.Create(fig.FileName)
	if err != nil {
		_CloseProgram(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range fig.Slice(false) {
		fmt.Fprintln(writer, line)
	}
	if err := writer.Flush(); err != nil {
		_CloseProgram(err)
	}

	msg := fmt.Sprintf("You have successfully written to the file: %q", fig.FileName)
	_ShowSuccessfulMessage(msg)
}

func checkParamsForOutput(fig *figure) {
	if fig.Align.Name != _TextAlignLeft {
		fig.Align.Name = _TextAlignLeft
		fig.Align.Func = getAlignLeft

		msg := fmt.Sprintf("The aligns for the text were not applied\nAlign parametr choosed to '%v'", _TextAlignLeft)
		_ShowWarningMessage(msg)
	}
	if fig.IsColoured {
		msg := fmt.Sprintf("The colors for the text were not applied")
		_ShowWarningMessage(msg)
	}
	if fig.FileName == "" {
		_CloseProgram(fmt.Errorf("You cant create file with name ''!"))
	}
}
