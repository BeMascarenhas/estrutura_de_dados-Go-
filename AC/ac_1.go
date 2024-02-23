package main

import (
	"fmt"
	
)

// 01.

func calculaMedia(num ...float64) float64 {
	soma := 0.0

	for _, n := range num {
		soma += n
	}
	return soma / float64(len(num))
}
//02.
func verificarParidade(num int) string{
	if num%2==0{
		return "Par"
	} else{
 
		return "Ímpar"
	}
}
//03.
func minhaPotencia(base,expoente int) int {
	ans := 1
	for i:=0;i<expoente;i++{
		ans*=base


	}
	return ans

}
//04.
func converteCelsiusParaFahrenheit(temp float64) float64{
	return (temp * 9/5) + 32
}

func main() {
	media := calculaMedia(10,20,30)
	fmt.Println("A média é:",media)


	var verificar_num,base,expoente int
	fmt.Print("Digite um numero para saber se é par ou ímpar: ")
	fmt.Scanln(&verificar_num)
	fmt.Println("O numero é:",verificarParidade(verificar_num))
	fmt.Print("Digite a base e o expoente que calcularemos a potencia: ")
	fmt.Scanln(&base, &expoente)
	fmt.Println("O numero é:",minhaPotencia(base,expoente))
	var temperatura float64
	fmt.Print("Digite uma temperatura em celcius que diremos quanto ela vale em farenheit: ")
	fmt.Scanln(&temperatura)
	fmt.Printf("A temperatura correta é: %f", temperatura)


}
