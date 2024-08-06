package entityB

import "github.com/reizzao/music/api/entitys/entityA"

type LetraTecnica struct {
	Frase []Frase
}

type Frase struct {
	PerguntaResposta []PerguntaResposta
}

type PerguntaResposta struct {
	Fato                 PropsArranjo
	Consequencia_do_Fato PropsArranjo
	Finaliza_Fato        PropsArranjo
	Momento              MomentoEstrofe
}

type MomentoEstrofe struct {
	Parte    string
	ClimaVoz string
}

type PropsArranjo struct {
	Letra           string
	CantoVogalFinal entityA.CantoVogalFinal
	Acorde          Acorde
}

type Acorde struct {
	Grau string
}
