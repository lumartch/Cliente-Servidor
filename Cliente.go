package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
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
	// Se captura el proceso enviado
	var p Proceso
	err = gob.NewDecoder(c).Decode(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mensaje: ", p)
	}
	go incremento(p)
}

func incremento(p Proceso) {
	for {
		// System "Clear"
		fmt.Print("\033[H\033[2J")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("|                  Procesos                  |")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("| id:", p.Id, "                  tiempo:", p.Tiempo, "         |")
		p = Proceso{Id: p.Id, Tiempo: p.Tiempo + 1}
		fmt.Println("+--------------------------------------------+")
		time.Sleep(time.Second / 2)
	}
}

func handleProceso(c net.Conn, p Proceso) {

}

func main() {
	go cliente()
	var input string
	fmt.Scanln(&input)
	/*fmt.Println("Enviando... ", pr)
	err := gob.NewEncoder(c).Encode(&pr)
	if err != nil {
		fmt.Println(err)
		return
	}*/
}
