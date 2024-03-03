package geometria

/*
4
Crie um pacote chamado geometria que contenha funções para calcular a área e o perímetro de um retângulo. 
Em seguida, use esse pacote em um programa principal para calcular a área e o perímetro 
de um retângulo com dimensões fornecidas pelo usuário.
*/
func Area(base, altura float64) float64 {
    return base * altura
}

func Perimetro(base, altura float64) float64 {
    return 2 * (base + altura)
}