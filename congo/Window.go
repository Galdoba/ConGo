package congo

//	"strings"
import (
	"encoding/hex"
	"strconv"
)

//WindowsMap -
var WindowsMap TWindowsMap

//TWindowsMap -
type TWindowsMap struct {
	ByTitle map[string]IWindow
	//byID map[int]IWindow
}

type fmtedRows struct {
	align string
	text  string
	color TColor
}

var id int

//TWindow -
type TWindow struct {
	//owner         *TWindow
	id            int
	posX          int
	posY          int
	width         int
	height        int
	title         string //name?
	titleVisible  bool
	border        string //style?
	borderVisible bool
	vScrollBar    bool
	autoScroll    bool
	scrollIndex   int
	isActive      bool
	outputRange   int
	containedText string
	storedRows    int
	//cursor ???
}

//IWindow -
type IWindow interface {
	InFocus() bool //нужны: (указатель на объект)
	SetFocus(bool)
	WDraw()
	SetBorderVisibility(bool) //Debug only
	SetAutoScroll(bool)
	SetOutputRange(int)
	SetSize(int, int)
	SetPosition(int, int)
	GetScrollIndex() int
	SetScrollIndex(int)
	GetStoredRows() int
	GetTitle() string
	WPrint(string, TColor)
	WPrintLn(string, TColor)
	WRead() string
	WClear()
	GetID() int
	GetPrintableHeight() int
	GetPrintableWidth() int
	//DrawScrollBar()

}

//GetID -
func (w *TWindow) GetID() int {
	return w.id
}

//DrawScrollBar -
func (w *TWindow) DrawScrollBar() {
	//panic("scroll")
	highB := w.posY
	lowB := w.posY + w.storedRows
	delta := lowB - highB
	pointer := w.scrollIndex / delta
	//FillRect(w.posX+w.width-w.checkBorderVisibility(), w.posY+w.checkBorderVisibility(), 1, w.height-2*w.checkBorderVisibility(), 'X', TColor(ColorYellow), TColor(ColorDarkGray))
	PrintText(w.posX+w.width-w.checkBorderVisibility(), w.posY+w.checkBorderVisibility()+pointer, "X")
	PrintText(1, 5, strconv.Itoa(w.scrollIndex))

}

//WPrint - Кодирует входящий стринг в HEX и добавляет его к (w.containedText) не создает новую строку
func (w *TWindow) WPrint(text string, col TColor) {
	var color string
	plainBytes := hex.EncodeToString([]byte(text))
	align := hex.EncodeToString([]byte(string('\x10')))
	//endLine := hex.EncodeToString([]byte(string('\x0c')))
	switch col {
	case ColorGreen:
		color = hex.EncodeToString([]byte(string('\x01')))
	case ColorRed:
		color = hex.EncodeToString([]byte(string('\x02')))
	case ColorYellow:
		color = hex.EncodeToString([]byte(string('\x03')))
	default:
		color = hex.EncodeToString([]byte(string('\x09')))
	}

	w.containedText += align + color + plainBytes// + endLine
	//	w.containedText = strings.TrimRight(w.containedText, "0d")

}

//WPrintLn - Кодирует входящий стринг в HEX и добавляет его к (w.containedText)
func (w *TWindow) WPrintLn(text string, col TColor) {
	var color string
	plainBytes := hex.EncodeToString([]byte(text))
	align := hex.EncodeToString([]byte(string('\x10')))
	endLine := hex.EncodeToString([]byte(string('\x0d')))
	switch col {
	case ColorGreen:
		color = hex.EncodeToString([]byte(string('\x01')))
	case ColorRed:
		color = hex.EncodeToString([]byte(string('\x02')))
	case ColorYellow:
		color = hex.EncodeToString([]byte(string('\x03')))
	default:
		color = hex.EncodeToString([]byte(string('\x09')))
	}
	w.containedText += align + color + plainBytes + endLine
}

//WRead -
func (w *TWindow) WRead() string {
	//decodedRow, _ := hex.DecodeString(row[topRow+i].text)
	text, _ := hex.DecodeString(w.containedText)
	return string(text)
}

//WClear - Сносит из HEX в (w.containedText)
func (w *TWindow) WClear() {
	w.containedText = ""
}

