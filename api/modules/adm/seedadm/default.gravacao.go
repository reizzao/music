package seedadm

import (
	"github.com/reizzao/music/api/modules/composicao/contractcomposicao"
)

var GravacaoDEFAULT = contractcomposicao.GravacoesComposicoes{
	Interprete:   "",
	TipoGravacao: "",
	Data:         "",
	Controle:     ControleDEFAULT,
}
