package main

import (
	"fmt"
)

var option int

const PI = 3.1415

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
			fmt.Print("\n")
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
			fmt.Print("\n")
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
			fmt.Print("\n")
			fmt.Println("Área do retângulo: ", area)
			fmt.Println("Perímetro do retângulo: ", perimetro)
			fmt.Print("\n")

		case 3:
			var lado float32
			fmt.Print("Digite o tamanho do lado do quadrado: ")
			fmt.Scanln(&lado)
			area := quadradoArea(lado)
			perimetro := quadradoPerimetro(lado)
			fmt.Print("\n")
			fmt.Println("Área do quadrado: ", area)
			fmt.Println("Perímetro do quadrado: ", perimetro)
			fmt.Print("\n")

		case 4:
			var raio float32
			fmt.Print("Digite o tamanho do raio do círculo: ")
			fmt.Scanln(&raio)
			area := circuloArea(raio)
			perimetro := circuloPerimetro(raio)
			fmt.Print("\n")
			fmt.Println("Área do círculo: ", area)
			fmt.Println("Perímetro do círculo: ", perimetro)
			fmt.Print("\n")

		case 5:
			var areabase, arealateral, altura float32
			fmt.Print("Digite o tamanho da área base da pirâmide: ")
			fmt.Scanln(&areabase)
			fmt.Print("Digite o tamanho da área lateral da pirâmide: ")
			fmt.Scanln(&arealateral)
			fmt.Print("Digite o tamanho da altura da pirâmide: ")
			fmt.Scanln(&altura)
			area := piramideArea(areabase, arealateral)
			volume := piramideVolume(areabase, altura)
			fmt.Print("\n")
			fmt.Println("Área da pirâmide: ", area)
			fmt.Println("Volume da pirâmide: ", volume)
			fmt.Print("\n")

		case 6:
			var aresta float32
			fmt.Print("Digite o tamanho da aresta do cubo: ")
			fmt.Scanln(&aresta)
			area := cuboArea(aresta)
			volume := cuboVolume(aresta)
			fmt.Print("\n")
			fmt.Println("Área do cubo: ", area)
			fmt.Println("Volume do cubo: ", volume)
			fmt.Print("\n")

		case 7:
			var aresta1, aresta2, aresta3 float32
			fmt.Print("Digite o tamanho das arestas do paralelepípedo: ")
			fmt.Scanln(&aresta1)
			fmt.Scanln(&aresta2)
			fmt.Scanln(&aresta3)
			area := paralelepipedoArea(aresta1, aresta2, aresta3)
			volume := paralelepipedoVolume(aresta1, aresta2, aresta3)
			fmt.Print("\n")
			fmt.Println("Área do paralelepípedo: ", area)
			fmt.Println("Volume do paralelepípedo: ", volume)
			fmt.Print("\n")

		case 8:
			var raio float32
			fmt.Print("Digite o tamanho do raio da esfera: ")
			fmt.Scanln(&raio)
			area := esferaArea(raio)
			volume := esferaVolume(raio)
			fmt.Print("\n")
			fmt.Println("Área do círculo: ", area)
			fmt.Println("Volume do círculo: ", volume)
			fmt.Print("\n")

		default:
			fmt.Print("\n")
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

func quadradoArea(lado float32) float32 {
	return (lado * lado)
}

func quadradoPerimetro(lado float32) float32 {
	return lado * 4
}

func circuloArea(raio float32) float32 {
	return PI * raio * raio
}

func circuloPerimetro(raio float32) float32 {
	return 2 * PI * raio
}

func piramideArea(areabase float32, arealateral float32) float32 {
	return areabase + arealateral
}

func piramideVolume(areabase float32, altura float32) float32 {
	return 1 / 3 * areabase * altura
}

func cuboArea(aresta float32) float32 {
	return 6 * aresta * aresta
}

func cuboVolume(aresta float32) float32 {
	return aresta * aresta * aresta
}

func paralelepipedoArea(aresta1 float32, aresta2 float32, aresta3 float32) float32 {
	return (2 * aresta1 * aresta2) + (2 * aresta1 * aresta3) + (2 * aresta2 * aresta3)
}

func paralelepipedoVolume(aresta1 float32, aresta2 float32, aresta3 float32) float32 {
	return aresta1 * aresta2 * aresta3
}

func esferaArea(raio float32) float32 {
	return 4 * PI * raio * raio
}

func esferaVolume(raio float32) float32 {
	return 4 / 3 * PI * raio * raio * raio
}
