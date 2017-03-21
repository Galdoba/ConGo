package congo

import "strings"

//StylePool -
var StylePool []BorderStyle

//BorderStyle -
type BorderStyle struct {
	name        string
	borderCells []string
	isCreated   bool
}

func InitBorders() {
	AddBorderStyle("Double Line", "═╗║╝╚╔")
	AddBorderStyle("Single Line", "─┐│┘└┌")
	AddBorderStyle("Block", "█████")
	AddBorderStyle("EX", "XXXXXX")
	AddBorderStyle("Default", "-+|+++")
}

//AddBorderStyle -
func AddBorderStyle(name, borderRunes string) *[]BorderStyle {
	var style BorderStyle
	style.name = name
	style.borderCells = make([]string, 6)
	style.SetBorderCells(borderRunes)
	style.isCreated = true
	StylePool = append(StylePool, style)
	return &StylePool
}

func pickStyle(name string) []string {
	var borderRune []string
	styleFound := false
	for i := range StylePool {
		if name == StylePool[i].name {
			borderRune = StylePool[i].GetBorderCells()
			styleFound = true
		}
	}
	if styleFound == false {
		panic("Style <" + name + "> not found in  'pickStyle()' ")
	}
	return borderRune
}

//DrawBorder -
func DrawBorder(x, y, w, h int, name string, fgCol, bgCol int) { //прийти к (Координаты,размер,тип рамки, цвета)
	borderRune := pickStyle(name)
	for i := 1; i < w-1; i++ {
		PrintText(x+i, y, borderRune[0], fgCol, bgCol)
		PrintText(x+i, y+h-1, borderRune[0], fgCol, bgCol)
	}
	for i := 1; i < h-1; i++ {
		PrintText(x, y+i, borderRune[2], fgCol, bgCol)
		PrintText(x+w-1, y+i, borderRune[2], fgCol, bgCol)
	}
	PrintText(x, y, borderRune[5], fgCol, bgCol)
	PrintText(x+w-1, y, borderRune[1], fgCol, bgCol)
	PrintText(x+w-1, y+h-1, borderRune[3], fgCol, bgCol)
	PrintText(x, y+h-1, borderRune[4], fgCol, bgCol)
}

//SetBorderCells -
func (style *BorderStyle) SetBorderCells(input string) {
	borderRunes := strings.Split(input, "")
	for i := 0; i < len(borderRunes); i++ {
		style.borderCells[i] = borderRunes[i]
	}
}

//SetStyleName -
func (style *BorderStyle) SetStyleName(input string) {
	style.name = input
}

//GetStyleName -
func (style *BorderStyle) GetStyleName() string {
	return style.name
}

//GetBorderCells -
func (style *BorderStyle) GetBorderCells() []string {
	return style.borderCells
}
