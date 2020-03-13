package main

import (
	"encoding/gob" // golang object, json
	"fmt"
	"net"
)

type Proceso struct {
	Id     uint64
	Tiempo uint64
}

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}
}

func handleCliente(c net.Conn) {
	var p Proceso
	err := gob.NewDecoder(c).Decode(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mensaje:", p)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
