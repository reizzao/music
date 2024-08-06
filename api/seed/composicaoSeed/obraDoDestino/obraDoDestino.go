package obraDoDestino

import (
	"github.com/reizzao/music/api/entitys/entityA"
	"github.com/reizzao/music/api/entitys/entityB"
	"github.com/reizzao/music/api/seed/seedMusical"
	"github.com/reizzao/music/api/seed/seedcompositor"
	"github.com/reizzao/music/api/seed/seeddefault"
	"github.com/reizzao/music/api/seed/seedguia"
)

var Musica_ObraDoDestino = entityA.Composicao{
	Musica: entityB.MusicaComposicao{
		Nome: "OBRA DO DESTINO",
		Compositor: []entityB.Compositor{
			seedcompositor.CompositorPaulinhoDC,
			seedcompositor.CompositorREIZAO,
		},
		Contatos: "#todo",

		Apresentacao: []entityB.Guia{
			seedguia.GuiaCaeto_300,
		},

		Gravacoes: []entityB.GravacoesComposicoes{
			seeddefault.GravacaoDEFAULT,
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
					Momento: seedMusical.Parte_A_Piano,
					Fato: entityB.PropsArranjo{
						Letra:           "SE É OBRA DO DESTINO",
						CantoVogalFinal: seedMusical.CantoVogalFinal_Baixo,
						Acorde:          seedMusical.Inicia_Repouso,
					},
					Consequencia_do_Fato: entityB.PropsArranjo{
						Letra:           "A GENTE SE CONHECER",
						CantoVogalFinal: seedMusical.CantoVogalFinal_Medio,
						Acorde:          seedMusical.Continua_Relativo,
					},
					Finaliza_Fato: entityB.PropsArranjo{
						Letra:           "NÃO HÁ PORQUE SE ESCONDER DE MIM.",
						CantoVogalFinal: seedMusical.CantoVogalFinal_Medio,
						Acorde:          seedMusical.Finaliza_SobeDistancia,
					},
				}, //
			},
			},
		},
	},
}
