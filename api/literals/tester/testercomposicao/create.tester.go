package testercomposicao

import (
	"fmt"

	"github.com/reizzao/music/api/literals/seed/composicoes/obraDoDestino"
	"github.com/reizzao/music/api/usecases/composicao"
)

func TesterComposicao() {
	obraDoDestino := composicao.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
