package main

import (
	"fmt"
	"strings"
)

func main(){
x:=2
y:=x

fmt.Println(x,y)
x=3
// ao mudar o x o valor de y se mantem intacto
fmt.Println(x,y)

z:=&x//referencia

fmt.Println(x,z)
fmt.Println(x,*z)//dereferencia

x=4
// por ser um ponteiro ao mudar x o valor de z tambem muda
fmt.Println(x,z)
fmt.Println(x,*z)//dereferencia


msg1 := "Ola, mundo"
alteraMensagem(&msg1)


n1 := Numero{Valor: 10}
fmt.Println(n1)
n1.Incremento()
fmt.Println(n1)


}
type Numero struct{
	Valor int
}

func(n *Numero) Incremento(){
	n.Valor++
}
/*
Usamos ponteiros como parametros de funcoes quando:
1. é necessario reduzir o espaço consumido em memoria
2.queremos alterar o valor dos paramentros


*/
func alteraMensagem(msg*string){
	novaMensagem:= strings.ReplaceAll(*msg,"mundo","turma")
	*msg = novaMensagem

}
