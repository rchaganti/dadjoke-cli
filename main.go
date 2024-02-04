package main

import (
	"fmt"

	"github.com/rchaganti/dadjoke-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
