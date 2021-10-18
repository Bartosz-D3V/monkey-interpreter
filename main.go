package main

import (
	"fmt"
	"monkey_interpreter/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		fmt.Print("Could not find system user", err)
		panic(err)
	}
	fmt.Printf("Hello %s. Start by typing monkey code \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
