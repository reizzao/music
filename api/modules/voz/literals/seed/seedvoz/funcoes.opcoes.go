package seedvoz

import (
	"github.com/reizzao/music/api/modules/voz/contract/contractvoz"
)

var Funcao_Inicio = contractvoz.OpcoesFuncoes{
	Caracteristicas: "Repousar",
	Graus_Normal:    Grau_1,
	Graus_Medio:     Grau_2,
	Graus_Forte:     Grau_6,
}

var Funcao_Meio = contractvoz.OpcoesFuncoes{
	Caracteristicas: "Continuar- Distanciar",
	Graus_Normal:    Grau_2,
	Graus_Medio:     Grau_4,
	Graus_Forte:     Grau_6,
}

var Funcao_Fim = contractvoz.OpcoesFuncoes{
	Caracteristicas: "Terminar",
	Graus_Normal:    Grau_3,
	Graus_Medio:     Grau_5,
	Graus_Forte:     Grau_7,
}
