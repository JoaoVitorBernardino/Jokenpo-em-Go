package main

import (
	"fmt"
	"playJokenpo.go/funcsJokenpo"
)

func main() {
	var nome string
	fmt.Println("Para começar jogar, informe seu nome:")
	_, err := fmt.Scan(&nome)

	if err != nil {
		return
	}

	jogador1 := funcsJokenpo.NewJogador(nome)

	funcsJokenpo.Menu(*jogador1)
}
