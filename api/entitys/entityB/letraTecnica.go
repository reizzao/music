package entityB

type LetraTecnica struct {
	Parte            Parte
	PerguntaResposta []PerguntaResposta
}

type PerguntaResposta struct {
	Fato               string
	Explicacao_do_Fato string
}

type Parte struct {
	Parte    string
	ClimaVoz string
}
