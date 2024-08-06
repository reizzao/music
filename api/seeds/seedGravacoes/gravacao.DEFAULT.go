package seedGravacoes

import (
	"github.com/reizzao/music/api/entitys/entityB"
	"github.com/reizzao/music/api/seeds/controleSeed"
)

var GravacaoDEFAULT = entityB.GravacoesComposicoes{
	Interprete:   "",
	TipoGravacao: "",
	Data:         "",
	Controle:     controleSeed.ControleDEFAULT,
}
