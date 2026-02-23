package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContract(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Println("Contact already exists")
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact: %+v\n", newContact)
}

func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]

	if exists {
		return &contactList[index]
	}

	return nil
}

func ListContacts() {
	fmt.Println("ListContacts")
	if len(contactList) == 0 {
		fmt.Println("No contacts")
		return
	}

	for _, contact := range contactList {
		fmt.Printf("Contact: %+v\n", contact)
	}
}

func main() {
	addContract("Alice wonderland", "alice@gmail.com", "111-222")
	addContract("Bob the builder", "bob@gmail.com", "111-222")
	addContract("Charlie brown", "charlie@gmail.com", "111-222")
	addContract("Bob the builder", "bob@gmail.com", "111-222")

	ListContacts()

	bob := findContact("Bob the builder")
	if bob == nil {
		fmt.Println("No Bob contact found")
	} else {
		fmt.Println("Bob contact found.", bob.Name)
	}
}
