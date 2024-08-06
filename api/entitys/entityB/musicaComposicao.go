package entityB

type MusicaComposicao struct {
	Nome       string
	Compositor string
	Contatos   string //TODO : []ContatosComposicao
	Gravacoes  []GravacoesComposicoes
}
