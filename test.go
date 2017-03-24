package main

import (
	"strconv"

	"fmt"

	"github.com/Galdoba/ConGo/congo"
)

var buf chan rune

var (
	total       string
	counter     int
	key         rune
	canClose    bool
	fillColorFg int
	fillColorBg int
	info        interface{}
	width       int
	height      int
)

func initialize() {
	congo.InitBorders()
	width, height = congo.GetSize()
}

func draw() {
	congo.ClearScreen(' ', fillColorFg, fillColorBg)
	congo.FillRect(width/4, height/4, width/2, height/2, '_', fillColorFg, fillColorBg)
	congo.Draw(0, 14, 20, 1, strconv.Itoa(counter), fillColorFg, fillColorBg) //выводит таймер и 4 варианта строки
	congo.Draw(10, 16, 20, 0, total, fillColorFg, fillColorBg)
	congo.Draw(10, 18, 20, 1, total, fillColorFg, fillColorBg)
	congo.Draw(10, 20, 20, 2, total, fillColorFg, fillColorBg)
	congo.Draw(10, 25, 20, 56, total, fillColorFg, fillColorBg)
	congo.Draw(10, 12, 20, 0, string(key), fillColorFg, fillColorBg)
	congo.PrintText(10, 27, fmt.Sprintf("%v-%T ", info, info), fillColorFg, fillColorBg)
	congo.PrintText(10, 3, fmt.Sprintf("%v-%v ", width, height), fillColorFg, fillColorBg)
	congo.DrawBorder(0, 0, 5, 5, "Default", fillColorFg, fillColorBg)
	congo.DrawBorder(0, 0, width, height, "Default", fillColorFg, fillColorBg)
	congo.Flush()
}

func handleEvent(ev congo.IEvent) {
	switch ev.GetType() {
	case "Keyboard":
		switch ev.GetKey() {
		case 27:
			canClose = true
		}
	case "Resize":
		//_ = ev.GetRune()
		width, height = ev.GetSize()
		/*congo.Flush()
		width, height = congo.GetSize()*/
	}

	/*if ev.GetRune() > 31 {
		key = ev.GetRune()
		total = total + string(key)
	}

	if ev.GetRune() == '4' {
		if fillColorBg == 0 {
			fillColorBg = 2
		} else {
			fillColorBg = 0
		}
	}
	counter++*/

}

func main() {
	err := congo.Init()
	if err != nil {
		panic(err)
	}
	defer congo.Close()
	//////////////////////////////////
	initialize()

	//////////////////
	kbd := congo.CreateKeyboard()
	kbd.StartKeyboard()

	//kbd.StopKeyboard()
	//congo.Close()

	for !canClose {
		draw()
		ev := kbd.ReadEvent()
		info = ev
		handleEvent(ev)

		/*if kbd.KeyPressed() {
			ev := kbd.ReadKey()
			handleEvent(ev)
			counter = 0
		}*/
		//total = total + string(key)

	}

	//defer congo.Close()
}
