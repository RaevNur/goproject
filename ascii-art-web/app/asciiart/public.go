package asciiart

import (
	"errors"
	"os"
	"regexp"
)

// ASCIIConfigs - Using for text params
type ASCIIConfigs struct {
	Text     string
	FontName string
}

// GetArtByConfigs - Gets arts by AsciiConfigs
func GetArtByConfigs(configs *ASCIIConfigs) (string, error) {
	if configs == nil {
		return "", errors.New(errNilPointerException)
	}
	myFigure := &figure{}
	if err := myFigure.Constructor(configs); err != nil {
		return "", err
	}
	return *myFigure.GetArt(), nil
}

// GetAllFonts - Returns All Fonts in Folder
func GetAllFonts() (map[string]bool, error) {
	files, err := os.ReadDir(fontsPath)
	if err != nil {
		return nil, err
	}
	// Getting All Files with extention
	pattern := regexp.MustCompile(`(.{1,}).txt$`)
	result := make(map[string]bool, len(files))
	for _, file := range files {
		fName := file.Name()
		isMatched := pattern.MatchString(fName)
		if isMatched {
			fName = pattern.FindStringSubmatch(fName)[1]
			result[fName] = true
		}
	}
	return result, nil
}

// IsValidInputText - Check For Valid Text input
func IsValidInputText(text string) bool {
	if checkText(text) != nil {
		return false
	}
	return true
}
