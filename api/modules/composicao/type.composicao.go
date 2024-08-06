package composicao

import "github.com/reizzao/music/api/modules/composicao/types_aux"


type Composicao struct {
	Musica         types_aux.MusicaComposicao
	Categoria      types_aux.CategoriaMusical
	Letra          string
	LetraTecnica   types_aux.LetraTecnica
	DataComposicao string // todo : date
	Tom            string
}
