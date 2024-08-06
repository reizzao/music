package testercomposicao

import (
	"fmt"

	"github.com/reizzao/music/api/modules/composicao/literals/seed/composicoes/obraDoDestino"
	"github.com/reizzao/music/api/modules/composicao/usecase"
)

func TesterComposicao() {
	obraDoDestino := usecase.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
