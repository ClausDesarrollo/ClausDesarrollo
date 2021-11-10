package main

import "fmt"

func main() {
	var age int = 23
	//fmt.Scanln(&age)

	var ageDaysE int
	var ageDaysM int

	ageDaysE = age * 365      //ej 2 años = 730  //aritmeticos
	ageDaysM = ageDaysE / 687 // 730 /687 = 0. ...

	//mars := mars_age(age)
	//fmt.Println(mars)
	fmt.Println(ageDaysM)
}
