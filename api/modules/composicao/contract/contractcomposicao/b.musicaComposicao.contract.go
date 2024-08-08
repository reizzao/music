package contractcomposicao

import "github.com/reizzao/music/api/modules/marketing/contract/contractmarketing"

type MusicaComposicao struct {
	Nome         string
	Compositor   []Compositor
	Contatos     string //TODO : []ContatosComposicao
	Apresentacao []contractmarketing.Guia
	Gravacoes    []GravacoesComposicoes
}

type Compositor struct {
	NomeArtistico string
	Redes         contractmarketing.RedesCompositor
}
