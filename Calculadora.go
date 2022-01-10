package main

import (
	"fmt"
)

var option int

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
		fmt.Scanln(&option)

		switch option {

		case 0:
			fmt.Println("Encerrando programa...")
			return

		case 1:
			var base, altura float32
			fmt.Print("Digite o tamanho da base do triângulo: ")
			fmt.Scanln(&base)
			fmt.Print("Digite o tamanho da altura do triângulo: ")
			fmt.Scanln(&altura)
			area := trianguloArea(base, altura)
			perimetro := trianguloPerimetro(base)
			fmt.Println("Área do triângulo: ", area)
			fmt.Println("Perímetro do triângulo: ", perimetro)
			fmt.Print("\n")

		case 2:
			var base, altura float32
			fmt.Print("Digite o tamanho da base do retângulo: ")
			fmt.Scanln(&base)
			fmt.Print("Digite o tamanho da altura do retângulo: ")
			fmt.Scanln(&altura)
			area := retanguloArea(base, altura)
			perimetro := retanguloPerimetro(base, altura)
			fmt.Println("Área do retângulo: ", area)
			fmt.Println("Perímetro do retângulo: ", perimetro)
			fmt.Print("\n")

		default:
			fmt.Println("Opção inválida. Selecione uma das opções disponíveis!")
			fmt.Print("\n")
		}

	}

}

func trianguloArea(base float32, altura float32) float32 {
	return (base * altura) / 2
}

func trianguloPerimetro(base float32) float32 {
	return base * 3
}

func retanguloArea(base float32, altura float32) float32 {
	return (base * altura)
}

func retanguloPerimetro(base float32, altura float32) float32 {
	return base*2 + altura*2
}
