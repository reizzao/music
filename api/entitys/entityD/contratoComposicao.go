package entityD

import "github.com/reizzao/music/api/entitys/entityE"

type ContratoComposicao struct {
	Vigente         bool
	DataContratacao string // TODO: date
	DataExpiracao   string // TODO: date
	Duracao         int
	Pagamentos      entityE.PagamentosLiberacoes
}
