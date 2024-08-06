package entityB

import "github.com/reizzao/music/api/entitys/entityC"

type GravacoesComposicoes struct {
	Interprete   string
	TipoGravacao string
	Data         string
	Controle     entityC.EstadoAtualTrabalho
}
