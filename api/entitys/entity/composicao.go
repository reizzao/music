package entity

import "github.com/reizzao/music/api/entitys/entitySub"

type Composicao struct {
	Musica         entitySub.MusicaComposicao
	Categoria      entitySub.CategoriaMusical
	Letra          string
	DataComposicao string // todo : date
	Tom            string
	Harmonia       []entitySub.Harmonia
}
