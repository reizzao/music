package seedadm

import (
	"github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"
)

var GravacaoDEFAULT = contractcomposicao.GravacoesComposicoes{
	Interprete:   "",
	TipoGravacao: "",
	Data:         "",
	Controle:     ControleDEFAULT,
}
