package types_aux

type MusicaComposicao struct {
	Nome         string
	Compositor   []Compositor
	Contatos     string //TODO : []ContatosComposicao
	Apresentacao []Guia
	Gravacoes    []GravacoesComposicoes
}

type Compositor struct {
	NomeArtistico string
	Redes         RedesCompositor
}

type Guia struct {
	Interprete string
	Custo      float64
}

type RedesCompositor struct {
	Instagram string
	Youtube   string
	PaginaWeb string
}
