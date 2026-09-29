package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Employee struct {
	Name    string
	Species string
	Status  string
}

func openFile(inputfile string) *os.File {
	f, err := os.Open(inputfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	return f
}

func (e Employee) isOnBadList() bool {
	return strings.EqualFold("flagged", e.Status)
}

func (e Employee) checkStatus() {
	if !e.isOnBadList() {
		fmt.Printf("%s: Safe. ...for now\n", e.Name)
	} else {
		fmt.Printf("%s, PIP TIME!\n", e.Name)
	}
}

func main() {
	empData := make(map[string]Employee)
	fileName := flag.String("filename", "registry.csv", "file name to open")
	flag.Parse()
	fmt.Println(*fileName)
	filePointer := openFile(*fileName)
	defer filePointer.Close()

	scanner := csv.NewReader(filePointer)

	_, err := scanner.Read() //skip header
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		row, err := scanner.Read()
		if err == io.EOF {
			break
		}
		newEmployee := Employee{row[1], row[2], row[4]}
		empData[row[0]] = newEmployee
	}
	for _, employee := range empData {
		employee.checkStatus()
	}
}
