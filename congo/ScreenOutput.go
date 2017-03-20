package congo

import "github.com/nsf/termbox-go"

//"time"

var cursorX int
var cursorY int
//FillRect - 
func FillRect(x, y, w, h int, bckground rune, fgCol, bgCol int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			termbox.SetCell(x+i, y+j, bckground, termbox.Attribute(fgCol), termbox.Attribute(bgCol))
		}
	}

}
//PrintText - 
func PrintText (x,y int, input string, fgCol, bgCol int) {
	rInput := []rune(input)
	var curentRune rune
	cursorY = y
	for i := range rInput {
		curentRune = rInput[i]
		termbox.SetCell(x, y, curentRune, termbox.Attribute(fgCol), termbox.Attribute(bgCol))
		x++
	}
}



//Output - 
func Output(x, y, w int, align int, input string, fgCol, bgCol int) {
	trimLeft := false
	trimRight := false
	leftBound := x
	switch align {
	default:
		leftBound = x
	case 0: //Left
		leftBound = x
	case 1: //Right
		x = w - len(input)
	case 2: //Center
		x = ((w - len(input)) / 2) + x
	}
	cursorY = y
	if x < leftBound {
		trimLeft = true
	}
	if x + len(input) > w {
		trimRight = true
	}
	if x >= leftBound && x+len(input) <= w {
		PrintText(x,y,input, fgCol, bgCol)
	}
	if trimLeft == true {
		termbox.SetCell(leftBound, y, '…', termbox.Attribute(fgCol), termbox.Attribute(bgCol))
	}
	if trimRight == true {
		termbox.SetCell(w-1, y, '…', termbox.Attribute(fgCol), termbox.Attribute(bgCol))
	}
}
//ClearScreen - 
func ClearScreen() {
	termbox.Flush()
	w, h := termbox.Size()
	FillRect(0, 0, w, h, ' ', 0, 0)
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

