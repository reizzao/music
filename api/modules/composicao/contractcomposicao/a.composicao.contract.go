package contractcomposicao

type Composicao struct {
	Musica         MusicaComposicao
	Categoria      CategoriaMusical
	Letra          string
	LetraTecnica   LetraTecnica
	DataComposicao string // todo : date
	Tom            string
}
