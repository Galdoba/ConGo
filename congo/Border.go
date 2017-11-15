package congo

import "strings"

//BorderMap -
var BorderMap TBorderMap

//DefaultBorder -
var DefaultBorder TBorder

//TBorderMap -
type TBorderMap struct {
	Border map[string]TBorder
}

//TBorder -
type TBorder struct {
	borderName  string
	borderCells []string
	isCreated   bool
}

//GetTBorderName -
func GetTBorderName() string {
	return DefaultBorder.borderName
}

//SetTBorder -
func SetTBorder(borderName string) {

	if bord, ok := BorderMap.Border[borderName]; ok {
		DefaultBorder = bord
	} else {
		DefaultBorder = BorderMap.Border["Default"]
	}
}

//AddBorder -
func AddBorder(name, borderRunes string) *TBorderMap {
	var style TBorder
	var borderCells []string
	borderCells = make([]string, 6)
	style.isCreated = true
	style.borderName = name
	style.borderCells = strings.Split(borderRunes, "")
	for i := range borderCells {
		borderCells[i] = string(borderRunes[i])
	}
	BorderMap.Border[name] = style

	return &BorderMap
}

//InitBorders -
func InitBorders() {
	BorderMap = TBorderMap{}
	BorderMap.Border = map[string]TBorder{} //Инициализируем МАР внутри типа

	AddBorder("Double Line", "═╗║╝╚╔")
	AddBorder("Single Line", "─┐│┘└┌")
	AddBorder("Block", "██████")
	AddBorder("EX", "XXXXXX")
	AddBorder("Default", "-+|+++")
	AddBorder("Test", "123456")
	AddBorder("None", "      ")
	DefaultBorder = BorderMap.Border["Default"]
}

func pickStyle(name string) []string {
	var borderRune []string
	styleFound := false
	if border, ok := BorderMap.Border[name]; ok {
		styleFound = true
		borderRune = border.borderCells
	}

	if styleFound == false {
		panic("Style <" + name + "> not found in  'pickStyle()' ")
	}
	return borderRune
}

//DrawBorder -
func DrawBorder(x, y, w, h int, name string) { //прийти к (Координаты,размер,тип рамки, цвета)
	borderRune := pickStyle(name)
	for i := 1; i < w-1; i++ {
		PrintText(x+i, y, borderRune[0])
		PrintText(x+i, y+h-1, borderRune[0])
	}
	for i := 1; i < h-1; i++ {
		PrintText(x, y+i, borderRune[2])
		PrintText(x+w-1, y+i, borderRune[2])
	}
	PrintText(x, y, borderRune[5])
	PrintText(x+w-1, y, borderRune[1])
	PrintText(x+w-1, y+h-1, borderRune[3])
	PrintText(x, y+h-1, borderRune[4])
}

//SetBorderCells -
func (style *TBorder) SetBorderCells(input string) {
	borderRunes := strings.Split(input, "")
	for i := 0; i < len(borderRunes); i++ {
		style.borderCells[i] = borderRunes[i]
	}
}

//GetNames -
func (style *TBorderMap) GetNames() []string {
	names := make([]string, 0, len(style.Border))
	for n := range style.Border {
		names = append(names, n)
	}
	return names
}
