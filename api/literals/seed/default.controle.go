package seed

import (
	"github.com/reizzao/music/api/modules/composicao/types_aux"
)

var ControleDEFAULT = types_aux.EstadoAtualTrabalho{
	ContratoComposicao: types_aux.ContratoComposicao{
		Vigente:         false,
		DataContratacao: "",
		DataExpiracao:   "",
		Duracao:         0,
		Pagamentos: types_aux.PagamentosLiberacoes{
			ValorContratado: 0,
			EmAberto:        0,
			Pago:            0,
			Debito:          0,
		},
	},
}
