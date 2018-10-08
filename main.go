package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name string
	help bool
)

func init() {
	flag.StringVar(&name, "name", "", "input name")
	flag.BoolVar(&help, "help", false, "show help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	msg := NewMessage(name)
	str := msg.String()
	msg.Destroy()

	fmt.Println(str)
}
