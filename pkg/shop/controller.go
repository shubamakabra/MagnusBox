package shop

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	peoplepb "github.com/shubamakabra/MagnusBox/golang/people"
	shoppb "github.com/shubamakabra/MagnusBox/golang/shop"
)

type Controller struct {
	shoppb.UnimplementedCustomerServer
	peoplepb.UnimplementedCustomerServer
	pc payment.Client
}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) GetCustomer(ctx context.Context, request *shoppb.GetCustomerRequest) (*shoppb.GetCustomerResponse, error) {
	RID := request.GetId()

	fmt.Println("1")

	if RID == "404" {
		e := &commonpb.Error{
			Status:  "404",
			Message: "Very much errored",
		}
		result := &shoppb.GetCustomerResponse_Error{Error: e} // note this here <------------------------------
		response := &shoppb.GetCustomerResponse{Result: result}
		return response, nil
	}
	pcr := MakeMockCustomer()

	var total float32
	for _, o := range pc.GetOrderHistory() {
		total = total + o.GetTotal()
	}

	cus := customerToPB(Customer{
		ID:        RID,
		PaymentID: "benis",
		Name:      "no pls",
	}, total)
	result := &shoppb.GetCustomerResponse_Customer{Customer: cus}
	response := &shoppb.GetCustomerResponse{Result: result}
	return response, nil
}

func (c *Controller) CreateNewCustomer(ctx context.Context, in *shoppb.Customer) (*shoppb.Customer, error) {
	log.Printf("recieved: %v", in.Name)
	var user_id string = string(rand.Intn(1000))
	return &shoppb.Customer{Name: in.Name, Id: user_id, MoneySpent: in.MoneySpent}, nil
}

func MakeMockCustomer() Customer {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000)
	pid := rand.Intn(1000)
	names := [5]string{"hans", "karl", "nanna", "dieter", "loeffe"}
	c := Customer{
		ID:        strconv.Itoa(id),
		PaymentID: strconv.Itoa(pid),
		Name:      names[rand.Intn(5)],
	}
	return c
}
func customerToPB(input Customer, total float32) *shoppb.Customer {
	r := &shoppb.Customer{
		Id:         input.ID,
		Name:       input.Name,
		MoneySpent: total,
	}
	return r
}
