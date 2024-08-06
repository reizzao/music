package seed

import "github.com/reizzao/music/api/modules/marketing"

var Artista_Mumuzinho = marketing.Artistas{
	Props: marketing.ArtistasPropsMKT{
		Nome:      "MUMUZINHO",
		Genero:    "SAMBA",
		Youtube:   "#mumuzinho",
		Instagram: "@mumuzinho",
	},
	Produtor: Produtor_Mumuzinho_Seed,
}

var Artista_Ferrugem = marketing.Artistas{
	Props: marketing.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Ferrugem_Seed,
}

var Artista_Pericles = marketing.Artistas{
	Props: marketing.ArtistasPropsMKT{
		Nome:      "FERRUGEM",
		Genero:    "SAMBA",
		Youtube:   "#ferrugem",
		Instagram: "@ferrugem",
	},
	Produtor: Produtor_Pericles_Seed,
}

/*
var Artista_TAL = marketing.Artistas{
	Props: marketing.MarketingRedesProps{
		Nome: "",
		Youtube: "",
		Instagram: "",
	},
}
*/
