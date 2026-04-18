// hola
package main

import "fmt"

func main() {
	var v1, v2 int
	var op string
	var c int = 0

	for {

		fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
		fmt.Println("Ingresa el primer número")
		fmt.Scanln(&v1)
		fmt.Println("Ingresa el segundo número")
		fmt.Scanln(&v2)

		fmt.Println("Ingresa la operación (+, -, *, /,^,!):")
		fmt.Scanln(&op)

		switch op {
		case "+":

			sum := v1 + v2
			fmt.Printf("Resultado: %d + %d = %d\n", v1, v2, sum)

		case "-":

			rest := v1 - v2
			fmt.Printf("Resultado: %d - %d = %d\n", v1, v2, rest)

		case "*":

			mult := v1 * v2
			fmt.Printf("Resultado: %d * %d = %d\n", v1, v2, mult)

		case "/":
			if v2 == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
			} else {
				div := float64(v1) / float64(v2)
				fmt.Printf("Resultado: %.d / %d = %.2f\n", v1, v2, div)
			}

		case "^":
			au := 1

			for i := 1; i <= v2; i++ {

				au = au * v1

			}
			fmt.Printf("Resultado: %.d^%d = %.d\n", v1, v2, au)

		case "!":
			au2 := 1
			for i := 1; i <= v1; i++ {

				au2 = au2 * i

			}
			fmt.Printf("Resultado: %.d!= %.d\n", v1, au2)

		default:
			fmt.Println("Operación no válida")

		}
		c = c + 1
		var r string
		fmt.Println("¿Deseas realizar otra operación? (s/n)")
		fmt.Scanln(&r)
		if r == "n" {
			fmt.Printf("Total de operaciones realizadas: %d\n", c)
			fmt.Println("¡Hasta luego!")
			break
		}
	}

}
