package testercomposicao

import (
	"fmt"

	"github.com/reizzao/music/api/modules/composicao/literals/seed/composicoes/comp_obraDoDestino"
	usecase "github.com/reizzao/music/api/modules/composicao/usecase/usecasecomposicao"
)

func TesterComposicao() {
	obraDoDestino := usecase.New(comp_obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
