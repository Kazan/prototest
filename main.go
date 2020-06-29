package main

import (
	"encoding/json"
	"fmt"

	"github.com/Kazan/prototest_gen_go/computers"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Client starting...")

	mm := new(computers.Memory)

	data, err := proto.Marshal(mm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%v\n%+#v\n", data, data, data)

	json, err := json.MarshalIndent(mm, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON", string(json))
}
