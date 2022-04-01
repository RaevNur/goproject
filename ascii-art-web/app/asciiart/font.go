package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Font Path Where takes Every Font
const fontsPath = "app/asciiart/fonts/"

type font struct {
	Name    string
	Banners [95]*banner
	Height  int
}

// InitFont - Read and Set Font From File
func (fnt *font) InitFont() error {
	filename := fmt.Sprintf("%v%v.txt", fontsPath, fnt.Name)
	file, err := os.Open(filename)
	if err != nil {
		msg := fmt.Sprintf(errUnknownFont, fnt.Name)
		log.Printf("fnt.InitFont > os.Open: %v", msg)
		return fmt.Errorf(msg)
	}
	defer file.Close()

	fnt.Height = 8
	charIdx, lineInd := 0, -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && charIdx < len(fnt.Banners) {
		if lineInd == -1 {
			// Initialize Banners
			fnt.Banners[charIdx] = &banner{View: make([]string, fnt.Height)}
		} else if lineInd > -1 {
			// Set Banner Line
			line := scanner.Text()
			if err := fnt.Banners[charIdx].SetBannerLine(lineInd, line); err != nil {
				return err
			}
		}
		lineInd++
		if lineInd == fnt.Height {
			// Reset Banner
			lineInd = -1
			charIdx++
		}
	}
	// Check For Banner Counts
	cntBanners := int(last_ascii-first_ascii) + 1
	if charIdx != cntBanners {
		msg := errIncorrectBannersInFile
		log.Printf("fnt.InitFont: %v", msg)
		return fmt.Errorf(msg)
	}
	return nil
}

func getSliceRowsByLine(fnt *font, line []*symbol) []string {
	if len(line) == 0 {
		return []string{""}
	}
	rows := make([]string, 0)
	for r := 0; r < fnt.Height; r++ {
		printRow := getTextRow(fnt, line, r)
		rows = append(rows, printRow)
	}
	return rows
}

func getTextRow(fnt *font, line []*symbol, row int) string {
	words := ""
	for _, char := range line {
		words += fnt.Banners[char.BannerIdx].View[row]
	}
	return words
}
