package main

import (
	"fmt"
	"strings"
)

// Person ...
// An example structure
type Person struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

func createPerson(name string, sName string, hobbies []string) (p Person) {

	p.Name = name
	p.Surname = sName
	p.Hobbies = hobbies

	return
}

// This method will not be visible outside the package
func (p *Person) toJSONString() string {
	return fmt.Sprintf("{\"Name\":\"%s\",\"Surname\":\"%s\", \"Hobbies\": [\"%s\"]}", p.Name, p.Surname, strings.Join(p.Hobbies, "\",\""))
}

// Jsonfy ...
// This method will not be visible outside the package
func (p *Person) Jsonfy() string {
	return p.toJSONString()
}

func runStruct() {
	var p Person = createPerson("John", "Doe", []string{"Cycling", "Painting"})
	fmt.Println(p)
	fmt.Println("Private method : ", p.toJSONString())
	fmt.Println("Public method : ", p.Jsonfy())
}