func (w *TWindow) cutter(width int) []fmtedRows {
	//	PrintText(2, 6, w.containedText)
	char := SplitSubN(w.containedText, 2)
	var lines []fmtedRows
	var word string
	var rowAlign string
	var color TColor
	var x string
	//var x fmtedRows
	for i := range char {
		if char[i] == "0c" {
			char[i] = ""
			x = word
			if x == "skdhgfasgfj" {
				panic(1)
			}
			rLine := fmtedRows{rowAlign, word, color}
			lines = append(lines, rLine)
			//lines = lines[:len(lines)-1]
			//word = x
			//lines = append(lines, x)
			//word = ""
			//word = word[:len(word)-6]
			//slice = slice[:len(slice)-1]
			//x := lines[len(lines)-1]
			//y := lines[len(lines)-2]
			//lines = append(lines[:len(lines)-2], lines[len(lines)-1:]...)
			//lines = append(lines[:len(lines)-2], lines[len(lines)-1:]...)

			//x, a := a[0], a[1:]
			//var x int
        	/*x, s = s[0], s[1:]
			sl[len(sl)-1]
        	fmt.Println(x)*/
		}
		//creating words
		//word = x
		if char[i] != "0d" {
			if char[i] == "10" {
				rowAlign = "left"
			} else if char[i] == "11" {
				rowAlign = "right"
			} else if char[i] == "09" {
				color = ColorDefault
			} else if char[i] == "01" {
				color = ColorGreen
			} else if char[i] == "02" {
				color = ColorRed
			} else if char[i] == "03" {
				color = ColorYellow
				/*} else if char[i] == "20" {
				word = word + " "
				if 2*width-len(word) > 0 {
					rLine := fmtedRows{rowAlign, word, color}
					lines = append(lines, rLine)
				}*/
			} else {
				if len(word) > 2*width-1 {
					rLine := fmtedRows{rowAlign, word, color}
					lines = append(lines, rLine)
					word = ""
				}
				word = word + string(char[i])
			}
		} else {
			rLine := fmtedRows{rowAlign, word, color}
			lines = append(lines, rLine)
			word = ""
		}
	}
	rLine := fmtedRows{rowAlign, word, color}
	lines = append(lines, rLine)
	return lines

}

//WDraw -
func (w *TWindow) WDraw() {
	FillRect(w.posX, w.posY, w.width, w.height, ' ', GetFgColor(), GetBgColor())
	//Border:
	SetBounds(w.posX+w.width, w.posY+w.height-2*w.checkBorderVisibility())
	if w.borderVisible {
		if w.isActive {
			SetFgColor(ColorGreen)
		}
		DrawBorder(w.posX, w.posY, w.width, w.height, w.border)
		if w.titleVisible {
			Draw(w.posX+2, w.posY, w.width-4, "left", w.title)
		}
		SetFgColor(ColorDefault)
	}
	//Printable:
	winX := w.posX + w.checkBorderVisibility()
	winY := w.posY + w.checkBorderVisibility()
	winHeight := w.height - 2*w.checkBorderVisibility()
	topRow := w.GetScrollIndex()
	//w.autoScroll = true
	winWidth := w.width - 2*w.checkBorderVisibility()
	row := w.cutter(winWidth)
	w.storedRows = len(row)
	if w.storedRows > winHeight {
		w.vScrollBar = true
	}

	if w.autoScroll { //scrollIndex++ ?? Выводить так чтобы осталось 1-2 пустых строки
		w.SetScrollIndex(w.storedRows - winHeight + 2)
	}
	if len(row) < winHeight {
		w.SetScrollIndex(0)
	}
	if w.vScrollBar == true {
		//w.DrawScrollBar()
	}

	for i := 0; i < winHeight; i++ {
		if topRow+i > len(row)-1 {
			break
		}
		decodedRow, _ := hex.DecodeString(row[topRow+i].text)
		row[topRow+i].text = string(decodedRow)
		fgCol = row[topRow+i].color
		if i >= 0 && i < winHeight {
			Draw(winX, winY, winWidth, row[topRow+i].align, row[topRow+i].text)
		}
		fgCol = ColorDefault
		winY++
		//w.putScrollMarker(topRow+i, len(row))

	}
	if len(row) > w.GetPrintableHeight() {
		w.putScrollMarker(topRow, len(row))
	}
}

