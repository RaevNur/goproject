package asciiart

import (
	"fmt"
)

type figure struct {
	Text []*symbol
	Font *font
}

func (fig *figure) Constructor(configs *ASCIIConfigs) error {
	if err := fig.InitText(configs.Text); err != nil {
		return err
	}
	if err := fig.InitFont(configs.FontName); err != nil {
		return err
	}
	return nil
}

func (fig *figure) InitText(text string) error {
	if err := checkText(text); err != nil {
		return err
	}
	for _, sym := range text {
		newSymbol := &symbol{Rune: sym, BannerIdx: int(sym - ascii_offset)}
		fig.Text = append(fig.Text, newSymbol)
	}
	return nil
}

func (fig *figure) InitFont(fontName string) error {
	fig.Font = &font{Name: fontName}
	return fig.Font.InitFont()
}

func (fig *figure) Slice() (rows []string) {
	rows = make([]string, 0)
	line := []*symbol{}
	for i := 0; i < len(fig.Text); i++ {
		char := fig.Text[i]
		if char.Rune == '\n' {
			rows = append(rows, getSliceRowsByLine(fig.Font, line)...)
			line = []*symbol{}
		} else {
			line = append(line, fig.Text[i])
		}
	}
	rows = append(rows, getSliceRowsByLine(fig.Font, line)...)
	return rows
}

func (fig *figure) GetArt() *string {
	result := ""
	for _, row := range fig.Slice() {
		result += fmt.Sprintf("%v\n", row)
	}
	return &result
}
