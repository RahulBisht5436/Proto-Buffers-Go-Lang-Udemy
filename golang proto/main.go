package main

import (
	"fmt"

	"example.com/proto/proto"
)

func doSimple() (*proto.Simple, error) {
	return &proto.Simple{
		Id:          1,
		IsSimple:    true,
		Name:        "Rahul Bisht",
		SampleLists: []int32{},
	}, nil
}

func main() {

	doSimpleInstance, err := doSimple()
	if err != nil {
		fmt.Print("Error in creating the instance")
		return
	}
	fmt.Println(doSimpleInstance)

}
