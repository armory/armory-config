package main

import (
	"fmt"
	"log"

	"github.com/austinthao5/golang_proto/config"
	"google.golang.org/protobuf/proto"
)

func main() {

	fmt.Println("Hello World")

	testEnv := pb.Env{
		name:  "testEnv",
		value: "testValue",
	}

	data, err := proto.Marshal(testEnv)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}

	// testContainer := &Containers{
	// 	name: "Test",
	// 	image: "testImage",
	// 	command: "yes",
	// 	env: testEnv,
	// }

	fmt.Println(data)
}

func test() (*config.Env, error) {

}
