package contractmarketing

type MarketingRedes struct {
	Artistas []Artistas
}

type Artistas struct {
	Props    ArtistasPropsMKT
	Produtor ProdutorProps
}

type ArtistasPropsMKT struct {
	Nome      string
	Genero    string
	Youtube   string
	Instagram string
}

type ProdutorProps struct {
	Youtube   string
	Instagram string
	Email     string
}

type RedesCompositor struct {
	Instagram string
	Youtube   string
	PaginaWeb string
}

type Guia struct {
	Interprete string
	Custo      float64
}
