package textcolor

import (
	"fmt"
)

// possible colors are:
// red, green, yellow, blue, purple, cyan, gray, white
func Print(text, color string) {
	fmt.Println(setColor(text, color))
}

// possible colors are:
// red, green, yellow, blue, purple, cyan, gray, white
func Change(text, color string) string {
	return setColor(text, color)
}

var colors = map[string]string{
	"red":    "[031m",
	"green":  "[032m",
	"yellow": "[033m",
	"blue":   "[034m",
	"purple": "[035m",
	"cyan":   "[036m",
	"gray":   "[037m",
	"white":  "[038m",
}

func setColor(text, color string) (coleredString string) {
	if a, found := colors[color]; found {
		coleredString = fmt.Sprintf("\033%s%s\033[0m", a, text)
	}
	return coleredString
}
