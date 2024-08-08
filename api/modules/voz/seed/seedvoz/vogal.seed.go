package seedvoz

import "github.com/reizzao/music/api/modules/voz/contractvoz"

var Vogal_A = contractvoz.Vogal{
	Vogal:         "A",
	Caracteristas: []string{"Explosao", "Libertacao", "Forte"},
	Altura:        contractvoz.Alta,
	GrausCombina: []contractvoz.SensacaoHarmonica{
		SensacaoRepouso, SensacaoTensao,
	},
}

// TODO CONTEUDO CADA CAMPO
var Vogal_E = contractvoz.Vogal{
	Vogal:         "E",
	Caracteristas: []string{"", "", ""},
	Altura:        contractvoz.Alta,
	GrausCombina: []contractvoz.SensacaoHarmonica{
		SensacaoRepouso,
	},
}

var Vogal_I = contractvoz.Vogal{
	Vogal:         "I",
	Caracteristas: []string{"", "", ""},
	Altura:        contractvoz.Alta,
	GrausCombina: []contractvoz.SensacaoHarmonica{
		SensacaoRepouso,
	},
}

var Vogal_O = contractvoz.Vogal{
	Vogal:         "O",
	Caracteristas: []string{"", "", ""},
	Altura:        contractvoz.Alta,
	GrausCombina: []contractvoz.SensacaoHarmonica{
		SensacaoRepouso,
	},
}

var Vogal_U = contractvoz.Vogal{
	Vogal:         "U",
	Caracteristas: []string{"", "", ""},
	Altura:        contractvoz.Alta,
	GrausCombina: []contractvoz.SensacaoHarmonica{
		SensacaoRepouso,
	},
}
