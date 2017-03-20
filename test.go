package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Galdoba/ConGo/congo"
)

var buf chan rune

func main() {
	err := congo.Init()
	if err != nil {
		panic(err)
	}
	//defer congo.Close()
	//////////////////////////////////
	kbd := congo.CreateKeyboard()
	kbd.StartKeyboard()
	congo.FillRect(0, 10, 60, 14, '|', 0, 2)

	congo.AddBorderStyle("Style1", "═╗║╝╚╔") //═╗║╝╚╔
	congo.AddBorderStyle("Style2", "000000")
	congo.AddBorderStyle("Style3", "111111")
	congo.DrawBorder(80, 10, 60, 14, "Style1", 0, 0)
	congo.DrawBorder(80, 25, 60, 14, "Style2", 0, 0)
	congo.DrawBorder(80, 40, 60, 14, "Style3", 0, 0)
	congo.DrawBorder(80, 40, 60, 14, "Style1", 0, 0)
	//borderStyleList.DrawBorder(80, 16, 60, 14, "Style1", 0, 0)
	//kbd.StopKeyboard()
	//congo.Close()
	var total string
	i := 0
	key := rune(0)
	for key != 27 {
		if kbd.KeyboardReady() == false {
			fmt.Println("\nKeyboard is not started...")
			os.Exit(2)
		} /* else if i > 140000 { //проверка автовыключения программы
			kbd.StopKeyboard()
			congo.Close()
		}*/

		if kbd.KeyPressed() {
			key = kbd.ReadKey()
			i = 0
			congo.Output(2, 11, 60, 0, "Клавиатура", 0, 0)
			if key > 31 {
				congo.Output(2, 12, 60, 0, string(key), 0, 0)
				total = total + string(key)
			}
		}
		i++
		//total = total + string(key)
		congo.Output(0, 14, 60, 1, strconv.Itoa(i), 0, 0)
		congo.Output(0, 16, 60, 0, total, 0, 0)
		congo.Output(0, 18, 60, 1, total, 0, 0)
		congo.Output(0, 20, 60, 2, total, 0, 0)
		congo.Output(0, 25, 60, 3, total, 0, 0)
		congo.Flush()
	}

	defer congo.Close()
}
