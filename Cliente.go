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

func cliente() {
	// Conexión inicial entre cliente servidor
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("¡Conexión!")
	var p Proceso
	err = gob.NewDecoder(c).Decode(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mensaje: ", p)
	}
	c.Close()
}

func handleProceso(c net.Conn, p Proceso) {

}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input)
}
