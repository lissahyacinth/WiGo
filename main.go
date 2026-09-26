package main

import (
	"fmt"
	"os"
	"os/user"

	"wigo/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey PL :)\n", user.Name)
	fmt.Printf("Please start writing commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
