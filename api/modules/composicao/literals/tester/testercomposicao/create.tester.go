package testercomposicao

import (
	"fmt"

	"github.com/reizzao/music/api/modules/composicao/literals/seed/composicoes/obraDoDestino"
	usecasess "github.com/reizzao/music/api/modules/composicao/usecases"
)

func TesterComposicao() {
	obraDoDestino := usecasess.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
