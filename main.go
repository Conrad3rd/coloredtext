package main

import "github.com/conrad3rd/goColor/color"

func main() {
	var text string = "This text is colored"

	color.PrintColoredText(text, "blue")
}
