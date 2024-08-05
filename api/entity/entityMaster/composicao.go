package entityMaster

import (
	"github.com/reizzao/music/api/entity/entitySub"
)

type Composicao struct {
	Musica    entitySub.MusicaComposicao
	Categoria entitySub.CategoriaMusical
}
