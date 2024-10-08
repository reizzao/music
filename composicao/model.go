package composicao

import help "github.com/reizzao/music/helpers"

type ComposicaoModel struct {
	Request  Request
	Computed Computed
}

// Response : Nivel_1

type Request struct {
	DadosMusica DadosMusicaProps
	Motivacoes  MotivacoesProps

	Estrofe_A1 Estrofe_AS_Props
	// Estrofe_A2 Estrofe_AS_Props
}

// Props : Nivel_2

type DadosMusicaProps struct {
	Grau_EscalaNatural_Define_TOM help.Grau_EscalaNatural_Options
	EntradaPadraoVoz              help.EntradaPadraoVozOptions
}

type MotivacoesProps struct {
	Sentimento_Triste                                 bool
	Emocao_Central                                    help.Emocao_Central_Options
	Quem_eo_Personagem_Principal                      help.Personagens_Options
	RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal string
	RefraoFecha_Declaracao_O_Jeito_e                  string
}

// Estrofes

type Estrofe_AS_Props struct {
	Estrofe                                      help.Estrofes_Options
	Verso1_Comum_RimaFraca1_PodeRimar            Verso_Props
	Verso2_Comum_RimaFraca11                     Verso_Props
	Verso3_MenosComum_RimaForte2_ApontaResolucao Verso_Props
	Verso4_MenosComum_RimaForte22                Verso_Props
}

// Versos
type Verso_Props struct {
	FraseNumero     int
	Marca_Rima      bool
	Gatilhos_Verso  help.Gatilhos_Versos_Options
	TempoVerbal_Def help.TempoVerbal_Options
	Personagem_Def  help.Personagens_Options
	Verso           string
	// Resposta        Duas_PalavraPoetica
}

// type Duas_PalavraPoetica struct {
// 	PalavraPoetica1 string
// 	PalavraPoetica2 string
// }

type Computed = string
