/*
	Trabajo Práctico - Canales en Go (UNSAM)
	Este programa demuestra el uso de goroutines y canales en Go para ejecutar tareas concurrentes.
	Cada goroutine envía un mensaje a través de un canal, y el programa principal los recibe e imprime.
*/

package main

import (
	"fmt"  // Paquete para imprimir mensajes en consola
	"time" // Paquete para manejar tiempos y retardos
)

// enviarMensaje simula una operación concurrente enviando un mensaje a través de un canal
func enviarMensaje(canal chan string, mensaje string) {
	time.Sleep(2 * time.Second) // Simula una tarea que tarda 2 segundos en completarse
	canal <- mensaje            // Envía el mensaje al canal
}

func main() {
	// Creación de un canal sin búfer para comunicar goroutines
	canal := make(chan string)

	// Se inician tres goroutines que enviarán mensajes al canal
	go enviarMensaje(canal, "Goroutine prueba 1")
	go enviarMensaje(canal, "Goroutine 2")
	go enviarMensaje(canal, "Goroutine ensayo 3")

	// Se reciben e imprimen los mensajes del canal (bloquea hasta que reciba los 3)
	for i := 0; i < 3; i++ {
		fmt.Println(<-canal)
	}
}
