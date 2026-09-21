package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	port := flag.Int("port", 8000, "port to listen on")
	offset := flag.Int("offset", 32, "offset, if incrementing")
	if len(os.Args) < 2 {
		fmt.Println("Not enough args.")
		os.Exit(1)
	}
	flag.Parse()

	fmt.Println(*port + *offset)
	os.Exit(0)

}
