package funcsJokenpo

import "fmt"

func (j jogador) VerStatus(){
	fmt.Println("Nome:", j.nome)
	fmt.Println("Quantidade de vitórias:", j.qtdVitorias)
	fmt.Println("Quantidade de empates:", j.qtdEmpates)
	fmt.Println("Quantidade de derrotas:", j.qtdDerrotas)
}

