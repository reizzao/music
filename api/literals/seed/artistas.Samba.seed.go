package seed

import "github.com/reizzao/music/api/entitys/entityA"

var Artista_Mumuzinho = entityA.Artistas{
	Props: entityA.ArtistasPropsMKT{
		Nome:      "MUMUZINHO",
		Genero:    "SAMBA",
		Youtube:   "#mumuzinho",
		Instagram: "@mumuzinho",
	},
	Produtor: Produtor_Mumuzinho_Seed,
}

var Artista_Ferrugem = entityA.Artistas{
	Props: entityA.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Ferrugem_Seed,
}

var Artista_Pericles = entityA.Artistas{
	Props: entityA.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Pericles_Seed,
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
