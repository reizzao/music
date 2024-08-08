package seedcomposicao

import (
	"github.com/reizzao/music/api/modules/composicao/contractcomposicao"
	"github.com/reizzao/music/api/modules/voz/seed/seedvoz"
)

var Frase_Amor_de_Milhoes = contractcomposicao.Frases_Auxiliares{
	CategoriaFrases: contractcomposicao.APAIXONADO,
	Frase:           "AMOR DE MILHOES",
	VogalTermino:    seedvoz.Vogal_U,
}
