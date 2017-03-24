package congo

import "github.com/nsf/termbox-go"

//"time"

var cursorX int
var cursorY int

func GetSize() (int, int) {
	return termbox.Size()
}

//FillRect -
func FillRect(x, y, w, h int, bckground rune, fgCol, bgCol int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			termbox.SetCell(x+i, y+j, bckground, termbox.Attribute(fgCol), termbox.Attribute(bgCol))
		}
	}

}

//PrintText -
func PrintText(x, y int, input string, fgCol, bgCol int) {
	rInput := []rune(input)
	var curentRune rune
	cursorY = y
	for i := range rInput {
		curentRune = rInput[i]
		termbox.SetCell(x, y, curentRune, termbox.Attribute(fgCol), termbox.Attribute(bgCol))
		x++
	}
}

func trim(sl []rune, w, align, offSet int) ([]rune, int, bool, bool) {
	start := 0
	end := 0
	trimR := false
	trimL := false
	switch align {
	case 0:
		offSet = 0
		if len(sl) > w {
			trimR = true
			end = len(sl) - w
			if end < 0 {
				end = w
			}
			sl = sl[start : w-1]
		}
	case 1:
		offSet = (w - len(sl))
		end = len(sl)
		if len(sl) > w {
			start = len(sl) - w
			offSet = 0
			trimL = true
		}
		sl = sl[start:end]
	case 2:
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
		align = 0
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
func Draw(x, y, w int, align int, input string, fgCol, bgCol int) {
	offSet := 0
	trimR := false
	trimL := false
	sl := []rune(input)
	sl, offSet, trimR, trimL = trim(sl, w, align, offSet)
	x = x + offSet
	input = string(sl)
	PrintText(x, y, input, fgCol, bgCol)
	if trimR == true {
		PrintText(x+w-1, y, "…", fgCol, bgCol)
	}
	if trimL == true {
		PrintText(x, y, "…", fgCol, bgCol)
	}
}

//ClearScreen -
func ClearScreen(r rune, fgCol, bgCol int) {
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
