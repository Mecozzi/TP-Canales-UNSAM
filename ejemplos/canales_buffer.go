/*
	Ejemplo de Canales con Búfer en Go
	- Se crea un canal con capacidad para 3 mensajes.
	- Se envían 3 mensajes sin bloquear la ejecución.
	- Luego, el consumidor recibe los mensajes.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Canal con un búfer de 3 posiciones
	canal := make(chan string, 3)

	// Enviar mensajes al canal (no se bloquea hasta que el búfer se llene)
	canal <- "Mensaje 1"
	canal <- "Mensaje 2"
	canal <- "Mensaje 3"

	fmt.Println("✅ Mensajes enviados al canal con búfer")

	// Recibir los mensajes del canal
	fmt.Println(" Recibiendo:", <-canal)
	fmt.Println(" Recibiendo:", <-canal)
	fmt.Println(" Recibiendo:", <-canal)

	// Intentar enviar otro mensaje bloquea hasta que haya espacio en el búfer
	go func() {
		canal <- "Mensaje 4"
		fmt.Println(" Mensaje 4 enviado después de liberar espacio en el canal")
	}()

	time.Sleep(1 * time.Second) // Simula tiempo de espera antes de recibir otro mensaje
	fmt.Println(" Recibiendo:", <-canal)

	close(canal) // Cierra el canal cuando ya no se usará
}
