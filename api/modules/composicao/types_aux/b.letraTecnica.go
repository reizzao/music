package types_aux

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
	Momento              MomentoEstrofe
}

type MomentoEstrofe struct {
	Parte    string
	ClimaVoz string
}

type PropsArranjo struct {
	Letra           string
	CantoVogalFinal []string
	Acorde          Acorde
}

type Acorde struct {
	Grau string
}
