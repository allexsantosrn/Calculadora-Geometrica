package main

import (
	"fmt"
)

var option int = 0

func main() {

	for {
		fmt.Println("Calculadora de geometria Plana e Espacial")
		fmt.Println("(1) Triângulo equilátero")
		fmt.Println("(2) Retângulo")
		fmt.Println("(3) Quadrado")
		fmt.Println("(4) Círculo")
		fmt.Println("(5) Pirâmide com base quadrangular")
		fmt.Println("(6) Cubo")
		fmt.Println("(7) Paralelepípedo")
		fmt.Println("(8) Esfera")
		fmt.Println("(0) Sair")

		fmt.Print("Digite a sua opção: ")
		fmt.Scanf("%d", &option)

		switch option {

		}

		if option == 0 {
			fmt.Println("Encerrando programa...")
			return
		}
	}

}
