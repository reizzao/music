package seed

import "github.com/reizzao/music/api/modules/composicao/contract"

var ControleDEFAULT = contract.EstadoAtualTrabalho{
	ContratoComposicao: contract.ContratoComposicao{
		Vigente:         false,
		DataContratacao: "",
		DataExpiracao:   "",
		Duracao:         0,
		Pagamentos: contract.PagamentosLiberacoes{
			ValorContratado: 0,
			EmAberto:        0,
			Pago:            0,
			Debito:          0,
		},
	},
}
