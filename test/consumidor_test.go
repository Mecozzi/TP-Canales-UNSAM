package test

import (
	"testing"
	"time"
)

// Función Consumidor (para la prueba)
func consumidorTest(canal chan int, resultados *[]int) {
	for num := range canal {
		*resultados = append(*resultados, num) // Guardamos el número recibido
	}
}

// Prueba unitaria para verificar que el consumidor recibe los datos correctamente
func TestConsumidor(t *testing.T) {
	canal := make(chan int, 3) // Canal con búfer
	resultados := []int{}      // Slice para almacenar valores recibidos

	// Enviar datos al canal (simulando un productor)
	go func() {
		canal <- 1
		canal <- 2
		canal <- 3
		close(canal)
	}()

	// Ejecutamos el consumidor
	go consumidorTest(canal, &resultados)

	// Esperamos un momento para que el consumidor procese los datos
	time.Sleep(1 * time.Second)

	// Verificamos que el consumidor recibió los valores esperados
	esperado := []int{1, 2, 3}
	if len(resultados) != len(esperado) {
		t.Errorf("Número incorrecto de valores recibidos. Esperado: %d, Recibidos: %d", len(esperado), len(resultados))
	}

	for i, val := range esperado {
		if resultados[i] != val {
			t.Errorf("Valor incorrecto en el consumidor. Esperado: %d, Recibido: %d", val, resultados[i])
		}
	}
}
