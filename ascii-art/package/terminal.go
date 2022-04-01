package asciiart

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var _TerminalHeight uint
var _TerminalWidth uint

// Can Panic on Debug
func _InitTerminal() {
	err := error(nil)
	_TerminalHeight, _TerminalWidth, err = getTerminalSize()

	if err != nil {
		_CloseProgram(err)
	}
}

// Gets Command result on System Terminal "stty size"
// Returns: "height width\n", error
func stty_size() (string, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	return string(output), err
}

// Returns: "height", "width", error
func getTerminalSizeString() (string, string, error) {
	sizeStr, err := stty_size()
	if err != nil {
		return "", "", err
	}
	size := strings.Split(strings.Replace(sizeStr, "\n", "", 1), " ")
	return size[0], size[1], nil
}

// Returns: height, width, error
func getTerminalSize() (uint, uint, error) {
	sizeStr, err := stty_size()
	if err != nil {
		return 0, 0, err
	}
	size := strings.Split(strings.Replace(sizeStr, "\n", "", 1), " ")
	height, err := strconv.Atoi(size[0])
	if err != nil {
		return 0, 0, err
	}
	width, err := strconv.Atoi(size[1])
	if err != nil {
		return 0, 0, err
	}
	return uint(height), uint(width), nil
}
