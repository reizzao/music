package seedvoz

import "github.com/reizzao/music/api/modules/voz/contract/contractvoz"

var Grau_1 = contractvoz.Grau{
	Nome:              "Grau_1",
	Regiao:            "Repouso",
	Similar_A_Grau:    "Grau_1 - ele Mesmo",
	Funcao:            "Repousar",
	Precisa_Resolucao: false,
}

var Grau_2 = contractvoz.Grau{
	Nome:              "Grau_2",
	Regiao:            "Continuacao",
	Similar_A_Grau:    "Grau_4",
	Funcao:            "Caminhar",
	Precisa_Resolucao: true,
}

var Grau_3 = contractvoz.Grau{
	Nome:              "Grau_3",
	Regiao:            "Repouso Tenso",
	Similar_A_Grau:    "Grau_1 + Forte",
	Funcao:            "Reiniciar e Reepousar no Lugar do Grau 1",
	Precisa_Resolucao: false,
}

var Grau_4 = contractvoz.Grau{
	Nome:              "Grau_4",
	Regiao:            "Distancia",
	Similar_A_Grau:    "Grau_2 + Forte",
	Funcao:            "Continuar Forte - Distanciar, Viajar",
	Precisa_Resolucao: true,
}

var Grau_5 = contractvoz.Grau{
	Nome:              "Grau_5",
	Regiao:            "Dominate Tensa",
	Similar_A_Grau:    "Grau_3 + Forte",
	Funcao:            "Resolver Acabar Finalizar",
	Precisa_Resolucao: true,
}

var Grau_6 = contractvoz.Grau{
	Nome:              "Grau_6",
	Regiao:            "Relativa",
	Similar_A_Grau:    "Grau_1 + Forte",
	Funcao:            "Iniciar Tensamente Forte no lugar do Grau 1",
	Precisa_Resolucao: false,
}

var Grau_7 = contractvoz.Grau{
	Nome:              "Grau_7",
	Regiao:            "Minuta, Repetina",
	Similar_A_Grau:    "Grau_5 + Forte",
	Funcao:            "Passagem para regiao Tensa - Pausa , Freio",
	Precisa_Resolucao: true,
}
