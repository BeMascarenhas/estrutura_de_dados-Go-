package main

import (
	"fmt"
)
/*1. 
Implemente em Go a solução recursiva para o problema da Torre de Hanói.*/
// origem = A, trabalho = B, destino = C
func hanoi(n int , origem string, trabalho string, destino string){
	if n > 0{
		hanoi(n-1,origem,destino,trabalho)
		fmt.Printf("Move o disco %v da origem %s para o destino %s\n",n,origem,destino)
		hanoi(n-1, trabalho, origem, destino)
	}

}


/* 2.
Implemente em Go um algoritmo que resolva o seguinte problema: dado um array de números inteiros positivos,
 considerado ordenado crescentemente, e um valor alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. 
 Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado. 
Resolva esse problema com um algoritmo cuja complexidade é O(n). */
func find_par(vetor[] int, valor int) [] int {
	i:= 0
	j:= len(vetor) - 1
	for i<j{
		if  (vetor[i]+vetor[j] == valor){
			return []int {vetor[i],vetor[j]}
			
		} else if ((vetor[i]+vetor[j])>valor){
			j--
		}else{
			i++
		}
	}
	return []int {-1,-1}


}

func main() {
	var qnt_discos int

	fmt.Scanln(&qnt_discos)

	hanoi(qnt_discos,"A","B","C")


	var vetor = []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(find_par(vetor,10))

}
