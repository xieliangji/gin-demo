package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgGreen).SprintFunc()
	green := color.New(color.FgBlue).Sprintf("hello blue")
	fmt.Println(red("This string is printed in red color!"), green)
}
