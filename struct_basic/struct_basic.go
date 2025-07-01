package struct_basic

import (
	"fmt"
)

// Person represents a person with a name and age
type Person struct {
	Name string // Name of the person
	Age  int    // Age of the person
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Greet returns a greeting message from the person
func (p *Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// UpdateAge updates the age of the person
func (p *Person) UpdateAge(newAge int) {
	if newAge < 0 {
		fmt.Println("Age cannot be negative.")
		return
	}
	p.Age = newAge
}

// DisplayInfo prints the person's information
func (p *Person) DisplayInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// Example usage of the Person struct
func StructBasicExample() {
	person := NewPerson("Alice", 30)
	fmt.Println(person.Greet()) // Output: Hello, my name is Alice and I am 30 years old.

	person.UpdateAge(31)
	person.DisplayInfo() // Output: Name: Alice, Age: 31

	person.UpdateAge(-5) // Output: Age cannot be negative.

	person.DisplayInfo() // Output: Name: Alice, Age: 31
}
