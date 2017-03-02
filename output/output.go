package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Banner outputs a summary of the command & flags
func Banner(name string, flags map[string]string) {
	fmt.Printf("%s:\t\"%s\"\n", changeColor("Command", color.FgBlue), name)
	for key, val := range flags {
		if len(key) > 5 {
			fmt.Printf("%s:\t\"%s\"\n", changeColor(strings.Title(key), color.FgBlue), val)
		} else {
			fmt.Printf("%s:\t\t\"%s\"\n", changeColor(strings.Title(key), color.FgBlue), val)
		}
	}
}

// Error ends the running process with a red error message
func Error(text string) {
	msg := changeColor(text, color.FgRed)
	fmt.Println(msg)
	os.Exit(1)
}

func changeColor(text string, code color.Attribute) string {
	c := color.New(code).SprintFunc()
	return c(text)
}
