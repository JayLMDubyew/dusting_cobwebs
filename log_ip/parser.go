package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ips := make(map[string]int)
	re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	for scanner.Scan() {
		line := scanner.Text()
		addr := re.FindString(line)
		ips[addr]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	fmt.Println(ips)
}

