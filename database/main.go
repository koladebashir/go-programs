package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (d *Database) Add(name string, data Person) error {
	_, ok := d.students[name]
	if ok {
		return errors.New("student already in database")
	}

	d.students[name] = data
	fmt.Printf("Added %v to database\n", name)
	return nil
}

func (d *Database) LookUp(name string) error {
	if _, ok := d.students[name]; !ok {
		return errors.New("student not in database")
	}

	jsonData, err := json.MarshalIndent(d.students[name], "", "    ")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("Biodata: %s\n", jsonData)
	return nil
}

func (d *Database) Delete(name string) error {
	if _, ok := d.students[name]; !ok {
		return errors.New("can not delete record as student does not exist")
	}

	delete(d.students, name)
	fmt.Printf("Deleted %v from database\n", name)
	return nil
}

func main() {
	studentsDatabase := Database{
		students: make(map[string]Person),
	}

	andra := Person{
		Name:   "Andra",
		Gender: "F",
		Age:    21,
		Course: "Law",
		GPA:    4.10,
	}


	err := studentsDatabase.Add("Andra", andra)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = studentsDatabase.LookUp("Andra")
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	err = studentsDatabase.Delete("Andra")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = studentsDatabase.LookUp("Andra")
	if err != nil {
		fmt.Println("Error:", err)
	}

}
