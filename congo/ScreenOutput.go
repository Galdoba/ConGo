package congo

import "github.com/nsf/termbox-go"

import "bytes"

//"time"

var cursorX int
var cursorY int
var fgCol TColor
var bgCol TColor
var boundX int
var boundY int

//InitOutput - 
func InitOutput() {
	SetFgColor(ColorDefault)
	SetBgColor(ColorDefault)
	//SetBounds(termbox.Size())
	ResetBounds()
}

//GetSize -
func GetSize() (int, int) {
	return termbox.Size()
}

//FillRect -
func FillRect(x, y, w, h int, bckground rune, fgCol, bgCol TColor) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			termbox.SetCell(x+i, y+j, bckground, termbox.Attribute(fgCol.color), termbox.Attribute(bgCol.color))
		}
	}

}

//PrintText -
func PrintText(x, y int, input string) {
	rInput := []rune(input)
	var curentRune rune
	maxX, _ := GetBounds()
	//cursorY = y
	/*var tbFgCol int32
	var tbBgCol int32
	tbFgCol = int32(GetFgColor().color)
	tbBgCol = int32(GetFgColor().color)*/
	for i := range rInput {
		if x < maxX {
		curentRune = rInput[i]
		termbox.SetCell(x, y, curentRune, termbox.Attribute(fgCol.color), termbox.Attribute(bgCol.color))
		x++
		}
		//ResetBounds()
	}
}

//ResetBounds - 
func ResetBounds() {
	boundX, boundY = termbox.Size()
}

//GetBounds - 
func GetBounds() (int, int) {
	maxX, maxY := boundX, boundY
	return maxX, maxY
}

//SetBounds - 
func SetBounds(maxX, maxY int) {
	boundX, boundY = maxX, maxY
}

//SetFgColor - 
func SetFgColor(col TColor) {
	fgCol = col
}

//SetBgColor - 
func SetBgColor(col TColor) {
	bgCol = col
}

//GetFgColor - 
func GetFgColor() TColor {
	return fgCol
}

//GetBgColor - 
func GetBgColor() TColor {
	return bgCol
}

func trim(sl []rune, w int, align string, offSet int) ([]rune, int, bool, bool) {
	start := 0
	end := 0
	trimR := false
	trimL := false

	switch align {
	case "LBLeft":
		offSet = 0
		if len(sl) > w {
			end = len(sl)
			if end < 0 {
				end = w
			}
			sl = sl[start:end]
		}
	case "left":
		offSet = 0
		if len(sl) > w {
			trimR = true
			end = len(sl) - w
			if end < 0 {
				end = w
			}
			sl = sl[start : w-1]
		}
	case "right":
		offSet = (w - len(sl))
		end = len(sl)
		if len(sl) > w {
			start = len(sl) - w
			offSet = 0
			trimL = true
		}
		sl = sl[start:end]
	case "center":
		offSet = ((w - len(sl)) / 2)
		if offSet < 0 {
			offSet = 0
		}
		if len(sl) > w {
			start = ((len(sl) - w) / 2)
			end = start + w
			sl = sl[start:end]
			trimL = true
			trimR = true
		}
	default:
		align = "left"
		offSet = 0
		if len(sl) > w {
			trimR = true
			end = len(sl) - w
			if end < 0 {
				end = w
			}
			sl = sl[start : w-1]
		}
	}
	return sl, offSet, trimR, trimL
}

//Draw -
func Draw(x, y, w int, align string, input string) {
	offSet := 0
	trimR := false
	trimL := false
	sl := []rune(input)
	_ , maxY := GetBounds()

	////////////////модуль переноса строки - возможно перенести в отдельную функцию
	if align == "LBLeft" {
		outRows := SplitSubN(input, w)
		for i := range outRows {
			if y <= maxY {
				PrintText(x, y, string(outRows[i]))
				y++
			}
		}
	} else {
		/////////////////////////////////////////////////////////////////////////////////
		sl, offSet, trimR, trimL = trim(sl, w, align, offSet)
		x = x + offSet
		input = string(sl)
		PrintText(x, y, input)
	}
	if trimR == true {
		PrintText(x+w-1, y, "…")
	}
	if trimL == true {
		PrintText(x, y, "…")
	}
}

//ClearScreen -
func ClearScreen(r rune, fgCol, bgCol TColor) {
	termbox.Flush()
	w, h := termbox.Size()
	FillRect(0, 0, w, h, r, fgCol, bgCol)
}

//Flush -
func Flush() {
	termbox.Flush()
}

//MoveCursor -
func MoveCursor(x, y int) {
	cursorX = x
	cursorY = y
	termbox.SetCursor(x, y)
}

// HideCursor -
func HideCursor() {
	termbox.SetCursor(-1, -1)
}

// ShowCursor -
func ShowCursor() {
	termbox.SetCursor(cursorX, cursorY)
}

//SplitSubN - 
func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
