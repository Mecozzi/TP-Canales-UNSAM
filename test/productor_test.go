package test

import (
	"testing"
	"time"
)

// Función Productor 
func productorTest(canal chan int) {
	for i := 1; i <= 3; i++ {
		canal <- i
	}
	close(canal) // Cerramos el canal al terminar
}

// Prueba unitaria para verificar que el productor envía los datos correctamente
func TestProductor(t *testing.T) {
	canal := make(chan int, 3) // Canal con búfer de 3
	go productorTest(canal)    // Ejecutamos el productor en una goroutine

	// Esperamos un tiempo breve antes de verificar los valores
	time.Sleep(1 * time.Second)

	// Verificamos que el canal recibe los valores esperados
	esperado := []int{1, 2, 3}
	i := 0
	for num := range canal {
		if num != esperado[i] {
			t.Errorf("Valor incorrecto en el canal. Esperado: %d, Recibido: %d", esperado[i], num)
		}
		i++
	}

	// Verificamos que recibimos exactamente 3 valores
	if i != 3 {
		t.Errorf("Número incorrecto de valores recibidos. Esperado: 3, Recibidos: %d", i)
	}
}
