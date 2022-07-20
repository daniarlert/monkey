package main

import (
	"fmt"
	"log"
	"monkey/pkg/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", u.Username)
	fmt.Println("Feel free to type in any command")
	repl.Start(os.Stdin, os.Stdout)
}
