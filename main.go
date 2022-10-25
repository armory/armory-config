package main

// func main() {

// 	fmt.Println("Hello World")

// 	testEnv := deploymentEnv.Env{
// 		Name:  "testEnv",
// 		Value: "testValue",
// 	}

// 	// data, err := proto.Marshal(&testEnv)
// 	// if err != nil {
// 	// 	log.Fatal("Marshaling error", err)
// 	// }

// 	fmt.Println(testEnv.Name)

// 	var strArray []*deploymentEnv.Env

// 	var cmdArray []string

// 	strArray = append(strArray, &testEnv)

// 	cmdArray = append(cmdArray, "yes")

// 	testContainer := deploymentEnv.Containers{
// 		Name:    "Test",
// 		Image:   "testImage",
// 		Command: cmdArray,
// 		Env:     strArray,
// 	}
// 	// data_again, err := proto.Marshal(&testContainer)
// 	// if err != nil {
// 	// 	log.Fatal("Marshaling error", err)
// 	// }

// 	var arr []*deploymentEnv.Containers

// 	arr = append(arr, &testContainer)

// 	fmt.Print(testContainer.Command)

// 	testDE := deploymentEnv.DeploymentEnvironment{
// 		Container: arr,
// 	}

// 	fmt.Print("\n------------\n")
// 	fmt.Print(testDE.Container)

// }

import (
	"fmt"

	"github.com/austinthao5/golang_proto_test/cmd/migrate"
	"github.com/austinthao5/golang_proto_test/internal/parser"
)

func main() {

	path := "/Users/austinthao/kustom_github/austin-spin-ftw/proto_tests/Austin.yml"

	test, err := parser.ParseHalConfig(path)

	fmt.Println("TESTING IN MAIN")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(test.DeploymentConfigurations[0].Security.Authn)
	}

	migrate.Execute()
}
