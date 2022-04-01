package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

func (fig *figure) InitReverse(args *[]string) {
	cntArgs := len(*args)
	parametrFlag := "--reverse="
	isFileInited := false
	for i := 0; i < cntArgs; i++ {
		parametr := (*args)[i]
		if len(parametr) > 9 && parametr[:10] == parametrFlag {
			fileName := parametr[10:]
			if isFileInited {
				err := fmt.Errorf("You can have only 1 %v", parametrFlag)
				_CloseProgram(err)
			} else if fileName == "" {
				err := fmt.Errorf("Incorrect FileName for param %v, %q", parametrFlag, fileName)
				_CloseProgram(err)
			} else {
				fig.FileName = fileName
				fig.SetMode(MODE_REVERSE)
			}
			tempArgs := append(make([]string, 0), (*args)[:i]...)
			*args = append(tempArgs, (*args)[i+1:]...)
			cntArgs, i = cntArgs-1, i-1
			isFileInited = true
		}
	}
}

func (fig *figure) DoReverse() {
	checkParamsForReverse(fig)
	bannerLines := fig.getBannerLinesFromFile(fig.FileName)
	result := fig.convertBannerLinesToString(bannerLines)
	fmt.Printf("%v\n", result)
}

//Get String From BannersLine
func (fnt *font) getStrFromBannersLine(bannersLine []string) string {
	cntLines := len(bannersLine)
	if cntLines != fnt.Height {
		err := fmt.Errorf("Incorrect Banners Line")
		_CloseProgram(err)
	}

	result := ""
	starterIdx := 0
	lineSize := len(bannersLine[0])
	for starterIdx < lineSize {
		interval, isSymgolFind := fnt.MinWidth, false
		for ; interval <= fnt.MaxWidth; interval++ {
			isEqual, symbol := false, ' '
			endIdx := starterIdx + interval

			for bannerIdx, banner := range fnt.Banners {
				// Is Banner Symbol Equal
				if banner.Width == interval {
					isEqual = true
					for y := 0; y < fnt.Height; y++ {
						line := bannersLine[y][starterIdx:endIdx]
						if banner.View[y] != line {
							isEqual = false
							break
						}
					}
					if isEqual {
						symbol = rune(bannerIdx + first_ascii)
						break
					}
				} else {
					isEqual = false
				}
			}
			if isEqual {
				isSymgolFind = true
				result += string(symbol)
				break
			}
		}
		if !isSymgolFind {
			err := fmt.Errorf("Incorrect Banners in Line")
			_CloseProgram(err)
		}
		starterIdx += interval
	}
	return result
}

//Get Banners Symbols in String
func (fig *figure) convertBannerLinesToString(lines [][]string) string {
	result := ""
	cntBannersLine := len(lines)

	for i, bannerLine := range lines {
		linesCnt := len(bannerLine)
		if linesCnt == fig.Font.Height {
			result += fig.Font.getStrFromBannersLine(bannerLine)
		} else if linesCnt == 1 {
		} else {
			err := fmt.Errorf("Incorrect Banners in file")
			_CloseProgram(err)
		}
		if i+1 != cntBannersLine {
			result += "\n"
		}
	}

	return result
}

func (fig *figure) getBannerLinesFromFile(fileName string) [][]string {
	if _, err := os.Stat(fileName); err != nil {
		_CloseProgram(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		_CloseProgram(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bannerLines := [][]string{{}}
	oldLen, lineIdx := 0, 0
	maxHeight, curHeight := fig.Font.Height, 0
	if scanner.Scan() {
		curHeight++
		part := scanner.Text()
		bannerLines[lineIdx] = append(bannerLines[lineIdx], part)
		oldLen = len(part)
	}
	for scanner.Scan() {

		part := scanner.Text()
		curLen := len(part)
		if curLen == 0 || curLen != oldLen || curHeight >= maxHeight {
			bannerLines = append(bannerLines, []string{})
			lineIdx++
			curHeight = 0
		}
		bannerLines[lineIdx] = append(bannerLines[lineIdx], part)
		oldLen = curLen
		curHeight++
	}
	return bannerLines
}

func checkParamsForReverse(fig *figure) {
	if fig.Align.Name != _TextAlignLeft {
		fig.Align.Name = _TextAlignLeft
		fig.Align.Func = getAlignLeft

		msg := fmt.Sprintf("Align params is not allowed!", _TextAlignLeft)
		_ShowWarningMessage(msg)
	}
	if fig.IsColoured {
		msg := fmt.Sprintf("The colors for the text were not applied", _TextAlignLeft, _TextAlignLeft)
		_ShowWarningMessage(msg)
	}
	if fig.FileName == "" {
		_CloseProgram(fmt.Errorf("You cant read file with name: ''!"))
	}
}
