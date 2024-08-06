package seeddefault

import (
	"github.com/reizzao/music/api/entitys/entityC"
	"github.com/reizzao/music/api/entitys/entityD"
	"github.com/reizzao/music/api/entitys/entityE"
)

var ControleDEFAULT = entityC.EstadoAtualTrabalho{
	ContratoComposicao: entityD.ContratoComposicao{
		Vigente:         false,
		DataContratacao: "",
		DataExpiracao:   "",
		Duracao:         0,
		Pagamentos: entityE.PagamentosLiberacoes{
			ValorContratado: 0,
			EmAberto:        0,
			Pago:            0,
			Debito:          0,
		},
	},
}
