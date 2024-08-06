package seed

import "github.com/reizzao/music/api/modules/composicao/contract"

var CompositorREIZAO = contract.Compositor{
	NomeArtistico: "Reizao",
	Redes: contract.RedesCompositor{
		Instagram: "@Compositor_REIZAO",
		Youtube:   "www.youtube.com/@ReizaoMusic",
		PaginaWeb: "www.sites.google.com/view/reizaomusic",
	},
}
