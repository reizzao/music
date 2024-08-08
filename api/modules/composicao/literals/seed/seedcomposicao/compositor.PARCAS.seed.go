package seedcomposicao

import (
	"github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"
	"github.com/reizzao/music/api/modules/marketing/contract/contractmarketing"
)

var CompositorPaulinhoDC = contractcomposicao.Compositor{
	NomeArtistico: "PaulinhoDC",
	Redes: contractmarketing.RedesCompositor{
		Instagram: "@PaulinhoDC15",
		Youtube:   "",
		PaginaWeb: "",
	},
}
