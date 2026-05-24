package main

import (
	"errors"
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

func doComplex() (*proto.MessageComplex, error) {
	return &proto.MessageComplex{
		OneDummy: &proto.Dummy{
			Id:   4,
			Name: "Rahul Bisht",
		},
		DummyArray: []*proto.Dummy{
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

func getPersonInformation() *proto.PersonalInformation {
	return &proto.PersonalInformation{
		Name:    "Rahul",
		Surname: "Bisht",
		Gender:  proto.GenderInfo_GENDER_MALE,
	}
}

func getOneofs(message any) (any, error) {
	switch x := message.(type) {
	case *proto.Results_Id:
		return x.Id, nil
	case *proto.Results_Message:
		return x.Message, nil
	default:
		return nil, errors.New("Invalid Operation")
	}
}

func doMap() *proto.MapExample {
	return &proto.MapExample{
		Ids: map[string]*proto.IdWrapper{
			"rahul bisht": {
				Id: 123,
			},
		},
	}
}

func main() {

	doSimpleInstance, err := doSimple()
	if err != nil {
		fmt.Print("Error in creating the instance")
		return
	}
	fmt.Println(doSimpleInstance)

	ComplexObject, error := doComplex()
	if error != nil {
		fmt.Println("error occured in line 52")
		return
	}
	fmt.Println(ComplexObject)

	//TODO NEED TO UNDERSTAND , RIGHT NOW HAVE PARTIAL KNOWLEDGE
	personalInformation := getPersonInformation()
	fmt.Println(personalInformation)

	value, err := getOneofs(&proto.Results_Message{Message: "Nice to have you around"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	println(&value)

	valueMap := doMap()
	fmt.Println(valueMap)
}
