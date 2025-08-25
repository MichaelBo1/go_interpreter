package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MichaelBo1/go_interpreter/repl"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := user.Current()
	check(err)

	fmt.Printf("Hello %s, This is the Monkey programming language.\n", user.Username)
	repl.Run(os.Stdin, os.Stdout)
}
