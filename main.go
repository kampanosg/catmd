package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
)

const (
	WordWrapLen int = 100
)

var errNotEnoughArgs = errors.New("not all arguments supplied")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments, missing filepath")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error: unable to open file %s", os.Args[1])
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error: unable to read file, %s", err.Error())
		return
	}

	r, err := glamour.NewTermRenderer(glamour.WithAutoStyle(), glamour.WithWordWrap(WordWrapLen))
	if err != nil {
		fmt.Println("error: unable to create parser, %s", err.Error())
	}

	out, err := r.Render(string(content))
	if err != nil {
		fmt.Println("error: parsing the file, %s", err.Error())
		return
	}

	fmt.Println(out)
}
