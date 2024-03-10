package main

//Calculo de Passos
import (
	"fmt"
)

func buscaMatriz(m [][]int, n int, x int) bool {
	var i int = 0

	for i < n {
		j := 0
		for j < n {
			if m[i][j] == x {
				return true
			}
			j++
		}
		i++
	}
	return false
}
func par(a []int, x int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == x {
				return []int{a[i], a[j]}
			}
		}
	}
	return []int{-1, -1}
}
func par2(a []int, x int) []int {
	indices := make(map[int]int)

	for i, num := range a {
		complement := x - num
		if j, found := indices[complement]; found {
			return []int{j, i}
		}
		indices[num] = i
	}

	return []int{-1, -1}
}

func main() {
	matriz := [][]int{{1, 2}, {3, 4}}
	b := buscaMatriz(matriz, 2, 2)
	fmt.Println(b)
	c := []int{1, 2, 3, 4}
	fmt.Println(par(c, 7))
}


