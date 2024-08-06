package obraDoDestino

import (
	"github.com/reizzao/music/api/entitys/entityA"
	"github.com/reizzao/music/api/entitys/entityB"
	"github.com/reizzao/music/api/literals/seed"
)

var Musica_ObraDoDestino = entityA.Composicao{
	Musica: entityB.MusicaComposicao{
		Nome: "OBRA DO DESTINO",
		Compositor: []entityB.Compositor{
			seed.CompositorPaulinhoDC,
			seed.CompositorREIZAO,
		},
		Contatos: "#todo",

		Apresentacao: []entityB.Guia{
			seed.GuiaCaeto_300,
		},

		Gravacoes: []entityB.GravacoesComposicoes{
			seed.GravacaoDEFAULT,
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
					Momento: seed.Parte_A_Piano,
					Fato: entityB.PropsArranjo{
						Letra:           "SE É OBRA DO DESTINO",
						CantoVogalFinal: seed.CantoVogalFinal_Baixo,
						Acorde:          seed.Inicia_Repouso,
					},
					Consequencia_do_Fato: entityB.PropsArranjo{
						Letra:           "A GENTE SE CONHECER",
						CantoVogalFinal: seed.CantoVogalFinal_Medio,
						Acorde:          seed.Continua_Relativo,
					},
					Finaliza_Fato: entityB.PropsArranjo{
						Letra:           "NÃO HÁ PORQUE SE ESCONDER DE MIM.",
						CantoVogalFinal: seed.CantoVogalFinal_Medio,
						Acorde:          seed.Finaliza_SobeDistancia,
					},
				}, //
			},
			},
		},
	},
}
