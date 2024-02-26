package main

import (
	"fmt"
	"projeto/modelos"
	"projeto/modelos/academico"
)

func main(){
	aluno := modelos.Aluno{}
	aluno.Nome = "bernardo"
	aluno.Matricula = "1234"

	curso := academico.Curso{Nome: "Engenharia"}
	fmt.Println(aluno,curso)
}