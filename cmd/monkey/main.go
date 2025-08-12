package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/farruh4eg/monkey_interpreter/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Now, type in the commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
