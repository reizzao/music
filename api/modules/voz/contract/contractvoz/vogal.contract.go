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
	Baixa AlturaVogal = 1
	Alta  AlturaVogal = 2
)

type GrausCombina = []Grau
