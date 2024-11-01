package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Person struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Course string  `json:"course"`
	GPA    float64 `json:"gpa"`
}

type Database struct {
	students map[string]Person
}

func (d *Database) Update(name string, data Person) error {
	_, ok := d.students[name]
	if ok {
		return errors.New("student already in database")
	}

	d.students[name] = data
	return nil
}

func (d *Database) LookUp(name string) {
	jsonData, err := json.MarshalIndent(d.students["Andra"], "", "    ")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Biodata: %s\n", jsonData)
}

func main() {
	andra := Person{
		Name:   "Andra",
		Gender: "F",
		Age:    21,
		Course: "Law",
		GPA:    4.10,
	}

	studentsDatabase := Database{
		students: make(map[string]Person),
	}

	err := studentsDatabase.Update("Andra", andra)
	if err != nil {
		log.Fatal(err)
	}

	studentsDatabase.LookUp("Andra")
}
