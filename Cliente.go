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

func cliente(p *Proceso) {
	// Conexión inicial entre cliente servidor
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Se captura el proceso enviado por el servidor
	fmt.Println("¡Conexión!")
	err = gob.NewDecoder(c).Decode(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mensaje: ", p)
		incremento(p)
	}
}

func incremento(p *Proceso) {
	for {
		// System "Clear"
		fmt.Print("\033[H\033[2J")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("|                  Procesos                  |")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("| id:", p.Id, "                  tiempo:", p.Tiempo, "         |")
		p.Tiempo += 1
		fmt.Println("+--------------------------------------------+")
		time.Sleep(time.Second / 2)
	}
}

func enviarProceso(p *Proceso) {
	// Conexión inicial entre cliente servidor
	c, err := net.Dial("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var p Proceso
	go cliente(&p)
	var input string
	fmt.Scanln(&input)
	// Conexión inicial entre cliente servidor
	enviarProceso(&p)
}
