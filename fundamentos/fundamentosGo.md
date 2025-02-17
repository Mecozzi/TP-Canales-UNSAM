** Fundamentos de Go (Golang) **  

Este documento cubre los conceptos esenciales del lenguaje **Go**, desde su sintaxis básica hasta la concurrencia con goroutines y canales.  

 **Objetivo:** Explicar de manera clara y práctica los fundamentos de Go, con ejemplos funcionales y bien comentados.  

---

##  1. Introducción a Go**  

Go es un lenguaje de programación moderno, rápido y eficiente, diseñado por Google en 2009. Se caracteriza por:  

✔ **Sintaxis sencilla y fácil de leer.**  
✔ **Compilación rápida y código optimizado.**  
✔ **Manejo eficiente de la concurrencia con goroutines.**  
✔ **Seguridad en memoria sin necesidad de recolección de basura manual.**  

---

##  2. Variables y Tipos de Datos **  


### 🔹 Declaración Explícita**  

```go
var mensaje string = "Hola, Go!"
var edad int = 25

temperatura := 30.5  // Go infiere que es float64
activa := true        // Go infiere que es bool

🔹 Arrays y Slices
var numeros [3]int = [3]int{1, 2, 3}

Slice (dinámico, más usado en Go):
numeros := []int{1, 2, 3}
numeros = append(numeros, 4)  // Agregar elementos dinámicamente

🔹 Maps (Diccionarios en Go)
usuarios := map[string]int{
	"Juan":  30,
	"Maria": 25,
}

fmt.Println(usuarios["Juan"])  // 30

3. Funciones en Go 

func sumar(a int, b int) int {
	return a + b
}

Funciones con múltiples retornos:

func dividir(dividendo, divisor int) (int, int) {
	return dividendo / divisor, dividendo % divisor
}


## 4.  Estructuras (struct) y Métodos

Las estructuras (struct) permiten definir objetos personalizados.

package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

// Método para la estructura
func (p Persona) Saludar() {
	fmt.Println("Hola, soy", p.Nombre, "y tengo", p.Edad, "años.")
}

func main() {
	p := Persona{"Juan", 30}
	p.Saludar()
}

Puntos Clave:
✔ Go no tiene clases ni herencia, en su lugar usa structs.
✔ Podemos definir métodos asociados a una estructura.

Si queremos modificar la estructura desde un método, usamos punteros (*Persona).


package main

import "fmt"

// Definimos una estructura
type Cuenta struct {
	Saldo float64
}

// Método con puntero para modificar el saldo
func (c *Cuenta) Depositar(monto float64) {
	c.Saldo += monto  // Modifica el saldo original
}

func main() {
	c := Cuenta{100}
	c.Depositar(50)  // Se modifica el saldo real
	fmt.Println("Saldo actual:", c.Saldo)  // Salida: 150
}

## 5. Interfaces en Go (Polimorfismo sin Herencia)
En lenguajes como Java o Python, usamos herencia para compartir métodos entre clases.
En Go, no hay herencia, pero sí interfaces para lograr polimorfismo.

package main

import "fmt"

// Definimos una interfaz
type Trabajador interface {
	Trabajar() string
}

// Estructura Persona que implementa la interfaz
type Persona struct {
	Nombre string
}

func (p Persona) Trabajar() string {
	return p.Nombre + " está trabajando"
}

// Estructura Robot que también implementa la interfaz
type Robot struct {
	ID int
}

func (r Robot) Trabajar() string {
	return fmt.Sprintf("Robot %d en ejecución", r.ID)
}

func main() {
	// Polimorfismo: Trabajador puede ser Persona o Robot
	var t Trabajador

	t = Persona{"Juan"}
	fmt.Println(t.Trabajar())  // "Juan está trabajando"

	t = Robot{123}
	fmt.Println(t.Trabajar())  // "Robot 123 en ejecución"
}


## 5. Concurrencia en Go  (Goroutines y Canales)
🔹 Goroutines (Ejecución Concurrente)

package main

import (
	"fmt"
	"time"
)

func imprimirMensaje(mensaje string) {
	for i := 0; i < 3; i++ {
		fmt.Println(mensaje)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go imprimirMensaje("Hola desde goroutine") // Se ejecuta en paralelo
	imprimirMensaje("Hola desde main")         // Se ejecuta en el hilo principal
}

🔹 Canales en Go (Comunicación entre Goroutines)

package main

import "fmt"

func main() {
	canal := make(chan string)

	go func() {
		canal <- "Mensaje desde goroutine"
	}()

	fmt.Println(<-canal) // Recibe el mensaje del canal
}

## 6. Manejo de Errores en Go ⚠️
Go no usa excepciones, sino valores de error explícitos.

package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}


