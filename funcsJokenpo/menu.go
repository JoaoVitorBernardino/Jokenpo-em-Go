package funcsJokenpo

import "fmt"

func Menu(j jogador){
	jogando := true
	for jogando{
		var op int
		fmt.Println("Digite a opção que deseja:")
		fmt.Println("1 - Jogar\n2 - Ver seus status\n3 - sair")
		_, err := fmt.Scan(&op)

		if err != nil {
			return
		}

		switch op {
		case 1:
			var quantosRounds int
			fmt.Println("Digite quantos rounds você quer jogar:")
			_, err := fmt.Scan(&quantosRounds)

			if err != nil {
				return
			}

			j.qtdVitorias, j.qtdEmpates, j.qtdDerrotas = j.Jogar(quantosRounds)
		case 2:
			j.VerStatus()
		case 3:
			jogando = false
		default:
			fmt.Println("Digite uma opção válida.")
		}

	}
}

