package main

import (
	"fmt"
	"log"
	"net"

	peoplepb "github.com/shubamakabra/MagnusBox/golang/people"
	"github.com/shubamakabra/MagnusBox/pkg/shop"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting service!")
	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Arrrrgghhh im dying!!", err)
	}
	fmt.Println("0")
	sc := shop.NewController()

	s := grpc.NewServer()
	peoplepb.RegisterCustomerServer(s, sc)

	log.Println("listening on port: ", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("error is: ", err)
	}

}
