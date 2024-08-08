package contractvoz

type Voz struct {
	ParteMusica ParteMusica
}

type ParteMusica struct {
	Parte string
	Clima []ClimaParte
}

type Grau struct {
	Nome              string
	Regiao            string
	Similar_A_Grau    string
	Funcao            string
	Precisa_Resolucao bool
}

type OpcoesFuncoes struct {
	Caracteristicas string
	Graus_Normal    Grau
	Graus_Medio     Grau
	Graus_Forte     Grau
}

type ClimaParte struct {
	Nome         string
	GrausCombina []Grau
}
