package testerComposicao

import (
	"fmt"

	"github.com/reizzao/music/api/seed/composicaoSeed/obraDoDestino"
	"github.com/reizzao/music/api/usecases/composicao"
)

func TesterComposicao() {
	obraDoDestino := composicao.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}