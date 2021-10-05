package funcsJokenpo

type jogador struct {
	nome        string
	meusPontos  int
	botPontos   int
	qtdVitorias int
	qtdEmpates  int
	qtdDerrotas int
}

func NewJogador(n string) *jogador {
	return &jogador{
		nome:        n,
		meusPontos: 0,
		botPontos: 0,
		qtdVitorias: 0,
		qtdEmpates: 0,
		qtdDerrotas: 0,
	}
}
