package seedcomposicao

import (
	"github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"
	"github.com/reizzao/music/api/modules/marketing/contract/contractmarketing"
)

var CompositorREIZAO = contractcomposicao.Compositor{
	NomeArtistico: "Reizao",
	Redes: contractmarketing.RedesCompositor{
		Instagram: "@Compositor_REIZAO",
		Youtube:   "www.youtube.com/@ReizaoMusic",
		PaginaWeb: "www.sites.google.com/view/reizaomusic",
	},
}
