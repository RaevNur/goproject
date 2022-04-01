package asciiart

import (
	"fmt"
	"log"
)

type banner struct {
	Width        int
	CountSymbols int
	View         []string
}

func (bnr *banner) SetBannerLine(lineIdx int, line string) error {
	bnr.View[lineIdx] = line
	width := len(line)
	cntSymbols := len([]rune(line))
	if bnr.CountSymbols != 0 && bnr.CountSymbols != cntSymbols {
		msg := errIncorrectBannersInFile
		log.Printf("SetBannerLine: %v", msg)
		return fmt.Errorf(msg)
	}
	bnr.Width = width
	bnr.CountSymbols = cntSymbols
	return nil
}
