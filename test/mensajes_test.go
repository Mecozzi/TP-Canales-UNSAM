package test

import (
	"testing"
	"time"
)

// Prueba para verificar que enviarMensaje envía el mensaje correctamente
func TestEnviarMensaje(t *testing.T) {
	canal := make(chan string, 1) // Canal con espacio para un mensaje

	go func() {
		time.Sleep(1 * time.Second) // Simula una espera antes de enviar
		canal <- "Hola prueba"
	}()

	select {
	case msg := <-canal:
		if msg != "Hola prueba" {
			t.Errorf("Mensaje incorrecto. Esperado: 'Hola prueba', Recibido: '%s'", msg)
		}
	case <-time.After(2 * time.Second):
		t.Error("Timeout: No se recibió el mensaje en el tiempo esperado")
	}
}
