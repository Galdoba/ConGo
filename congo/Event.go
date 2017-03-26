package congo

type IEvent interface {
	GetKey() int
	GetRune() rune
	GetWidth() int
	GetHeight() int
	GetSize() (int, int)
	GetType() string
	GetMouseCoords() (int, int)
	GetMouseButton() int
}

type KeyboardEvent struct {
	key int
	ch  rune
}

type MouseEvent struct {
	mouseButton int
	mouseX int
	mouseY int
}

type ResizeEvent struct {
	width  int
	height int
}

func (ev *KeyboardEvent) GetType() string {
	return "Keyboard"
}

func (ev *KeyboardEvent) GetKey() int {
	return ev.key
}

func (ev *KeyboardEvent) GetRune() rune {
	return ev.ch
}

func (ev *KeyboardEvent) GetWidth() int {
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
	return 0,0
}

func (ev *KeyboardEvent) GetMouseButton() int {
	panic("Abstract Func Call")
	return 0
}

/////////////////////////////////////////////////

func (ev *ResizeEvent) GetType() string {
	return "Resize"
}

func (ev *ResizeEvent) GetWidth() int {
	return ev.width
}

func (ev *ResizeEvent) GetHeight() int {
	return ev.height
}

func (ev *ResizeEvent) GetSize() (int, int) {
	return ev.width, ev.height
}

func (ev *ResizeEvent) GetKey() int {
	panic("Abstract Func Call")
	return 0
}

func (ev *ResizeEvent) GetRune() rune {
	panic("Abstract Func Call")
	return ' '
}

func (ev *ResizeEvent) GetMouseCoords() (int, int) {
	panic("Abstract Func Call")
	return 0,0
}

func (ev *ResizeEvent) GetMouseButton() int {
	panic("Abstract Func Call")
	return 0
}
//MOUSE///////////////////////////////////////



func (ev *MouseEvent) GetType() string {
	return "MouseEvent"
}

func (ev *MouseEvent) GetMouseCoords() (int, int) {
	return ev.mouseX, ev.mouseY
}

func (ev *MouseEvent) GetMouseButton() int {
	return ev.mouseButton
}

func (ev *MouseEvent) GetWidth() int {
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
