package main

import (
	"fmt"
	"os"

	"github.com/meliadamian17/atleastitsnotjs/interpreter/repl"
)

func main() {
	fmt.Printf("Hey, it might be shit, but, AtLeastItsNotJS\n")
	repl.Start(os.Stdin, os.Stdout)
}
