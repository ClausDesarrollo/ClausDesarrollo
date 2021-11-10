/*
Para que las descargas sean más rápidas, decide utilizar la simultaneidad.
Cada descarga de archivo se ejecutará como un Goroutine separado

Para simular una descarga de archivos, la función download() debe tomar el tamaño
del archivo como argumento
y contar la suma de todos los enteros desde 0 hasta ese número.

El programa dado toma tres tamaños de archivo como entradas
y llama a la función download() como Goroutines para cada archivo.
Completa el programa declarando la función download() y enviando el resultado
de su operación a main() utilizando canales.
Los resultados deben sumarse en main() y output.
Puede usar canales para obtener los datos de las funciones download() y agregarlos en main().

ej:
*/
package main

import "fmt"

//define el download() function
func download(s int, c chan int) {
	var sum int
	for i := 0; i <= s; i++ {
		sum += i
	}
	c <- sum // <- es igual mente conocido como un select, con una sitaxis similar a switch
}

func main() {
	ch1 := make(chan int) //make corresponde a canales permite recivir y comunicar con las rutinas go
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1 int = 10
	var s2 int = 50
	var s3 int = 42
	//fmt.Scanln(&s1)
	//fmt.Scanln(&s2)
	//fmt.Scanln(&s3)

	go download(s1, ch1) //gorutines
	go download(s2, ch2) // las rutinas go no necesaria mente esperan
	go download(s3, ch3) //  a que termine una rutina antes de iniciar otra
	//para lograr la concurencia iniciamos la gorutina
	//usando la intrucion go

	//output the sum of all results
	fmt.Println(<-ch2 + <-ch1 + <-ch3)
}
