package seedcomposicao

import (
	"github.com/reizzao/music/api/modules/composicao/contractcomposicao"
	"github.com/reizzao/music/api/modules/marketing/contractmarketing"
)

var CompositorPaulinhoDC = contractcomposicao.Compositor{
	NomeArtistico: "PaulinhoDC",
	Redes: contractmarketing.RedesCompositor{
		Instagram: "@PaulinhoDC15",
		Youtube:   "",
		PaginaWeb: "",
	},
}
