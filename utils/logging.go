package utils

// Helped by github.com/BRO3886

import (
	"os"

	"github.com/fatih/color"
)

var (
	// WarnStyle is a bold yellow hue
	WarnStyle = color.New(color.FgHiYellow, color.Bold)
	// ErrorStyle is a bold red hue
	ErrorStyle = color.New(color.FgHiRed, color.Bold)
	// InfoStyle is a bold green hue
	InfoStyle = color.New(color.FgHiCyan, color.Bold)
	// PrintStyle is a white hue
	PrintStyle = color.New(color.FgWhite)
)

var (
	// Warn is a printf function with WarnStyle
	Warn = WarnStyle.Printf
	// Error is a printf function with ErrorStyle
	Error = ErrorStyle.Printf
	// Info is a printf function with InfoStyle
	Info = InfoStyle.Printf
	// Print is a printf function with PrintStyle
	Print = PrintStyle.Printf
)

// PrintF prints the given string with custom color formats
func PrintF(data string, formats ...color.Attribute) {
	style := color.New(formats...)
	style.Printf(data)
}

// ErrorP prints errors with a format and interface like in printf and exits the program.
func ErrorP(format string, a ...interface{}) {
	ErrorStyle.Printf(format, a...)
	os.Exit(1)
}
