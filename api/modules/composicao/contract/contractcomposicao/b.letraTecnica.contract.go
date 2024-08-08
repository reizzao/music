package contractcomposicao

import "github.com/reizzao/music/api/modules/musical/contract/contractmusical"

type LetraTecnica struct {
	Frase []Frase
}

type Frase struct {
	PerguntaResposta []PerguntaResposta
}

type PerguntaResposta struct {
	Fato                 PropsArranjo
	Consequencia_do_Fato PropsArranjo
	O_Que_Resolve        PropsArranjo
	Momento              contractmusical.MomentoEstrofe
}

type PropsArranjo struct {
	Letra           string
	CantoVogalFinal []string
	Acorde          contractmusical.Acorde
}
