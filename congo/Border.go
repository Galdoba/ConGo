package congo

import "strings"

var StylePool []BorderStyle

type BorderStyle struct {
	name        string
	borderCells []string
	isCreated   bool
}

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

func DrawBorder(x, y, w, h int, name string, fgCol, bgCol int) { //прийти к (Координаты,размер,тип рамки, цвета)
	borderRune := pickStyle(name)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if j == 0 || j == h-1 {
				PrintText(x+i, y+j, borderRune[0], fgCol, bgCol)
			}
			if i == 0 || i == w-1 {
				PrintText(x+i, y+j, borderRune[2], fgCol, bgCol)
			}
			if i == 0 && j == 0 {
				PrintText(x+i, y+j, borderRune[5], fgCol, bgCol)
			}
			if i == w-1 && j == 0 {
				PrintText(x+i, y+j, borderRune[1], fgCol, bgCol)
			}
			if i == w-1 && j == h-1 {
				PrintText(x+i, y+j, borderRune[3], fgCol, bgCol)
			}

			if i == 0 && j == h-1 {
				PrintText(x+i, y+j, borderRune[4], fgCol, bgCol)
				//termbox.SetCell(x+i, y+j, borderRune[4], termbox.Attribute(fgCol), termbox.Attribute(bgCol))
			}
		}
	}
}

func (style *BorderStyle) SetBorderCells(input string) {

	//borderRunes := []rune(input)
	borderRunes := strings.Split(input, "")
	//fmt.Println(borderRunes)
	for i := 0; i < len(borderRunes); i++ {
		style.borderCells[i] = borderRunes[i]

		//	fmt.Println(style.borderCells[i])

	}
}

func (style *BorderStyle) SetStyleName(input string) {
	style.name = input
}

func (style *BorderStyle) GetStyleName() string {
	return style.name
}

func (style *BorderStyle) GetBorderCells() []string {
	return style.borderCells
}
