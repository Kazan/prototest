package main

import (
	"encoding/json"
	"fmt"
)

type pepe struct {
	Master bool
}

func (p *pepe) speak() {
	fmt.Println("Im pepe")
}

func main() {
	fmt.Println("Client starting...")

	// p := new(pepe)

	// data, err := proto.Marshal(p)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%s\n%v\n%+#v\n", data, data, data)

	ell := new(pepe)
	// err = proto.Unmarshal(data, ell)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+#v\n", ell)

	json, err := json.MarshalIndent(ell, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}
