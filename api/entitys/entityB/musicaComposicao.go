package entityB

type MusicaComposicao struct {
	Nome         string
	Compositores string
	Contatos     string //TODO : []ContatosComposicao
	Gravacoes    []GravacoesComposicoes
}
