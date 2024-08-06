package composicaoSeed

import (
	"github.com/reizzao/music/api/entitys/entityA"
	"github.com/reizzao/music/api/entitys/entityB"
	"github.com/reizzao/music/api/entitys/entityC"
	"github.com/reizzao/music/api/entitys/entityD"
	"github.com/reizzao/music/api/entitys/entityE"
	"github.com/reizzao/music/api/seeds/seedMusical"
)

var TODO = entityA.Composicao{
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
SE É OBRA DO DESTINO A GENTE SE CONHECER
NÃO HÁ PORQUE SE ESCONDER DE MIM

O TEU JEITO NÃO ME ENGANA
SE FAZ DIFICIL MAS ME QUÉR

SÓ CONFUNDE O PENSAMENTO
POIS TENHO VOCÊ AQUI DENTRO

PERTO DO CORAÇÃO
QUE QUASE PARA AO TE VER
É DE BANDEJA ESSA PAIXÃO SÓ PRA VOCÊ...

NÃO MALTRATA VE SE ENTENDE
QUE UM AMOR TÃO LINDO ASSIM
NÃO É FACIL DE ACHAR
VEM QUE AQUI É SEU LUGAR

UM GRANDE AMOR
GUARDEI A VIDA INTEIRA PRA TE DAR
VEM LOGO DE UMA VEZ
NÃO SOU BRINQUEDO PRA VOCE BRINCAR

UNIR TEU CORPO AO MEU
NAQUELE DIA QUE TE CONHECI
FEZ RENASCER, FEZ DESPERTAR
A PAIXÃO EM MIM.
	`,

	LetraTecnica: entityB.LetraTecnica{

		Frase: []entityB.Frase{

			entityB.Frase{PerguntaResposta: []entityB.PerguntaResposta{

				entityB.PerguntaResposta{
					Parte: entityB.Parte{
						Parte: "A", ClimaVoz: "PIANO",
					},
					Fato: entityB.Fato{
						Letra:   "SE É OBRA DO DESTINO A GENTE SE CONHECER",
						Acordes: []entityB.Cadencia{seedMusical.Cadencia_InicioRelativo_16},
					},
					Explicacao_do_Fato: entityB.Explicacao_do_Fato{
						Letra:   "NÃO HÁ PORQUE SE ESCONDER DE MIM",
						Acordes: []entityB.Cadencia{seedMusical.Cadencia_Finaliza_SobeeTenso},
					},
				}, //
			},
			},
		},
	},
}
