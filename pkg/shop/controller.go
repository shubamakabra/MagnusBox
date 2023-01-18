package shop

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	commonpb "github.com/shubamakabra/MagnusBox/golang/common"
	peoplepb "github.com/shubamakabra/MagnusBox/golang/people"
	"github.com/shubamakabra/MagnusBox/pkg/structs"
)

type Controller struct {
	peoplepb.UnimplementedPersonServer
}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) GetPerson(ctx context.Context, request *peoplepb.GetPersonRequest) (*peoplepb.GetPersonResponse, error) {
	RID := request.GetId()

	if RID == "404" {
		e := &commonpb.Error{
			Status:  "404",
			Message: "Very much errored",
		}
		result := &peoplepb.GetPersonResponse_Error{Error: e}
		response := &peoplepb.GetPersonResponse{Result: result}
		return response, nil
	}

	var person Person
	result := &peoplepb.GetPersonResponse_Person{}
	for _, p := range structs.MockPeople {
		if strconv.Itoa(int(p.Id)) == RID {
			fmt.Println("Found person!" + person.Name)

			person, err := personToPB(Person{
				ID:    RID,
				Name:  p.Name,
				Phone: p.Phone.GetNumber(),
			})
			if err != nil {
				fmt.Print("errored as a frick")
			}

			result = &peoplepb.GetPersonResponse_Person{Person: person}
			break
		}
	}

	response := &peoplepb.GetPersonResponse{Result: result}
	return response, nil
}

func (c Controller) GetNumber(ctx context.Context, request *peoplepb.GetNumberRequest) (*peoplepb.GetNumberResponse, error) {
	RID := request.GetId()

	if RID == "404" {
		e := &commonpb.Error{
			Status:  "404",
			Message: "Very much errored",
		}
		result := &peoplepb.GetNumberResponse_Error{Error: e}
		response := &peoplepb.GetNumberResponse{Result: result}
		return response, nil
	}

	result := &peoplepb.GetNumberResponse_Number{}
	for _, p := range structs.MockPeople {

		if strconv.Itoa(int(p.Id)) == RID {

			if p.Phone.GetNumber() == "" {

				return nil, errors.New("Could not find phone number for " + p.Name)
			} else {
				fmt.Println("And their number: " + p.Phone.GetNumber())
			}

			nr, err := strconv.Atoi(p.Phone.GetNumber())
			if err != nil {
				return nil, errors.New("Could not convert number to string....?")
			}

			result = &peoplepb.GetNumberResponse_Number{Number: numberToPB(nr)}
			break
		}
	}

	response := &peoplepb.GetNumberResponse{Result: result}
	return response, nil
}

func (c *Controller) CreateNewCustomer(ctx context.Context, in *peoplepb.Person) (*peoplepb.Person, error) {
	log.Printf("recieved: %v", in.Name)
	var user_id int32 = int32(rand.Intn(1000))
	return &peoplepb.Person{Name: in.Name, Id: user_id, Phone: in.Phone}, nil
}

func numberToPB(input int) *peoplepb.Person_PhoneNumber {
	return &peoplepb.Person_PhoneNumber{Number: strconv.Itoa(input)}
}

func personToPB(input Person) (*peoplepb.Person, error) {
	id, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}

	r := &peoplepb.Person{
		Id:    int32(id),
		Name:  input.Name,
		Phone: &peoplepb.Person_PhoneNumber{Number: input.Phone},
	}
	return r, nil
}
