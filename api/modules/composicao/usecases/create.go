package usecases

import "github.com/reizzao/music/api/modules/composicao"

func New(c composicao.Composicao) composicao.Composicao {
	return c
}
