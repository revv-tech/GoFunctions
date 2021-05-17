package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	genPseudoAleatorios()

}

//Auxiliar que verifica si es primo
func esPrimo(n int) bool {

	if n <= 1 {
		fmt.Printf("No es primo")
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 && n != i {
			fmt.Printf("No es primo")
			return false
		}
	}
	fmt.Printf("Es primo")
	return true
}

//Generador de arreglo de n numeros pseudo-aleatorios
func genPseudoAleatorios() []int {
	//Creador del scanner para recibir input para recibir n
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Ingresa la cantidad de numeros pseudo-aleatorios que desea generar. Debe estar en el intervalo [200,999] : ")
	scanner.Scan()
	//n := cantidad de numeros a generar
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Print("La cantidad es: ", n, "\n")

	//Creador del scanner para recibir input para recibir semilla
	scanner_2 := bufio.NewScanner(os.Stdin)
	fmt.Printf("Ingresa la semilla con la que deseas empezar debe ser primo y estar en el intervalo [11,101]: ")
	scanner_2.Scan()
	//semilla
	semilla, _ := strconv.Atoi(scanner_2.Text())
	fmt.Print("La cantidad es: ", semilla, "\n")

	//Verifica si la semilla es primo y se encuentran en el rango y si la semilla es prima
	if ((200 < n) && (n < 1000)) && ((11 <= semilla && semilla <= 101) && esPrimo(int(semilla))) {

		return genPseudoAleatorios_Aux(n, semilla)

	}

	fmt.Println("No se cumplen los requisitos para realizar el metodo Pseudo-random Sequences!")
	//Si no cumple las condiciones, retorna un array vascio
	return []int{}

}

//Auxiliar que realiza operaciones, y crear el array
func genPseudoAleatorios_Aux(n int, x int) []int {

	var arrayInts []int
	m := 4624 //Modulo (Periodo)
	a := 73   //Multiplicador
	b := 956  //Incremento

	tmp := x

	for i := 1; i <= n; i++ {

		numAleatorio := (a*tmp + b) % m
		fmt.Println(numAleatorio)
		arrayInts = append(arrayInts, numAleatorio)
		tmp = numAleatorio

	}

	for _, num := range arrayInts {
		fmt.Println(num)
	}
	return arrayInts
}