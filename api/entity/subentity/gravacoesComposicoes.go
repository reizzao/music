package subentity

import "github.com/reizzao/music/api/entity/entitySubCooper"

type GravacoesComposicoes struct {
	Interprete          string
	TipoGravacao        string
	Data                string
	EstadoAtualTrabalho entitySubCooper.EstadoAtualTrabalho
}
