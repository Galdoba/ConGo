package main

import (
	"fmt"
	"sort"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/ConGo/utils"
)

var buf chan rune

//var Selector int

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
	Selector         int
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
		Selector = 0

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
				congo.WindowsMap.ByTitle[windowList[i]].WPrintLn("So reading up on Jammers for my Rigger/Decker hybrid (emphasis on the decker part, I was wondering how viable using an area jammer as Matrix defense is. ")
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
				/*congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Min(index+1, congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows()-congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight()+2))
				if congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows() < congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight() {
					congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(0)

				}*/
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(index - 1)
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
				//congo.WindowsMap.ByTitle[windowList[i]].SetAutoScroll(false)
				index := congo.WindowsMap.ByTitle[windowList[i]].GetScrollIndex()
				//congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Max(index-1, 0))
				congo.WindowsMap.ByTitle[windowList[i]].SetAutoScroll(false)
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(index + 1)
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
		Selector++
		Selector = utils.Min(Selector, len(actions)-1)
		return true
	})

	congo.NewKeyboardAction("ActWin/Move_selector_up", "ActWin/<up>", "", func(ev *congo.KeyboardEvent) bool {
		Selector--
		Selector = utils.Max(Selector, 0)
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

	w1 := congo.NewWindow(width/100*60, 0, width/100*40, height, "Window1", "Double Line")
	w2 := congo.NewWindow(20, 12, 40, 15, "Window2", "Double Line")

	congo.NewWindow(30, 00, 40, 48, "Window3", "Double Line")
	w1.SetFocus(true)
	w2.SetFocus(true)
	w2.SetAutoScroll(true)
	congo.WindowsMap.ByTitle["Window3"].SetAutoScroll(true)
	for i := 0; i < 150; i++ {
		congo.WindowsMap.ByTitle["Window3"].WPrintLn(12345, "{RED}", 67890, "{DEFAULT}")
		/*congo.WindowsMap.ByTitle["Window3"].WPrint2("{RED}Row1_Word1VERY_LOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOONG")
		congo.WindowsMap.ByTitle["Window3"].WPrint2("{YELLOW}Row2_Word1", "{RED}NEws{BLACK}{BG:RED}tring")
		congo.WindowsMap.ByTitle["Window3"].WPrintLn2("R{BLACK}ow2_{red}Wor{green}d2")
		congo.WindowsMap.ByTitle["Window3"].WPrint2("Row3_Word1")
		congo.WindowsMap.ByTitle["Window3"].WPrint2("Row3_Word2{/N}")
		congo.WindowsMap.ByTitle["Window3"].WPrint2("Row3_Word3")
		congo.WindowsMap.ByTitle["Window3"].WPrint2(999)
		congo.WindowsMap.ByTitle["Window3"].WPrint2("VERY_LOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOONG STRING")*/
		draw()
	}

	//congo.WindowsMap.ByTitle["Window1"].SetAutoScroll(true)

	//congo.WindowsMap.ByTitle["Window2"].WPrintLn("someText1", congo.ColorDefault)
	//congo.WindowsMap.ByTitle["Window2"].WPrintLn("qqqqqqqqqqqqqqqq1164654e", congo.ColorYellow)
	//congo.WindowsMap.ByTitle["Window2"].WPrintLn("3333333334465464646464646464444444444444444444646464644444444444444444444444444444444", congo.ColorRed)
	//congo.WindowsMap.ByTitle["Window2"].WPrintLn("someText1someText1", congo.ColorDefault)
	//congo.WindowsMap.WApply()
	congo.ActionMap.Apply()
}

var styles []string
var actions []string
var w5Test string

func draw() {

	congo.ClearScreen(' ', congo.GetFgColor(), congo.GetBgColor())
	//congo.PrintText(1, 1, fmt.Sprintf("%v-%T-%s ", info, info, info))
	//congo.PrintText(1, 3, congo.ActionMap.GetState())
	//congo.PrintText(5, 0, strconv.Itoa(Selector))
	//t := time.Now()
	//congo.WindowsMap.ByTitle["Window2"].WPrintLn("-", congo.ColorDefault)

	//congo.WindowsMap.ByTitle["Window3"].WPrint(fmt.Sprintf("%v", info), congo.ColorDefault)

	//	congo.WindowsMap.ByTitle["Window1"].WDraw()
	//congo.WindowsMap.ByTitle["Window2"].WDraw()
	congo.WindowsMap.ByTitle["Window3"].WDraw()
	//congo.WindowsMap.ByTitle["Window4"].WDraw2()
	//congo.WindowsMap.ByTitle["Window2"].WPrint(fmt.Sprintf("%v", info), congo.ColorDefault)
	//congo.WindowsMap.ByTitle["Window2"].WPrint(fmt.Sprintf("%v", info), congo.ColorDefault)
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
		congo.WindowsMap.ByTitle["Window3"].WPrint(fmt.Sprintf("%v", info))
	}

	//defer congo.Close()
}
