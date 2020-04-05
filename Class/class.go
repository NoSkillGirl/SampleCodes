package main

import "fmt"

// Person - struct
type Person struct {
	Name string
	Age  int
}

var persons []*Person

//Class - struct
type Class struct {
	Name   string
	People []*Person
}

func main() {
	person1 := Person{
		Name: "Pooja",
		Age:  25,
	}

	person2 := Person{
		Name: "Praveen",
		Age:  25,
	}
	persons = append(persons, &person1, &person2)
	fmt.Println("Person defined: ", persons)

	result := Class{
		Name:   "class 10",
		People: persons,
	}

	//display
	fmt.Println("Class: ", result)

	// allPersons := getPersons()

	for index, value := range result.People {
		fmt.Printf("Index is: %v, Value is %s\n", index, *value)
	}

}

/* ----------------------------------------------------------------------------------------------------------
//Class - struct
type Class struct {
	Name    string
	Details Person
}

func main() {
	result := Class{
		Name:    "class 10",
		Details: Person{"Sona", 24},
	}

	//display
	fmt.Println("Class Name: ", result.Name)
	fmt.Println("Person Name: ", result.Details.Name)
	fmt.Println("Person Age: ", result.Details.Age)

	result1 := Class{
		Name:    "class 10",
		Details: Person{"Mona", 25},
	}
	fmt.Println("Class Name: ", result1.Name)
	fmt.Println("Person Name: ", result1.Details.Name)

	fmt.Println(result, result1)
}
------------------------------------------------------------------------------------------------------------------*/
/*
func main() {
	person1 := Person{
		Name: "Pooja",
		Age:  25,
	}

	person2 := Person{
		Name: "Praveen",
		Age:  25,
	}

	persons = append(persons, &person1, &person2)

	allPersons := getPersons()

	for index, value := range allPersons {
		fmt.Printf("Index is: %v, Value is %s\n", index, *value)
	}
}
*/
func getPersons() []*Person {
	return persons
}
