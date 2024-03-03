package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"ac2/geometria"
)

/*
1.
Faça um programa que cria um vetor de inteiros com 10 elementos.
Então inicialize este vetor com números quaisquer e imprima na tela todos os seus elementos, um abaixo do outro.
*/
func cria_vector(){
	vetor :=[10]int{1,2,3,4,5,6,7,8,9,10}

	for i:=0; i<10;i++{
		fmt.Println(vetor[i])
	}
}

/*
2.
Faça uma função/método que receba uma string como parâmetro e que retorne uma nova string, 
onde a sequência dos caracteres foi invertida. 
Dentro da parte principal (main), leia uma string digitada pelo usuário e passe para a função/método criada, 
imprimindo em seguida a string devolvida. Para esse exercício, 
pesquise sobre o comportamento e a interação entre dados do tipo rune e do tipo string.
*/
func palavra_invertida(palavra string) string{
	p := []rune(palavra)
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return string(p)
}
/*
3
Crie um tipo chamado Ponto que represente um ponto no plano cartesiano, com coordenadas X e Y. 
Em seguida, implemente um método chamado DistanciaOrigem() que calcule a distância desse ponto até a origem (0,0).
*/
type Ponto struct{
	x float64
	y float64
}
func (p *Ponto) DistanciaOrigem() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}



func main(){
	//1
	cria_vector()
	//2
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Informe uma palavra: ")
	palavra,_:=leitor.ReadString('\n')
	fmt.Println(palavra_invertida(palavra))
//3
	ponto := Ponto{x:4.5,y:3}

	fmt.Printf("A distância do ponto (%.2f, %.2f) até a origem é %.2f\n", ponto.x, ponto.y,ponto.DistanciaOrigem())
//4
	var base, altura float64
    fmt.Print("digite a base do retângulo: ")
    fmt.Scanln(&base)
    fmt.Print("digite a altura do retângulo: ")
    fmt.Scanln(&altura)
    area := geometria.Area(base, altura)
    fmt.Printf("area do retângulo: %.2f\n", area)
    perimetro := geometria.Perimetro(base, altura)
    fmt.Printf("perimetro do retângulo: %.2f\n", perimetro)
}

