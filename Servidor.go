package main

import (
	"container/list"
	"encoding/gob" // golang object, json
	"fmt"
	"net"
	"time"
)

type Proceso struct {
	Id     uint64
	Tiempo uint64
}

// Función dedicada al incremento del tiempo de los procesos dentro de una lista
func incrementoProceso(lista list.List) {
	for {
		// System "Clear"
		fmt.Print("\033[H\033[2J")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("|                  Procesos                  |")
		fmt.Println("+--------------------------------------------+")
		// Ciclo que incrementa el tiempo de cada uno de los Procesos.
		for e := lista.Front(); e != nil; e = e.Next() {
			fmt.Println("| id:", e.Value.(Proceso).Id, "                  tiempo:", e.Value.(Proceso).Tiempo, "         |")
			// Al nodo actual se le asigna un valor actualizado del proceso.
			e.Value = Proceso{Id: e.Value.(Proceso).Id, Tiempo: e.Value.(Proceso).Tiempo + 1}
		}
		fmt.Println("+--------------------------------------------+")
		time.Sleep(time.Second / 2)
	}
}

// Función de servidor que estará escuchando para cuando se conecte un Cliente en el puerto :9999
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

// Función que manejará la información que sea proporcionada por el cliente
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
	// Inicialización de la lista
	var listaProcesos list.List
	listaProcesos.PushBack(Proceso{Id: 0, Tiempo: 0})
	listaProcesos.PushBack(Proceso{Id: 1, Tiempo: 0})
	listaProcesos.PushBack(Proceso{Id: 2, Tiempo: 0})
	listaProcesos.PushBack(Proceso{Id: 3, Tiempo: 0})
	listaProcesos.PushBack(Proceso{Id: 4, Tiempo: 0})
	// Hilo de los procesos
	go incrementoProceso(listaProcesos)
	// Hilo del servidor
	go server()
	// Condicionante de paro
	var input string
	fmt.Scanln(&input)
}
