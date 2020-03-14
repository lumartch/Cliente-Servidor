package main

import (
	"container/list" // golang object, json
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type Proceso struct {
	Id     uint64
	Tiempo uint64
}

// Función dedicada al incremento del tiempo de los procesos dentro de una lista
func incrementoProceso(lista *list.List) {
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
func server(listaProcesos *list.List) {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil || listaProcesos.Front() == nil {
			fmt.Println(err)
			continue
		}
		p := listaProcesos.Front().Value.(Proceso)
		listaProcesos.Remove(listaProcesos.Front())
		go handleCliente(c, p, listaProcesos)
	}
}

// Función que manejará la información que sea proporcionada por el cliente
func handleCliente(c net.Conn, p Proceso, listaProcesos *list.List) {
	err := gob.NewEncoder(c).Encode(&p)
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	c.Close()
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
	go incrementoProceso(&listaProcesos)
	// Hilo del servidor
	go server(&listaProcesos)
	// Condicionante de paro
	var input string
	fmt.Scanln(&input)
}
