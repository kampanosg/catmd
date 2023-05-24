package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
)

const (
	Theme string = "dark"
)

var (
	errNotEnoughArgs = errors.New("not all arguments supplied")
)

func main() {

	if len(os.Args) < 2 {
		panic(errNotEnoughArgs)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	out, err := glamour.Render(string(content), Theme)
	if err != nil {
		fmt.Errorf("error parsing the file, %w", err)
		return
	}

	fmt.Println(out)
}
