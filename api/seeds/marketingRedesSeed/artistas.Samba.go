package marketingRedesSeed

import "github.com/reizzao/music/api/entitys/entityA"

var Artista_Mumuzinho = entityA.Artistas{
	Props: entityA.MarketingRedesProps{
		Nome:      "",
		Genero:    "",
		Youtube:   "",
		Instagram: "",
	},
	Produtor: Produtor_Mumuzinho_Seed,
}

var Artista_Ferrugem = entityA.Artistas{
	Props: entityA.MarketingRedesProps{
		Nome:      "",
		Genero:    "",
		Youtube:   "",
		Instagram: "",
	},
	Produtor: Produtor_Ferrugem_Seed,
}

/*
var Artista_TAL = entityA.Artistas{
	Props: entityA.MarketingRedesProps{
		Nome: "",
		Youtube: "",
		Instagram: "",
	},
}
*/
