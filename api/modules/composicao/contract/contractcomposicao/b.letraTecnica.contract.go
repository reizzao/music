package contractcomposicao

import "github.com/reizzao/music/api/modules/voz/contract/contractvoz"

type LetraTecnica struct {
	Frase []Frase
}

type Frase struct {
	Fato_Consequencia_Conclusao []Props_Frase
}

type Props_Frase struct {
	Fato_GRAVE                 PropsFraseArranjo
	Consequencia_do_Fato_AGUDO PropsFraseArranjo
	Conclusao_Resolve_MEDIO    PropsFraseArranjo
}

type PropsFraseArranjo struct {
	Letra     string
	Altura    Altura
	OpcaoGrau []contractvoz.Grau
}

type Altura = uint

const (
	Grave Altura = iota
	Agudo
	Medio
)
