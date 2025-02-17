/*
	 Productor-Consumidor en Go con Canales
	- Un productor genera números y los envía por un canal.
	- Un consumidor recibe los números y los procesa.
*/

package ejemplos

import (
	"fmt"
	"time"
)

// productor genera números y los envía al canal
func productor(canal chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Produciendo: %d\n", i)
		canal <- i // Enviar número al canal
		time.Sleep(time.Second) // Simula tiempo de procesamiento
	}
	close(canal) // Cierra  cuando termina
}

// consumidor recibe y procesa los números del canal
func consumidor(canal chan int) {
	for num := range canal {
		fmt.Printf(" Consumido: %d\n", num)
		time.Sleep(500 * time.Millisecond) // Simula tiempo de consumo
	}
}

func main() {
	canal := make(chan int) 

	// Iniciamos las goroutines
	go productor(canal)
	go consumidor(canal)

	// Esperamos para que las goroutines terminen 
	time.Sleep(6 * time.Second)
}
