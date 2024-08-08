package contractvoz

type Vogal struct {
	Vogal         string
	Caracteristas []string
	GrausCombina  GrausCombina
	Altura        AlturaVogal
}

type AlturaVogal = uint64

const (
	Media AlturaVogal = 0
	Baixa             = 1
	Alta              = 2
)

type GrausCombina = []SensacaoHarmonica
