package funcsJokenpo

import (
	"fmt"
	"math/rand"
)

func (j jogador) Jogar(q int) (int, int, int) {

	var jogada int

	for i := 0; i < q; i++ {
		numRandom := rand.Intn(3)
		fmt.Println("Escolha sua jogada:")
		fmt.Println("0 - Pedra\n1 - Papel\n2 - tesoura")
		fmt.Scanln(&jogada)

		if jogada == 0 && numRandom == 0 {
			fmt.Println("Pedra VS Pedra: Empate!")
			j.qtdEmpates += 1
		} else if jogada == 0 && numRandom == 1 {
			fmt.Println("Pedra VS Papel: Vitória do Adversário!")
			j.qtdDerrotas += 1
			j.botPontos += 1
		} else if jogada == 0 && numRandom == 2 {
			fmt.Println("Pedra VS Tesoura: Você Venceu!")
			j.qtdVitorias += 1
			j.meusPontos += 1
		} else if jogada == 1 && numRandom == 0 {
			fmt.Println("Papel VS Pedra: Você Venceu!")
			j.qtdVitorias += 1
			j.meusPontos += 1
		} else if jogada == 1 && numRandom == 1 {
			fmt.Println("Papel VS Papel: Empate!")
			j.qtdEmpates += 1
		} else if jogada == 1 && numRandom == 2 {
			fmt.Println("Papel VS Tesoura: Vitória do Adversário!")
			j.qtdDerrotas += 1
			j.botPontos += 1
		} else if jogada == 2 && numRandom == 0 {
			fmt.Println("Tesoura VS Pedra: Vitória do Adversário!")
			j.qtdDerrotas += 1
			j.botPontos += 1
		} else if jogada == 2 && numRandom == 1 {
			fmt.Println("Tesoura VS Papel: Você Venceu!")
			j.qtdVitorias += 1
			j.meusPontos += 1
		} else if jogada == 2 && numRandom == 2 {
			fmt.Println("Tesoura VS Tesoura: Empate!")
			j.qtdEmpates += 1
		}
	}
	j.placar()

	return j.qtdVitorias, j.qtdEmpates, j.qtdDerrotas
}
