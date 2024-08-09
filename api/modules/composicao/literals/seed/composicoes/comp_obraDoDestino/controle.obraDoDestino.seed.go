package comp_obraDoDestino

import (
	"github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"
)

var ControleObraDoDestino = contractcomposicao.EstadoAtualTrabalho{
	ContratoComposicao: contractcomposicao.ContratoComposicao{
		Vigente:         false,
		DataContratacao: "",
		DataExpiracao:   "",
		Duracao:         0,
		Pagamentos: contractcomposicao.PagamentosLiberacoes{
			ValorContratado: 0,
			EmAberto:        0,
			Pago:            0,
			Debito:          0,
		},
	},
}
