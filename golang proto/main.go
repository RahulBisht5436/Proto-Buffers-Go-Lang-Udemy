package main

import (
	"errors"
	"fmt"
	"log"

	pb "example.com/proto/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() (*pb.Simple, error) {
	return &pb.Simple{
		Id:          1,
		IsSimple:    true,
		Name:        "Rahul Bisht",
		SampleLists: []int32{},
	}, nil
}

func doComplex() (*pb.MessageComplex, error) {
	return &pb.MessageComplex{
		OneDummy: &pb.Dummy{
			Id:   4,
			Name: "Rahul Bisht",
		},
		DummyArray: []*pb.Dummy{
			{
				Id:   1,
				Name: "sheetal Bisht",
			},
			{
				Id:   2,
				Name: "Pareshwari Bisht",
			},
			{
				Id:   3,
				Name: "Kamal Bisht",
			},
		},
	}, nil
}

func getPersonInformation() *pb.PersonalInformation {
	return &pb.PersonalInformation{
		Name:    "Rahul",
		Surname: "Bisht",
		Gender:  pb.GenderInfo_GENDER_MALE,
	}
}

func getOneofs(message any) (any, error) {
	switch x := message.(type) {
	case *pb.Results_Id:
		return x.Id, nil
	case *pb.Results_Message:
		return x.Message, nil
	default:
		return nil, errors.New("Invalid Operation")
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"rahul bisht": {
				Id: 123,
			},
		},
	}
}

func doFile(p proto.Message) {
	path := "simplePlan.bin"

	writeFile(path, p)

	message := &pb.Simple{}
	readFile(path, message)

	fmt.Println("Read from file:", message)
}

func main() {

	// doSimpleInstance, err := doSimple()
	// if err != nil {
	// 	fmt.Print("Error in creating the instance")
	// 	return
	// }
	// fmt.Println(doSimpleInstance)

	// ComplexObject, error := doComplex()
	// if error != nil {
	// 	fmt.Println("error occured in line 52")
	// 	return
	// }
	// fmt.Println(ComplexObject)

	// //TODO NEED TO UNDERSTAND , RIGHT NOW HAVE PARTIAL KNOWLEDGE
	// personalInformation := getPersonInformation()
	// fmt.Println(personalInformation)

	// value, err := getOneofs(&proto.Results_Message{Message: "Nice to have you around"})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// println(&value)

	// valueMap := doMap()
	// fmt.Println(valueMap)

	simpleObj, err := doSimple()
	if err != nil {
		log.Fatal(err)
	}
	doFile(simpleObj)
}
