package main

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func main() {
	in := `# Hello World

This is a simple example of Markdown rendering with Glamour!
Check out the [other examples](https://github.com/charmbracelet/glamour/tree/master/examples) too.

Bye!
`

	out, err := glamour.Render(in, "dark")
	if err != nil {
		fmt.Errorf("error, %s", err.Error())
		return
	}
	fmt.Print(out)
}
