package funcsJokenpo

import "fmt"

func (j jogador) placar() {

	if j.meusPontos <= 9 && j.botPontos <=9{
		fmt.Println("\n\tPlacar:")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|   Você   |  JokeBot |")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|                     |")
		fmt.Printf("\t\t|    %d    VS    %d     |\n", j.meusPontos, j.botPontos)
		fmt.Println("\t\t|                     |")
		fmt.Println("\t\t-----------------------")
	} else if j.meusPontos >= 10 && j.botPontos >= 10{
		fmt.Println("\n\tPlacar:")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|   Você   |  JokeBot |")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|                     |")
		fmt.Printf("\t\t|   %d    VS    %d    |\n", j.meusPontos, j.botPontos)
		fmt.Println("\t\t|                     |")
		fmt.Println("\t\t-----------------------")
	} else {
		fmt.Println("\n\tPlacar:")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|   Você   |  JokeBot |")
		fmt.Println("\t\t-----------------------")
		fmt.Println("\t\t|                     |")
		fmt.Printf("\t\t|    %d    VS    %d    |\n", j.meusPontos, j.botPontos)
		fmt.Println("\t\t|                     |")
		fmt.Println("\t\t-----------------------")
	}

}

