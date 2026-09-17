package main

import (
	"bufio"
	"fmt"
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

func main() {
	filename := "data.csv"
	file, err := os.Open(filename)
	var employees []Employee
	if err != nil {
		fmt.Print("ERROR\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		divided_string := strings.Split(scanner.Text(), ",")
		day_started, _ := strconv.Atoi(divided_string[2])
		newEmployee := Employee{
			Name:    divided_string[0],
			Species: divided_string[1],
			Day:     day_started,
			Status:  divided_string[3],
		}

		employees = append(employees, newEmployee)

	}
	//fmt.Println(employees)
	//fmt.Println(len(employees))
	harvested := 0
	producing := 0

	for _, emp := range employees {
		//fmt.Println(emp.Status)
		if strings.EqualFold("producing", emp.Status) {
			producing += 1
		} else if emp.Status == "harvested" {
			harvested += 1
		}

	}
	fmt.Printf("Harvested: %d\n", harvested)
	fmt.Printf("Producing: %d\n", producing)
}

