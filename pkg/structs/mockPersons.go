package structs

import (
	"github.com/shubamakabra/MagnusBox/golang/people"
)

var MockPeople = []people.Person{
	{
		Name:  "Hans Kulbert",
		Id:    123,
		Email: "Halbert@aaberg.com",
		Phone: &people.Person_PhoneNumber{Number: "123123"},
	}, {
		Name:  "Hanna Shulbert",
		Id:    321,
		Email: "Hanbert@Shulberg.com",
		Phone: &people.Person_PhoneNumber{Number: "123019501257019275019275015019275012975019275019275012751"},
	},
}
