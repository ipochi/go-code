// To achieve a similar OOPs pattern in Go , we must reduce Type Hierarchies
// rather we must group types based on the behaviour.

// Using the same example of Human , Employee and Manager
// note in this example without using Type Hierarchy , we removed the need for
// Human Type.

package main

import "fmt"

// lets create an interface that denotes the behaviour of any Human Speak.
type Speaker interface {
	Speak()
}

// Create a type of Employee.

type Employee struct {
	Name         string
	Gender       string
	Age          int
	City         string
	Role         string
	Phone        string
	WorkLocation string
}

// Create a Speak method , which would then mean that Employee type would
// satisfy the Speaker interface.

func (e *Employee) Speak() {

	fmt.Println(
		"Hi My name is ", e.Name,
		"I am a ", e.Age,
		" year old ", e.Gender,
		" currently living in ", e.City,
		". I work as a ", e.Role,
		", in ", e.WorkLocation,
		". You can reach me at ", e.Phone)
}

// Create a type Of Manager

type Manager struct {
	Name            string
	Gender          string
	Age             int
	City            string
	Role            string
	Phone           string
	WorkLocation    string
	ManagesDivision string
}

// Creating Speak method so that Speaker interface is satisfied by Manager type.

func (m *Manager) Speak() {
	fmt.Println(
		"Hi My name is ", m.Name,
		"I am a ", m.Age,
		" year old ", m.Gender,
		" currently living in ", m.City,
		". I work as a ", m.Role,
		", in ", m.WorkLocation,
		". I manage the ", m.ManagesDivision, " Division",
		". You can reach me at ", m.Phone)
}

func main() {

	// The good thing about this way is that we can create a reference of our interface
	// and use that to refer our Employee and Manager instances

	// We could not do in the earlier exmample of oopsInGoWithTypeHierarchy

	var e Speaker
	var m Speaker

	e = &Employee{
		Name:         "Alfred Pitch",
		Gender:       "Male",
		Age:          27,
		City:         "London",
		Role:         "Sales Executive",
		Phone:        "334-445-446",
		WorkLocation: "Manchester",
	}

	m = &Manager{
		Name:            "Mark Bell",
		Gender:          "Male",
		Age:             45,
		City:            "Newcastle",
		Role:            "Accounts Manager",
		Phone:           "334-445-560",
		WorkLocation:    "Manchester",
		ManagesDivision: "Accounts",
	}

	e.Speak()
	m.Speak()
}
