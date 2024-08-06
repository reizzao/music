package testerComposicao

import (
	"fmt"

	"github.com/reizzao/music/api/seeds/composicaoSeed"
	"github.com/reizzao/music/api/usecases/composicao"
)

func TesterComposicao() {
	obraDoDestino := composicao.New(composicaoSeed.ObraDoDestino)

	fmt.Println(obraDoDestino)
}