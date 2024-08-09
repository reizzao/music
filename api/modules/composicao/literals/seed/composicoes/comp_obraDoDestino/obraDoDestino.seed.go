package comp_obraDoDestino

// TODO ARRUMAR

import (
	"github.com/reizzao/music/api/modules/adm/literals/seed/seedadm"
	"github.com/reizzao/music/api/modules/composicao/contract/contractcomposicao"
	"github.com/reizzao/music/api/modules/composicao/literals/seed/seedcomposicao"
	"github.com/reizzao/music/api/modules/marketing/contract/contractmarketing"
	"github.com/reizzao/music/api/modules/marketing/literals/seed/seedmarketing"
)

var Musica_ObraDoDestino = contractcomposicao.Composicao{
	Musica: contractcomposicao.MusicaComposicao{
		Nome: "OBRA DO DESTINO",
		Compositor: []contractcomposicao.Compositor{
			seedcomposicao.CompositorPaulinhoDC,
			seedcomposicao.CompositorREIZAO,
		},
		Contatos: "#todo",

		Apresentacao: []contractmarketing.Guia{
			seedmarketing.GuiaCaeto_300,
		},

		Gravacoes: []contractcomposicao.GravacoesComposicoes{
			seedadm.GravacaoDEFAULT,
		},
	},

	Categoria: contractcomposicao.CategoriaMusical{
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

	LetraTecnica: contractcomposicao.LetraTecnica{

		Frase: []contractcomposicao.Frase{

			{
				Fato_Consequencia_Conclusao: []contractcomposicao.Props_Frase{

					{
						// Momento: seedmusical.Parte_A_Piano,
						Fato_GRAVE: contractcomposicao.PropsFraseArranjo{
							Letra: "SE É OBRA DO DESTINO",
							// CantoVogalFinal: seedmusical.// CantoVogalFinal_Baixo,
							// Acorde:          seedmusical.Inicia_Repouso,
						},
						Consequencia_do_Fato_AGUDO: contractcomposicao.PropsFraseArranjo{
							Letra: "A GENTE SE CONHECER",
							// CantoVogalFinal: seedmusical.// CantoVogalFinal_Medio,
							// Acorde:          seedmusical.Continua_Relativo,
						},
						Conclusao_Resolve_MEDIO: contractcomposicao.PropsFraseArranjo{
							Letra: "NÃO HÁ PORQUE SE ESCONDER DE MIM.",
							// CantoVogalFinal: seedmusical.// CantoVogalFinal_Medio,
							// Acorde:          seedmusical.Finaliza_SobeDistancia,
						},
					}, //
				},
			},
		},
	},
}
