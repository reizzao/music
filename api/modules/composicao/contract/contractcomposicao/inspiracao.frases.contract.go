package contractcomposicao

import "github.com/reizzao/music/api/modules/voz/contract/contractvoz"

type Inspiracao_Frases struct {
	CategoriaFrases CategoriaFrases
	Frase           string
	VogalTermino    contractvoz.Vogal
}

type CategoriaFrases = uint64

const (
	APAIXONADO CategoriaFrases = 0
	SOFRENDO   CategoriaFrases = 1
)
