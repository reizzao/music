package seedvoz

import (
	"github.com/reizzao/music/api/modules/voz/contract/contractvoz"
)

var Parte_A_Baixo_Piano = contractvoz.ClimaParte{
	Nome: "Parte_A_Baixo_Piano",
	GrausCombina: []contractvoz.Grau{
		Grau_1,
	},
}

var Parte_B_Medio_Crescendo = contractvoz.ClimaParte{
	Nome: "Parte_B_Medio_Crescendo",
	GrausCombina: []contractvoz.Grau{
		Grau_4, Grau_6,
	},
}

var Parte_REFRAO_Forte = contractvoz.ClimaParte{
	Nome: "Parte_REFRAO_Forte",
	GrausCombina: []contractvoz.Grau{
		Grau_5, Grau_7,
	},
}
