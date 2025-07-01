package struct_basic

import "fmt"

// StructWithNestedTypes demonstrates a struct with nested types
type Address struct {
	Street string // Street address
	City   string // City
	State  string // State
}

// Contact represents a contact with a name and address
type Contact struct {
	Name    string  // Name of the contact
	Age     int     // Age of the contact
	Address Address // Address of the contact
}

// NewContact creates a new Contact instance
func NewContact(name string, age int, street, city, state string) *Contact {
	return &Contact{
		Name: name,
		Age:  age,
		Address: Address{
			Street: street,
			City:   city,
			State:  state,
		},
	}
}

// DisplayContactInfo prints the contact's information
func (c *Contact) DisplayContactInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, %s, %s\n", c.Name, c.Age, c.Address.Street, c.Address.City, c.Address.State)
}

// Example usage of the Contact struct
func StructWithNestedTypesExample() {
	contact := NewContact("Bob", 25, "123 Main St", "Springfield", "IL")
	contact.DisplayContactInfo() // Output: Name: Bob, Age: 25, Address: 123 Main St, Springfield, IL
}
