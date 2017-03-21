package main

import (
	"strconv"

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
)

func initialize() {
	congo.InitBorders()
	//congo.AddBorderStyle("Default", "═╗║╝╚╔")
}

func draw() {
	congo.ClearScreen(' ', fillColorFg, fillColorBg)
	congo.FillRect(3, 10, 60, 20, '_', fillColorFg, fillColorBg)
	congo.Draw(0, 14, 60, 1, strconv.Itoa(counter), fillColorFg, fillColorBg) //выводит таймер и 4 варианта строки
	congo.Draw(3, 16, 60, 0, total, fillColorFg, fillColorBg)
	congo.Draw(3, 18, 60, 1, total, fillColorFg, fillColorBg)
	congo.Draw(3, 20, 60, 2, total, fillColorFg, fillColorBg)
	congo.Draw(3, 25, 60, 56, total, fillColorFg, fillColorBg)
	congo.Draw(3, 12, 60, 0, string(key), fillColorFg, fillColorBg)
	congo.Flush()
}

func handleEvent(ev rune) {
	if ev > 31 {
		key = ev
		total = total + string(key)
	}
	if ev == 27 {
		canClose = true
	}
	if ev == '4' {
		if fillColorBg == 0 {
			fillColorBg = 2
		} else {
			fillColorBg = 0
		}
	}
	counter++

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
		ev := kbd.ReadKey()
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
