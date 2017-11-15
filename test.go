package main

import (
	"fmt"
	"time"

	"sort"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/ConGo/utils"
)

var buf chan rune

var (
	total            string
	counter          int
	key              rune
	canClose         bool
	fillColorFg      int
	fillColorBg      int
	info             interface{}
	width            int
	height           int
	mouseX           int
	mouseY           int
	mouseButton      int
	temp             string
	mouseColor       int
	borderNames      []string
	activeBorderName string
	selector         int
	curs             int
	windowLayer      int
	totalWindows     int
	actWin           bool
)

func initialize() {
	congo.InitBorders()
	width, height = congo.GetSize()
	congo.SetTBorder("Default")
	activeBorderName = congo.GetTBorderName()
	congo.InitOutput() // -для инициализации/настройки стартовых цветов
	congo.NewKeyboardAction("Exit Programm", "<esc>", "", func(ev *congo.KeyboardEvent) bool {
		canClose = true
		return true
	})

	congo.NewKeyboardAction("___", "<#65532>", "", func(ev *congo.KeyboardEvent) bool { //KeyboardEvent
		selector = 0

		//borderName = congo.StylePool[selector].GetStyleName()
		return true
	})

	congo.NewKeyboardAction("Choose_selector_position", "5", "", func(ev *congo.KeyboardEvent) bool {
		bord := congo.BorderMap.GetNames()
		sort.Strings(bord)
		if windowLayer == 0 {
			activeBorderName = bord[0]
		}
		return true
	})

	congo.NewKeyboardAction("Choose_next_window", "<TAB>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				congo.WindowsMap.ByTitle[windowList[i]].SetFocus(false)
				if i == len(windowList)-1 {
					i = 0
				} else {
					i++
				}
				congo.ActionMap.SetState(congo.WindowsMap.ByTitle[windowList[i]].GetTitle())
				congo.WindowsMap.ByTitle[windowList[i]].SetFocus(true)
				break
			}
		}

		return true
	})

	congo.NewKeyboardAction("Add_to_window", "3", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				congo.WindowsMap.ByTitle[windowList[i]].WPrintLn("So reading up on Jammers for my Rigger/Decker hybrid (emphasis on the decker part, I was wondering how viable using an area jammer as Matrix defense is. ", congo.ColorDefault)
			}
		}
		return true
	})

	congo.NewKeyboardAction("ActWin/Choose_selector_position", "ActWin/5", "", func(ev *congo.KeyboardEvent) bool {
		color := congo.GetFgColor()
		if color == congo.ColorRed {
			congo.SetFgColor(congo.ColorDefault)
		} else {
			congo.SetFgColor(congo.ColorRed)
		}

		return true
	})

	congo.NewKeyboardAction("Increase_active_window_layer", "+", "", func(ev *congo.KeyboardEvent) bool {
		congo.ActionMap.SetState("ActWin")
		actWin = true
		return true
	})

	congo.NewKeyboardAction("Decrease_active_window_layer", "ActWin/-", "", func(ev *congo.KeyboardEvent) bool {
		congo.ActionMap.SetState("")
		actWin = false
		return true
	})

	congo.NewKeyboardAction("Move_selector_down", "<down>", "", func(ev *congo.KeyboardEvent) bool { //KeyboardEvent
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				index := congo.WindowsMap.ByTitle[windowList[i]].GetScrollIndex()
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Min(index+1, congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows()-congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight()+2))
				if congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows() < congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight() {
					congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(0)
				}
				congo.WindowsMap.ByTitle[windowList[i]].WDraw()

			}
		}
		curs++
		curs = utils.Min(curs, len(styles)-1)
		return true
	})

	congo.NewKeyboardAction("Move_selector_up", "<up>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				congo.WindowsMap.ByTitle[windowList[i]].SetAutoScroll(false)
				index := congo.WindowsMap.ByTitle[windowList[i]].GetScrollIndex()
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Max(index-1, 0))
				congo.WindowsMap.ByTitle[windowList[i]].WDraw()
			}
		}
		curs--
		curs = utils.Max(curs, 0)
		return true
	})

	congo.NewKeyboardAction("Delete_content", "<BACKSPACE>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(0)
				congo.WindowsMap.ByTitle[windowList[i]].WClear()
				congo.WindowsMap.ByTitle[windowList[i]].WDraw()
			}
		}
		curs--
		curs = utils.Max(curs, 0)
		return true
	})

	congo.NewKeyboardAction("ActWin/Move_selector_down", "ActWin/<down>", "", func(ev *congo.KeyboardEvent) bool { //KeyboardEvent
		selector++
		//selector = utils.Min(selector, len(actions)-1)
		return true
	})

	congo.NewKeyboardAction("ActWin/Move_selector_up", "ActWin/<up>", "", func(ev *congo.KeyboardEvent) bool {
		selector--
		selector = utils.Max(selector, 0)
		return true
	})

	congo.NewResizeAction("Resize_Window", "<Resize>", "", func(ev *congo.ResizeEvent) bool { //KeyboardEvent
		congo.Flush()
		width, height = congo.GetSize()
		congo.ClearScreen(' ', congo.GetFgColor(), congo.GetBgColor())
		draw()
		return true
	})

	congo.NewMouseAction("Change_color_up", "<MWup>", "", func(ev *congo.MouseEvent) bool { //KeyboardEvent
		mouseColor++
		if mouseColor > 7 {
			mouseColor = 0
		}
		congo.SetFgColor(congo.ColorRed)
		return true
	})

	congo.NewMouseAction("Change_color_down", "<MWdown>", "", func(ev *congo.MouseEvent) bool { //KeyboardEvent
		mouseColor--
		if mouseColor < 0 {
			mouseColor = 7
		}
		congo.SetFgColor(congo.ColorDefault)
		return true
	})
	congo.InitWindowsMap()

	congo.NewWindow(width/100*60, 0, width/100*40, height, "Window1", "Double Line")
	congo.NewWindow(20, 12, width/2, 34, "Window2", "Double Line")
	congo.WindowsMap.ByTitle["Window1"].WPrintLn("0123456780090155555LN", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].WPrintLn("0123456789000155555LN", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].WPrint("01234567890155555 --- !!", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].WPrintLn("So reading up on Jammers for my Rigger/Decker hybrid (emphasis on the decker part, I was wondering how viable using an area jammer as Matrix defense is. ", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].WPrint("newLine: Kill999999999999999999999999999sadfasdfasdfalasfasjkdhfkl33956", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].SetAutoScroll(true)

	congo.WindowsMap.ByTitle["Window2"].WPrintLn("someText1someText1", congo.ColorDefault)
	//congo.WindowsMap.WApply()
	congo.ActionMap.Apply()
}

var styles []string
var actions []string
var w5Test string

func draw() {

	congo.ClearScreen(' ', congo.GetFgColor(), congo.GetBgColor())
	congo.PrintText(1, 1, fmt.Sprintf("%v-%T-%s ", info, info, info))
	congo.PrintText(1, 3, congo.ActionMap.GetState())
	t := time.Now()
	congo.WindowsMap.ByTitle["Window1"].WPrintLn(fmt.Sprintf("%v    -%v ", info, t), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Window1"].WDraw()
	congo.WindowsMap.ByTitle["Window2"].WDraw()

	congo.Flush()
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

	for !canClose {
		draw()
		ev := kbd.ReadEvent()
		info = ev
		congo.HandleEvent(ev)

	}

	//defer congo.Close()
}
