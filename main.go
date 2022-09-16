package main

import (
	"config/deploymentconfigurations"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	fmt.Println("Hello World")

	testEnv := deploymentconfigurations.Env{
		name:  "testEnv",
		value: "testValue",
	}

	data, err := proto.Marshal(testEnv)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}

	fmt.Println(data)

	testContainer := deploymentconfigurations.Containers{
		name:    "Test",
		image:   "testImage",
		command: "yes",
		env:     testEnv,
	}
	data_again, err := proto.Marshal(testContainer)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}

	fmt.Print(data_again)

}
