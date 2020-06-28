package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Pepe struct {
	Master bool
}

func (p *Pepe) speak() {
	fmt.Println("Im pepe")
}

func main() {
	fmt.Println("Client starting...")

	p := &prototest.Person{
		Name: "pepe",
		Age:  14,
		Followers: &prototest.Followers{
			Facebook: 12,
			Youtube:  2,
		},
		PersonOfInterest: "yes he is",
	}

	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%v\n%+#v\n", data, data, data)

	ell := &prototest.Person{}
	err = proto.Unmarshal(data, ell)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+#v\n", ell)

	json, err := json.MarshalIndent(ell, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
