package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// ← apne package ko "pb" alias do
	"google.golang.org/protobuf/proto" // ← official library
)

func writeFile(fname string, message proto.Message) {
	out, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("Cant Serialize to bytes ", err)
	}

	errWrite := os.WriteFile(fname, out, 0644)
	if errWrite != nil {
		log.Fatal("Cant Write the file ", errWrite)
	}

	fmt.Println("Written to the file")
}

func readFile(fname string, pb proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("Can't Read the File", err)
		return
	}
	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatal("Could not unmarshal", err)
		return
	}
	fmt.Println(pb)
}
