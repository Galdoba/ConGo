package congo

import "time"
import "strconv"

//IEvent -
type IEvent interface {
	GetEventTime() time.Time
	GetEventType() string
	GetTrigger() string
	/*GetKey() int
	GetRune() rune
	GetWidth() int
	GetHeight() int
	GetSize() (int, int)
	GetMouseCoords() (int, int)
	GetMouseButton() int*/
}

//AbstractEvent -
type AbstractEvent struct {
	eventTime time.Time
}

//KeyboardEvent -
type KeyboardEvent struct {
	AbstractEvent
	key int
	ch  rune
}

//MouseEvent -
type MouseEvent struct {
	AbstractEvent
	mouseButton int
	mouseX      int
	mouseY      int
}

//ResizeEvent -
type ResizeEvent struct {
	AbstractEvent
	width  int
	height int
}

//GetEventType -
func (ev *AbstractEvent) GetEventType() string {
	panic("Abstract Func Call")
	return ""
}

//GetEventTime -
func (ev *AbstractEvent) GetEventTime() time.Time {
	return ev.eventTime
}

//GetTrigger -
func (ev *AbstractEvent) GetTrigger() string {
	panic("Abstract Func Call")
	return ""
}

////////////////////////////////////////////

//GetEventType -
func (ev *KeyboardEvent) GetEventType() string {
	return "Keyboard"
}

func (ev *KeyboardEvent) String() string {
	return ev.GetTrigger()
}

//GetTrigger -
func (ev *KeyboardEvent) GetTrigger() string {
	if ev.GetRune() == 0 {
		trigger, ok := keyTranslateMap[ev.GetKey()]
		if !ok {
			trigger = "<#" + strconv.Itoa(ev.GetKey()) + ">"
		}
		return trigger
	}
	return string(ev.GetRune())
}

//GetKey -
func (ev *KeyboardEvent) GetKey() int {
	return ev.key
}

//GetRune -
func (ev *KeyboardEvent) GetRune() rune {
	return ev.ch
}

/*func (ev *KeyboardEvent) GetWidth() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *KeyboardEvent) GetHeight() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *KeyboardEvent) GetSize() (int, int) {
	panic("Abstract Func Call")
	return 0, 0
}

func (ev *KeyboardEvent) GetMouseCoords() (int, int) {
	panic("Abstract Func Call")
	return 0, 0
}

func (ev *KeyboardEvent) GetMouseButton() int {
	panic("Abstract Func Call")
	return 0
}*/

/////////////////////////////////////////////////

//GetEventType -
func (ev *ResizeEvent) GetEventType() string {
	return "Resize"
}

//GetTrigger -
func (ev *ResizeEvent) GetTrigger() string {
	return "<Resize>"
}

//GetWidth -
func (ev *ResizeEvent) GetWidth() int {
	return ev.width
}

//GetHeight -
func (ev *ResizeEvent) GetHeight() int {
	return ev.height
}

//GetSize -
func (ev *ResizeEvent) GetSize() (int, int) {
	return ev.width, ev.height
}

/*func (ev *ResizeEvent) GetKey() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *ResizeEvent) GetRune() rune {
	panic("Abstract Func Call")
	return ' '
}

func (ev *ResizeEvent) GetMouseCoords() (int, int) {
	panic("Abstract Func Call")
	return 0, 0
}

func (ev *ResizeEvent) GetMouseButton() int {
	panic("Abstract Func Call")
	return 0
}*/

//MOUSE///////////////////////////////////////

//GetEventType -
func (ev *MouseEvent) GetEventType() string {
	return "MouseEvent"
}

//GetTrigger -
func (ev *MouseEvent) GetTrigger() string {
	trigger, ok := mouseTranslateMap[ev.GetMouseButton()]
	if !ok {
		trigger = "<#" + strconv.Itoa(ev.GetMouseButton()) + ">"
	}
	return trigger
}

//GetMouseCoords -
func (ev *MouseEvent) GetMouseCoords() (int, int) {
	return ev.mouseX, ev.mouseY
}

//GetMouseButton -
func (ev *MouseEvent) GetMouseButton() int {
	return ev.mouseButton
}

/*func (ev *MouseEvent) GetWidth() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *MouseEvent) GetHeight() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *MouseEvent) GetSize() (int, int) {
	panic("Abstract Func Call")
	return 0, 0
}

func (ev *MouseEvent) GetKey() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *MouseEvent) GetRune() rune {
	panic("Abstract Func Call")
	return ' '
}
*/
/////////////////////////////////////////Key translate map

var keyTranslateMap = map[int]string{
	8:     "<BACKSPACE>",
	9:     "<TAB>",
	13:     "<ENTER>",
	27:    "<esc>",
	32:    "<space>",
	65535: "<F1>",
	65534: "<F2>",
	65533: "<F3>",
	65524: "<F12>",
	65517: "<up>",
	65516: "<down>",
	65514: "<right>",
	
}

var mouseTranslateMap = map[int]string{
	
	65512: "<LMB>",
	65511: "<MMB>",
	65510: "<RMB>",
	65509: "<MBRelese>",
	65508: "<MWup>",
	65507: "<MWdown>",
}
