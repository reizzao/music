package entitySub

import "github.com/reizzao/music/api/entitys/entitySubCooper"

type GravacoesComposicoes struct {
	Interprete          string
	TipoGravacao        string
	Data                string
	EstadoAtualTrabalho entitySubCooper.EstadoAtualTrabalho
}
