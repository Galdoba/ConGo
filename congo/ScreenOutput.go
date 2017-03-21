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

func cutInput(sl []rune, w, offSet int) []rune {
	if len(sl) > w {
		end := w // + offSet
		//start := end - w

		sl = sl[0:end]
	}
	return sl
}

//Draw -
func Draw(x, y, w int, align int, input string, fgCol, bgCol int) {
	//sl := []rune(input)
	/*offSet := 0
	switch align {
	default:
		offSet = 0
	case 0:
		offSet = 0
	case 1:
		offSet = w - len(input) - 1
	case 2:
		offSet = ((w - len(input) - 1) / 2)
	}
	x = x + offSet
	sl = cutInput(sl, w, offSet)*/
	//input = string(sl)

	//PrintText(x, y, input, fgCol, bgCol)

	//trimLeft := false
	//trimRight := false
	offSet := 0
	switch align {
	default:
		offSet = 0
	case 0:
		offSet = 0
	case 1:
		offSet = w - len(input)
	case 2:
		offSet = ((w - len(input) - 1) / 2)
	}
	x = x + offSet
	sl := []rune(input)
	sl = cutInput(sl, w, offSet)
	input = string(sl)

	if x < offSet {
		//trimLeft = true
	}
	if x+len(input) > w {
		//trimRight = true
	}
	/*if x >= offSet && x+len(input) <= w {
		PrintText(x, y, input, fgCol, bgCol)
	}*/
	PrintText(x, y, input, fgCol, bgCol)
	/*if trimLeft == true {
		sl := []rune(input)
		input = string(sl)
		sl = append(sl[len(sl)-leftBound:], sl[:len(sl)-w+leftBound]...)
		PrintText(x, y, input, fgCol, bgCol)
		PrintText(leftBound, y, "…", fgCol, bgCol)
	}
	if trimRight == true {

		sl = append(sl[:0], sl[:w]...)
		input = string(sl)
		PrintText(x, y, input, fgCol, bgCol)
		PrintText(offSet, y, "…", fgCol, bgCol)
	}*/
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
