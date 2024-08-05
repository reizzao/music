package entityA

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
}
