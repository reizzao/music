package entityB

type LetraTecnica struct {
	Parte Parte
	Frase []Frase
}

type Frase struct {
	PerguntaResposta []PerguntaResposta
}

type PerguntaResposta struct {
	Fato               Fato
	Explicacao_do_Fato Explicacao_do_Fato
}

type Parte struct {
	Parte    string
	ClimaVoz string
}

type Fato struct {
	Letra   string
	Acordes []string
}
type Explicacao_do_Fato struct {
	Letra   string
	Acordes []string
}
