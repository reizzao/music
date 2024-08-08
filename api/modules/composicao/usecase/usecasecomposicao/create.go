package usecase

import "github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"

func New(c contractcomposicao.Composicao) contractcomposicao.Composicao {
	return c
}
