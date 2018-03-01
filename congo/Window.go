package congo

//	"strings"
import (
	"encoding/hex"
	"strconv"
	"strings"
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
	text  []string
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
	containedText string
	storedText    string
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
	SetSize(int, int)
	SetPosition(int, int)
	GetScrollIndex() int
	SetScrollIndex(int)
	GetStoredRows() int
	GetTitle() string
	WPrint(...interface{})
	WPrintLn(...interface{})
	WGetContent() string
	WSetContent(string)
	WRead() string
	WClear()
	GetID() int
	GetPrintableHeight() int
	GetPrintableWidth() int
	//WOutPut()
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
	//PrintText(1, 5, strconv.Itoa(w.scrollIndex))

}

//WPrint - Кодирует входящий стринг в HEX и добавляет его к (w.containedText) не создает новую строку
func (w *TWindow) WPrint(allData ...interface{}) {
	for _, data := range allData {
		//color := GetFgColor()
		var text string
		switch data.(type) {
		case int:
			text = strconv.Itoa(data.(int))
		case uint8:
			text = ""
		default:
			text = data.(string)
		}

		w.storedText += text
	}
}

//WPrint - Кодирует входящий стринг в HEX и добавляет его к (w.containedText) не создает новую строку
func (w *TWindow) WPrintLn(allData ...interface{}) {
	for _, data := range allData {
		//color := GetFgColor()
		var text string
		switch data.(type) {
		case int:
			text = strconv.Itoa(data.(int))
		case uint8:
			text = ""
		default:
			text = data.(string)
		}
		w.storedText += text //+ "{/N}"
	}
	w.storedText += "{/N}"
}

func (w *TWindow) WDraw() {
	//col := ColorYellow
	FillRect(w.posX, w.posY, w.width, w.height, ' ', GetFgColor(), GetBgColor())
	//Border:
	SetBounds(w.posX+w.width, w.posY+w.height-2*w.checkBorderVisibility())
	if w.borderVisible {
		if w.isActive {
			SetFgColor(ColorGreen)
		}
		DrawBorder(w.posX, w.posY, w.width, w.height, w.border)
		if w.titleVisible {
			if w.isActive {
				SetFgColor(ColorDefault)
			} else {
				SetFgColor(ColorGreen)
			}
			Draw(w.posX+2, w.posY, w.width-4, "left", w.title)

		}
		SetFgColor(ColorDefault)
	}
	winX := w.posX + w.checkBorderVisibility()
	winY := w.posY + w.checkBorderVisibility()
	//////
	tSlice := strings.Split(w.storedText, "")
	if w.scrollIndex > 0 {
		w.scrollIndex = 0
	}
	//w.autoScroll = true
	moveByX := 0
	moveByY := w.scrollIndex
	PrintText(1, 5, strconv.Itoa(w.scrollIndex))
	tag := ""
	readTag := false
	line := 0
	for x := range tSlice {

		if tSlice[x] == "{" {
			readTag = true
			continue
		}
		if tSlice[x] == "}" {

			readTag = false
			switch strings.ToUpper(tag) {
			case "RED":
				SetFgColor(ColorRed)
			case "BLACK":
				SetFgColor(ColorBlack)
			case "BG:RED":
				SetBgColor(ColorRed)
			case "YELLOW":
				SetFgColor(ColorYellow)
			case "GREEN":
				SetFgColor(ColorGreen)
			case "DEFAULT":
				SetFgColor(ColorDefault)
			case "/N":
				moveByY++
				line++
				moveByX = 0
			}
			tag = ""
			continue
		}
		if readTag {
			tag = tag + tSlice[x]
			continue
		}

		if moveByX >= w.GetPrintableWidth() {
			moveByY++
			moveByX = 0
			line++
		}

		if moveByY > w.GetPrintableHeight() {
			break
		}

		/*	PrintText(1, 6, "MOVEBYX =                         ")
			PrintText(1, 7, "MOVEBYy =                         ")
			PrintText(1, 8, "w.GetPrintableHeight() =          ")
			PrintText(1, 9, "moveByY =                         ")
			PrintText(1, 10, "x =                         ")
			PrintText(1, 11, "Lines =                         ")
			PrintText(1, 6, "MOVEBYX = "+strconv.Itoa(moveByX))
			PrintText(1, 7, "Scroll = "+strconv.Itoa(w.scrollIndex))
			PrintText(1, 8, "w.GetPrintableHeight() = "+strconv.Itoa(w.GetPrintableHeight()))
			PrintText(1, 9, "moveByY = "+strconv.Itoa(moveByY))
			PrintText(1, 10, "x = "+strconv.Itoa(x))
			PrintText(1, 11, "Lines = "+strconv.Itoa(line+1))*/
		if winY+moveByY < winY {
			moveByX++
			continue
		}
		if moveByY < 0 {
			continue
		}
		Draw(winX+moveByX, winY+moveByY, w.GetPrintableWidth(), "left", tSlice[x])
		moveByX++
	}
	if w.scrollIndex <= 0-(line-w.GetPrintableHeight()+2) {
		w.scrollIndex = 0 - (line - w.GetPrintableHeight() + 2)
		//panic(2)
	}
	if w.autoScroll {
		w.scrollIndex = 0 - (line - w.GetPrintableHeight() + 2)
	}
	if line > w.GetPrintableHeight() {
		w.putScrollMarker(w.scrollIndex, line)
	}
	SetFgColor(ColorDefault)
	SetBgColor(ColorBlack)
}

//WRead -
func (w *TWindow) WRead() string {
	//decodedRow, _ := hex.DecodeString(row[topRow+i].text)
	text, _ := hex.DecodeString(w.containedText)
	return string(text)
}

//WGetContent -
func (w *TWindow) WGetContent() string {
	return w.containedText
}

//WSetContent -
func (w *TWindow) WSetContent(cont string) {
	w.containedText = cont
}

//WClear - Сносит из HEX в (w.containedText)
func (w *TWindow) WClear() {
	w.containedText = ""
}

func (w *TWindow) putScrollMarker(marker, rowQty int) {
	if marker >= 0 {
		marker = -1
	}

	if marker < w.scrollIndex-w.GetPrintableHeight() {
		marker = w.scrollIndex - w.GetPrintableHeight()
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
	PrintText(w.posX+w.width-w.checkBorderVisibility(), int(float32((0-marker)+w.GetPrintableHeight()-1)*scrollJump)+w.posY, "X")
}

//DeleteWindow -
func (w *TWindow) DeleteWindow() {

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
	//window.storedText = append(window.storedText, "")
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

func insert(s []string, at int, val string) []string {
	// Move all elements of s up one slot
	s = append(s[:at+1], s[at:]...)
	// Insert the new element at the now free position
	s[at] = val
	return s
}
