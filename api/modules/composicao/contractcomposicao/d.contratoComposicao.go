package contractcomposicao

type ContratoComposicao struct {
	Vigente         bool
	DataContratacao string // TODO: date
	DataExpiracao   string // TODO: date
	Duracao         int
	Pagamentos      PagamentosLiberacoes
}
