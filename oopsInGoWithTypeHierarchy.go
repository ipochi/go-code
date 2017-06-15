/*
 In this example we will try to create a classic example of OOPs paradigm in Go  ,but by embedding Types.
 This is not the way to do in go , but we'll do it anyway to understandd the bette way in our next example.

 Human , Employee , Manager example

*/

package main

import "fmt"

// Here we have created 3 Types Human , Employee and Manager  ,simple enough

type Human struct {
	Name   string
	Gender string
	Age    int
	City   string
}

// Lets create some methods that are common to a human
// Speak

func (h *Human) Speak() {
	fmt.Println("Hi My name is ", h.Name,
		"I am a ", h.Age,
		" year old ", h.Gender,
		" , from ", h.City)
}

// Human is embedded in Employee
type Employee struct {
	Human
	Role         string
	Phone        string
	WorkLocation string
}

// Lets create a method Speak for Employee  , so that we override the method of Human

func (e *Employee) Speak() {
	fmt.Println("Hi My name is ", e.Name,
		"I am a ", e.Age,
		" year old ", e.Gender,
		" currently living in ", e.City,
		". I work as a ", e.Role,
		", in ", e.WorkLocation,
		". You can reach me at ", e.Phone)
}

// Employee is embedded in Manager as every manager is also an employee

type Manager struct {
	Employee
	ManagesDivision string
}

// Creating a Speak function for Manager

func (m *Manager) Speak() {
	fmt.Println("Hi My name is ", m.Name,
		"I am a ", m.Age,
		" year old ", m.Gender,
		" currently living in ", m.City,
		". I work as a ", m.Role,
		", in ", m.WorkLocation,
		". I manage the ", m.ManagesDivision, " Division",
		". You can reach me at ", m.Phone)
}

func main() {

	// an Employee instance
	e := &Employee{
		Human: Human{
			Name:   "Alfred Pitch",
			Gender: "Male",
			Age:    27,
			City:   "London",
		},
		Role:         "Sales Executive",
		Phone:        "334-445-446",
		WorkLocation: "Manchester",
	}

	m := &Manager{

		Employee: Employee{
			Human: Human{
				Name:   "Mark Bell",
				Gender: "Male",
				Age:    45,
				City:   "Newcastle",
			},
			Role:         "Accounts Manager",
			Phone:        "334-445-560",
			WorkLocation: "Manchester"},

		ManagesDivision: "Accounts",
	}

	// Call method Speak of Employee
	e.Speak()
	// call method speak of Manager
	m.Speak()

}
