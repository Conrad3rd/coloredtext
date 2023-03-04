package printcoloredtext

import "fmt"

// possible colors are
// red, green, yellow, blue, purple, cyan, gray, white
func Paint(text, color string) {
	switch color {
	case "red" : color = "[031m"
	case "green" : color = "[032m"
	case "yellow" : color = "[033m"
	case "blue" : color = "[034m"
	case "purple" : color = "[035m"
	case "cyan" : color = "[036m"
	case "gray" : color = "[037m"
	case "white" : color = "[038m"
	default : color = "[031m"
	}
	fmt.Println("\033" + color + text + "\033" + "[0m")
}
