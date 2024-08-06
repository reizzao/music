package testerComposicao

import (
	"fmt"

	"github.com/reizzao/music/api/literals/seed/obras/obraDoDestino"
	"github.com/reizzao/music/api/usecases/composicao"
)

func TesterComposicao() {
	obraDoDestino := composicao.New(obraDoDestino.Musica_ObraDoDestino)

	fmt.Println(obraDoDestino)
}
