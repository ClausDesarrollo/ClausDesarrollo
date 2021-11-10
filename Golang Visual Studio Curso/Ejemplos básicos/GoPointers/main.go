//Está creando una aplicación timer que debe contar hasta un número determinado.
//Su programa debe tomar un número como entrada y hacer que el temporizador marque ese número de veces.

//El código en main inicializa un temporizador y toma un número como entrada. Luego
//llama al método tick() para el temporizador el número dado de veces.

//Defina la estructura del temporizador con dos campos: id y value, y defina el método tick(),
//que debe incrementar el valor en uno y generar su valor actual.
//Utilice un puntero de estructura como receptor del método para poder cambiar el valor del temporizador.

package main

import "fmt"

type Timer struct {
	id    string
	value int
}

/* define the tick() method, which
should increment the value by one
and output its current value. */
func (t *Timer) tick() { // Apuntadores son variables especiales
	//que contienen la dirección de memoria de los valores var p *int
	for i := 0; i < t.value; i++ { // > < = relacionales
		fmt.Println(i + 1)
	}
}

func main() {
	var x int = 32
	//fmt.Scanln(&x)
	t := Timer{"timer1", x}
	t.tick()
}
