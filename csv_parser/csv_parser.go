package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name    string
	Species string
	Day     int
	Status  string
}

func NewEmployee() Employee {
	return Employee{
		Name:    "",
		Species: "",
		Day:     0,
		Status:  "",
	}
}

func main() {
	filename := "data-messy.csv"
	file, err := os.Open(filename)
	var employees []Employee
	if err != nil {
		fmt.Print("ERROR\n", err)
		return
	}
	defer file.Close()

	scanner := csv.NewReader(file)
	//skip header
	_, err = scanner.Read()

	for {
		row, err := scanner.Read()
		if err == io.EOF {
			break
		} else if errors.Is(err, csv.ErrFieldCount) {
			fmt.Printf("Skipping malformed row: %v\n", row)
			continue
		}
		newEmployee := NewEmployee()
		fmt.Println(row)
		if len(row) > 3 {
			newEmployee.Name = strings.TrimSpace(row[0])
			newEmployee.Species = strings.TrimSpace(row[1])
			newEmployee.Day, _ = strconv.Atoi(strings.TrimSpace(row[2]))
			newEmployee.Status = strings.TrimSpace(row[3])
			employees = append(employees, newEmployee)
		}
	}

	harvested := 0
	producing := 0

	for _, emp := range employees {
		if strings.EqualFold("producing", emp.Status) {
			producing += 1
		} else if emp.Status == "harvested" {
			harvested += 1
		}

	}
	fmt.Printf("Harvested: %d\n", harvested)
	fmt.Printf("Producing: %d\n", producing)

}
