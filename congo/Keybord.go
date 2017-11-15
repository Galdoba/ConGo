package congo

import "github.com/nsf/termbox-go"
import "fmt"
import "os"
import "time"

var buf chan IEvent
var congoInitiated bool
var keyboard *Keyboard

// Keyboard -
type Keyboard struct {
	initiated bool
	working   bool
}

// Init -
func Init() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	congoInitiated = true
	buf = make(chan IEvent, 16)
	go iLoop()
	return err
}

//Close -
func Close() {

	time.Sleep(1)
	termbox.Close()
}

//CreateKeyboard -
func CreateKeyboard() *Keyboard {
	if congoInitiated == false {
		fmt.Println("Congo is not initiated...")
		os.Exit(3)
	}
	if keyboard != nil {
		fmt.Println("Keyboard already initiated...")
		os.Exit(3)
	}

	keyboard = new(Keyboard)
	keyboard.initiated = true
	return keyboard
}

func GetKeyboard() *Keyboard {
	return keyboard
}

// StartKeyboard -
func (kbd *Keyboard) StartKeyboard() {
	if kbd.initiated == false {
		fmt.Println("ERROR! Keyboard is not initiated...  see StartKeyboard()")
		os.Exit(1)
	}
	kbd.working = true
}

// StopKeyboard -
func (kbd *Keyboard) StopKeyboard() {
	if kbd.working == false {
		fmt.Println("ERROR! Keyboard already stopped...  see StopKeyboard()")
		os.Exit(1)
	}
	kbd.working = false
}

//KeyboardReady -
func (kbd *Keyboard) KeyboardReady() bool {
	if kbd.initiated == true {
		if kbd.working == true {
			return true
		} else if kbd.working == false {
			panic("\nKeyboard not working...")
		}
	} else {
		panic("\nKeyboard not initiated...")
	}
	return false
}

/*func runeTranslator() rune {
	var key rune
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		key = rune(ev.Key)
		if ev.Ch != 0 {
			key = ev.Ch
		}
	}
	return key
}*/

// ReadEvent -
func (*Keyboard) ReadEvent() IEvent {
	return <-buf
}

// KeyPressed -
func (kbd *Keyboard) KeyPressed() bool {
	if kbd.working == false {
		panic("Keyboard not working...   see 'KeyPressed()'")
	}
	return len(buf) > 0
}

func iLoop() {
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			event := &KeyboardEvent{}
			event.eventTime = time.Now()
			event.key = int(ev.Key)
			event.ch = ev.Ch
			buf <- event
		case termbox.EventResize:
			event := &ResizeEvent{}
			event.eventTime = time.Now()
			event.width = ev.Width
			event.height = ev.Height
			buf <- event
		case termbox.EventMouse:
			event := &MouseEvent{}
			event.eventTime = time.Now()
			event.mouseX = ev.MouseX
			event.mouseY = ev.MouseY
			event.mouseButton = int(ev.Key)
			buf <- event
		}

	}
}
