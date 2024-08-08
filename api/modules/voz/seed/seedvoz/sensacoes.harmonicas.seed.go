package seedvoz

import "github.com/reizzao/music/api/modules/voz/contractvoz"

var SensacaoRepouso = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Repouso", "Estabilidade"},
}

var SensacaoInicioForte = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Inicio forte", "Estabilidade"},
}

var SensacaoTensao = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Tensao", "Perguntando", "Indefinido Forte", "Precisa Resolver", "Subiu", "Pede para: Fazer outra Coisa"},
}

var SensacaoRepetina = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Freio", "Tensao + Forte"},
}

var SensacaoContinuaForte_Distanciamento = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Distancia", "Viajou", "Faz outra Coisa"},
}

var SensacaoCaminhar = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Caminhar"},
}

var SensacaoReinicioForte = contractvoz.SensacaoHarmonica{
	Sensacoes: []string{"Acaba Sofrendo"},
}
