package seedvoz

import "github.com/reizzao/music/api/modules/voz/contractvoz"

var SensacaoRepouso = contractvoz.Sensacao{
	Grau: "1",
	Sensacoes: []string{"Repouso", "Estabilidade"},
}

var SensacaoTensao = contractvoz.Sensacao{
	Grau: "5",
	Sensacoes: []string{"Tensao", "Precisa Resolver", "Subiu", "Pede para: Fazer outra Coisa"},
}

var SensacaoDistanciamento = contractvoz.Sensacao{
	Grau: "4",
	Sensacoes: []string{"Distancia", "Viajou", "Faz outra Coisa"},
}

var SensacaoCaminhar = contractvoz.Sensacao{
	Grau: "2",
	Sensacoes: []string{"Caminhar",},
}

var SensacaoTerminaFraseSofrer = contractvoz.Sensacao{
	Grau: "3",
	Sensacoes: []string{"Acaba Sofrendo",},
}