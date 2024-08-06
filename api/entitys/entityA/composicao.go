package entityA

import "github.com/reizzao/music/api/entitys/entityB"


type Composicao struct {
	Musica         entityB.MusicaComposicao
	Categoria      entityB.CategoriaMusical
	Letra          string
	LetraTecnica   entityB.LetraTecnica
	DataComposicao string // todo : date
	Tom            string
}

type CantoVogalFinal = []string
