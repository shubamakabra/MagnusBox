package file

import (
	"encoding/gob"
	"errors"
	"os"

	"github.com/shubamakabra/MagnusBox/golang/people"
	"google.golang.org/protobuf/proto"
)

func WritePerson(person people.Person) error {
	if !Validate(person) {
		return errors.New("Could not valdate person.")
	}
	byted, err := proto.Marshal(&person)
	if err != nil {
		return err
	}
	err = WritetoFile(byted)
	if err != nil {
		return err
	}

	return nil
}

func Validate(p people.Person) bool {
	if p.Name == "" {
		return false
	}
	if p.Id == 0 {
		return false
	}
	return true
}

func ReadfromFile() {

}

func WritetoFile(input interface{}) error {
	file, err := os.Create("book.gob")
	if err != nil {
		return err
	}

	dEncoder := gob.NewEncoder(file)
	dEncoder.Encode(input)
	return nil
}
