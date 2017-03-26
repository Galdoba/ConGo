package congo

func mouseTranslate(mouseButton int) string {
	var button string
	switch mouseButton {
	case 65509:
		button = "LMBrelease"
	case 65510:
		button = "RMB"
	case 65511:
		button = "MMB"
	case 65512:
		button = "LMB"
	case 65507:
		button = "MWDown"
	case 65508:
		button = "MWUp"						
	}
	return button
}