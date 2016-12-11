package main

import (
	"fmt"
	"os"
)

func main() {
	run(os.Args[1:])
}

func run(args []string) {
	for _, v := range args {
		fmt.Println(v)
	}
}
