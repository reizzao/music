package contractcomposicao

import "github.com/reizzao/music/api/modules/voz/contractvoz"

type Frases_Auxiliares struct {
	CategoriaFrases CategoriaFrases
	Frase           string
	VogalTermino    contractvoz.Vogal
}

type CategoriaFrases = uint64

const (
	APAIXONADO CategoriaFrases = 0
	SOFRENDO                   = 1
)
