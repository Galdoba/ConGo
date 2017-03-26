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
	mouseX 		int
	mouseY 		int
	mouseButton int
	temp string
	mouseColor int

)

func initialize() {
	congo.InitBorders()
	width, height = congo.GetSize()
}

func draw() {
	congo.ClearScreen(' ', fillColorFg, fillColorBg)
	congo.FillRect(width/4, height/4-5, width/2, height/2, '-', mouseColor, fillColorBg)
	congo.Draw(0, 14, 20, 1, strconv.Itoa(counter), fillColorFg, fillColorBg) //выводит таймер и 4 варианта строки
	congo.Draw(width/4, (height/4 - 4), width/2, 0, total, fillColorFg, fillColorBg)
	congo.Draw(width/4, (height/4 - 2), width/2, 1, total, fillColorFg, fillColorBg)
	congo.Draw(width/4, (height/4), width/2, 2, total, fillColorFg, fillColorBg)
	congo.Draw(width/4, (height/4 + 2), width/2, 3, total, fillColorFg, fillColorBg)
	congo.Draw(width/4, (height/4 + 4), width/2, 3, string(temp), fillColorFg, fillColorBg)
	congo.Draw(10, 12, 20, 0, string(key), fillColorFg, fillColorBg)
	congo.PrintText(10, 27, fmt.Sprintf("%v-%T-%s ", info, info, info), fillColorFg, fillColorBg)
	congo.PrintText(10, 3, fmt.Sprintf("%v-%v ", width, height), fillColorFg, fillColorBg)
	congo.DrawBorder(0, 0, width, height, "Double Line", fillColorFg, fillColorBg)
	congo.PrintText(mouseX,mouseY, "█", 2, fillColorBg)
	congo.Flush()
}

func handleEvent(ev congo.IEvent) {
	switch ev.GetType() {
	case "Keyboard":
	counter++
	if ev.GetRune() > 31 {
		temp = string(ev.GetRune())
		key = ev.GetRune()
		
		total = total + temp
	}
	if ev.GetKey() == 32 {
		temp = " " 
		total = total + temp
	}
	if ev.GetKey() == 8 {
		if len(total) > 0 {
			total = total[:len(total)-1]	
		}
	}
	if ev.GetRune() == '4' {
		if fillColorBg == 0 {
			fillColorBg = 2
		} else {
			fillColorBg = 0
		}
	}
		switch ev.GetKey() {
		case 27:
			canClose = true
		}
	case "Resize":
		//_ = ev.GetRune()
		width, height = ev.GetSize()
		/*congo.Flush()
		width, height = congo.GetSize()*/
	
	case "MouseEvent":
		mouseX, mouseY = ev.GetMouseCoords()
		mouseButton = ev.GetMouseButton()
		if mouseButton == 65512 && mouseX >= width/4 && mouseX <= width/4*3  {
			mouseColor = mouseColor + 1
			if mouseColor > 7 {
				mouseColor = 0
			}
		}
		/*} else {
			fillColorBg = 0
		}*/
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
