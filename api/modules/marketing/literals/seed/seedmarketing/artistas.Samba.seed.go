package seedmarketing

import "github.com/reizzao/music/api/modules/marketing/contract/contractmarketing"

var Artista_Mumuzinho = contractmarketing.Artistas{
	Props: contractmarketing.ArtistasPropsMKT{
		Nome:      "MUMUZINHO",
		Genero:    "SAMBA",
		Youtube:   "#mumuzinho",
		Instagram: "@mumuzinho",
	},
	Produtor: Produtor_Mumuzinho_Seed,
}

var Artista_Ferrugem = contractmarketing.Artistas{
	Props: contractmarketing.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Ferrugem_Seed,
}

var Artista_Pericles = contractmarketing.Artistas{
	Props: contractmarketing.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Pericles_Seed,
}

/*
var Artista_TAL = contractmarketing.Artistas{
	Props: edesProps{
		Nome: "",
		Youtube: "",
		Instagram: "",
	},
}
*/
