package entitySubCooperAux

import "github.com/reizzao/music/api/entity/entitySubCooperAux2"

type ContratoComposicao struct {
	Vigente         bool
	DataContratacao string // TODO: date
	DataExpiracao   string // TODO: date
	Duracao         int
	Pagamentos      entitySubCooperAux2.PagamentosLiberacoes
}
