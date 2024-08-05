package entitySubCooperAux

import "github.com/reizzao/music/api/entitys/entitySubCooperAuxService"


type ContratoComposicao struct {
	Vigente         bool
	DataContratacao string // TODO: date
	DataExpiracao   string // TODO: date
	Duracao         int
	Pagamentos      entitySubCooperAuxService.PagamentosLiberacoes
}
