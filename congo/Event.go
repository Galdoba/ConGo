package congo

type IEvent interface {
	GetKey() int
	GetRune() rune
	GetWidth() int
	GetHeight() int
	GetSize() (int, int)
	GetType() string
}

type KeyboardEvent struct {
	key int
	ch  rune
}

type ResizeEvent struct {
	width  int
	height int
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

func (ev *KeyboardEvent) GetType() string {
	return "Keyboard"
}

/////////////////////////////////////////////////

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

func (ev *ResizeEvent) GetType() string {
	return "Resize"
}
