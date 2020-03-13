package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Proceso struct {
	Id     uint64
	Tiempo uint64
}

func cliente(p Proceso) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enviando", p)
	err = gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	p := Proceso{
		Id:     2,
		Tiempo: 5,
	}
	go cliente(p)

	var input string
	fmt.Scanln(&input)
}