func (w *TWindow) putScrollMarker(marker, rowQty int) {
	if marker <= 0 {
		marker = 1
	}
	/*
			http://stackoverflow.com/questions/1406546/calculating-scrollbar-position
			// c is between a and b
		pos = (c-a)/(b-a) // pos is between 0 and 1
		result = pos * (y-x) + x // result is between x and y

	*/

	viewableRatio := float32(w.GetPrintableHeight()) / float32(rowQty) // 1/3 or 0.333333333n
	scrollBarArea := float32(w.GetPrintableHeight())                   // 150px
	thumbHeight := scrollBarArea * viewableRatio                       // 50px

	scrollTrackSpace := float32(rowQty) - float32(w.GetPrintableHeight()) // (600 - 200) = 400
	scrollThumbSpace := float32(w.GetPrintableHeight()) - thumbHeight     // (200 - 50) = 150
	scrollJump := scrollThumbSpace / scrollTrackSpace                     //  (400 / 150 ) = 2.666666666666667
	PrintText(w.posX+w.width-w.checkBorderVisibility(), int(float32(marker+w.GetPrintableHeight()-4)*scrollJump)+w.posY, "X")

	//PrintText(w.posX+w.width-w.checkBorderVisibility(), int(float32(marker + w.GetPrintableHeight() - 2) * scrollJump) + w.posY, "X")
}

//DeleteWindow -
func (w *TWindow) DeleteWindow() {

}

//GetOutputRange -
func (w *TWindow) GetOutputRange() int {
	return w.outputRange
}

//SetOutputRange  -
func (w *TWindow) SetOutputRange(or int) {
	w.outputRange = or
}

//GetPrintableHeight -
func (w *TWindow) GetPrintableHeight() int {
	return w.height - 2*w.checkBorderVisibility()
}

//GetPrintableWidth -
func (w *TWindow) GetPrintableWidth() int {
	return w.width - 2*w.checkBorderVisibility()
}

//GetScrollIndex -
func (w *TWindow) GetScrollIndex() int {
	return w.scrollIndex
}

//SetScrollIndex -
func (w *TWindow) SetScrollIndex(scrIn int) {
	w.scrollIndex = scrIn
}

//GetStoredRows -
func (w *TWindow) GetStoredRows() int {
	return w.storedRows
}

//SetBorderVisibility -
func (w *TWindow) SetBorderVisibility(bv bool) {
	w.borderVisible = bv
}

func (w *TWindow) checkBorderVisibility() int {
	if w.borderVisible {
		return 1
	}
	return 0
}

//GetBorderVisibility -
func (w *TWindow) GetBorderVisibility() bool {
	return w.borderVisible
}

//SetVerticalScrollBar -
func (w *TWindow) SetVerticalScrollBar(bv bool) {
	w.vScrollBar = bv
}

func (w *TWindow) checkVerticalScrollBar() int {
	if w.vScrollBar {
		return 1
	}
	return 0
}

//GetVerticalScrollBar  -
func (w *TWindow) GetVerticalScrollBar() bool {
	return w.vScrollBar
}

//SetTitleVisibility -
func (w *TWindow) SetTitleVisibility(bv bool) {
	w.titleVisible = bv
}

func (w *TWindow) checkTitleVisibility() int {
	if w.titleVisible {
		return 1
	}
	return 0
}

//GetAutoScroll -
func (w *TWindow) GetAutoScroll() bool {
	return w.autoScroll
}

//SetAutoScroll -
func (w *TWindow) SetAutoScroll(bv bool) {
	w.autoScroll = bv
}

//GetTitleVisibility -
func (w *TWindow) GetTitleVisibility() bool {
	return w.titleVisible
}

//GetTitle -
func (w *TWindow) GetTitle() string {
	return w.title
}

//InFocus -
func (w *TWindow) InFocus() bool {
	return w.isActive
}

//SetFocus -
func (w *TWindow) SetFocus(act bool) {
	w.isActive = act
}

//SetSize -
func (w *TWindow) SetSize(width, height int) {
	//w, h := GetSize()
	w.width = width
	w.height = height
}

//SetPosition -
func (w *TWindow) SetPosition(posX, posY int) {
	w.posX = posX
	w.posY = posY
}

//NewWindow -
func NewWindow(posX, posY, width, height int, title, border string) IWindow {
	window := &TWindow{}
	window.id = id
	id++
	window.posX = posX
	window.posY = posY
	window.width = width
	window.height = height
	window.title = title
	window.titleVisible = true
	window.border = border
	window.borderVisible = true

	WindowsMap.ByTitle[window.title] = window
	return window
}

//InitWindowsMap -
func InitWindowsMap() {
	WindowsMap = TWindowsMap{}
	WindowsMap.ByTitle = map[string]IWindow{}
}

//WApply -
func (wm *TWindowsMap) WApply() {
	wm.ByTitle = map[string]IWindow{}
	for _, winExist := range wm.ByTitle {
		wm.ByTitle[winExist.GetTitle()] = winExist
	}
}

//GetNames  -
func (wm *TWindowsMap) GetNames() []string {
	names := make([]string, 0, len(wm.ByTitle))
	for n := range wm.ByTitle {
		names = append(names, n)
	}
	return names
}
