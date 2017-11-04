package main

import (
	"fmt"
	"time"
)

// User interface
type User interface {
	PrintName()
	PrintDetails()
}

// Person is cool
type Person struct {
	FirstName, LastName string
	DateOfBirth         time.Time
	Email, Location     string
}

// PrintName prints the person's name
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

// PrintDetails prints person's details
func (p Person) PrintDetails() {
	fmt.Printf("[ Date of birth: %s, Email: %s, Location: %s ]\n", p.DateOfBirth, p.Email, p.Location)
}

// ChangeLocation updates person's location
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

// FakeChangeLocation tries to update person's location, but it can't
func (p Person) FakeChangeLocation(newFakeLocation string) {
	p.Location = newFakeLocation
}

// Admin person
type Admin struct {
	Person

	Roles []string
}

// PrintDetails prints admin details
func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

// Member person
type Member struct {
	Person

	Skills []string
}

// PrintDetails prints member details
func (m Member) PrintDetails() {
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

// Team struct
type Team struct {
	Name, Description string
	Users             []User
}

// GetDetails prints team's information
func (t Team) GetDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

func main() {
	var vic Person
	vic.FirstName = "Victor"
	vic.LastName = "Quiroz Castro"
	vic.DateOfBirth = time.Date(1989, time.March, 01, 0, 0, 0, 0, time.UTC)
	vic.Email = "vic@foo.com"
	vic.Location = "Berlin"

	vic.ChangeLocation("Guadalajara")

	julie := Person{
		FirstName:   "Julieta",
		LastName:    "Campos",
		DateOfBirth: time.Date(1989, time.February, 19, 0, 0, 0, 0, time.UTC),
		Email:       "julie@foo.com",
		Location:    "Berlin",
	}

	julie.FakeChangeLocation("Shina")

	adminVic := Admin{
		Person: Person{
			FirstName:   "Victor",
			LastName:    "Quiroz",
			DateOfBirth: time.Date(1988, time.December, 18, 0, 0, 0, 0, time.UTC),
			Email:       "victorhqc@gmail.com",
			Location:    "Berlin",
		},
		Roles: []string{"Manage Team", "Manage Tasks"},
	}

	memberJulie := Member{
		Person: julie,
		Skills: []string{"JavaScript", "Nintendo Switch"},
	}

	team := Team{
		Name:        "Pushkin",
		Description: "Awesome team",
		Users:       []User{adminVic, memberJulie},
	}

	team.GetDetails()
}
