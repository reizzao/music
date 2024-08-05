package entityA

type MarketingRedes struct {
	Artistas []Artistas
}

type MarketingRedesProps struct {
	Nome      string
	Genero    string
	Youtube   string
	Instagram string
}

type Gravadoras struct {
	Props MarketingRedesProps
}

type Produtores struct {
	Props MarketingRedesProps
}

type Escritorios struct {
	Props MarketingRedesProps
}

type Artistas struct {
	Props    MarketingRedesProps
	Produtor ProdutorProps
}

type ProdutorProps struct {
	Youtube   string
	Instagram string
}
