package testercomposicao

import (
	"fmt"

	"github.com/reizzao/music/api/modules/composicao/literals/seed/composicoes/obraDoDestino"
	usecase "github.com/reizzao/music/api/modules/composicao/usecase/usecasecomposicao"
)

func TesterComposicao() {
	obraDoDestino := usecase.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
