package composicaoSeed

import (
	"github.com/reizzao/music/api/entitys/entityA"
	"github.com/reizzao/music/api/entitys/entityB"
	"github.com/reizzao/music/api/entitys/entityC"
	"github.com/reizzao/music/api/entitys/entityD"
	"github.com/reizzao/music/api/entitys/entityE"
)

var FOO = entityA.Composicao{
	Musica: entityB.MusicaComposicao{
		Nome:       "FOO",
		Compositor: "Reizao",
		Contatos:   "#todo",
		Gravacoes: []entityB.GravacoesComposicoes{
			entityB.GravacoesComposicoes{
				Interprete:   "@cantorCaeto",
				TipoGravacao: "Video e Mp3",
				Data:         "03/08/2024",
				EstadoAtualTrabalho: entityC.EstadoAtualTrabalho{
					ContratoComposicao: entityD.ContratoComposicao{
						Vigente:         false,
						DataContratacao: "",
						DataExpiracao:   "",
						Duracao:         0,
						Pagamentos: entityE.PagamentosLiberacoes{
							ValorContratado: 0,
							EmAberto:        0,
							Pago:            0,
							Debito:          0,
						},
					},
				},
			},
		},
	},
	Categoria: entityB.CategoriaMusical{
		Nome: "Dolente",
	},
	Letra: `
	`,

	LetraTecnica: entityB.LetraTecnica{

		Frase: []entityB.Frase{

			entityB.Frase{PerguntaResposta: []entityB.PerguntaResposta{

				entityB.PerguntaResposta{
					Parte: entityB.Parte{
						Parte: "A", ClimaVoz: "PIANO",
					},
					Fato: entityB.Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
					Explicacao_do_Fato: entityB.Explicacao_do_Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
				}, //

				entityB.PerguntaResposta{
					Parte: entityB.Parte{
						Parte: "A", ClimaVoz: "PIANO",
					},
					Fato: entityB.Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
					Explicacao_do_Fato: entityB.Explicacao_do_Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
				}, //

			},
			},
		},
	},
}

/*
entityB.PerguntaResposta{
					Parte: entityB.Parte{
						Parte: "A", ClimaVoz: "PIANO",
					},
					Fato: entityB.Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
					Explicacao_do_Fato: entityB.Explicacao_do_Fato{
						Letra:   "",
						Acordes: []string{"", "", "", "", "", ""},
					},
				}, //

*/
