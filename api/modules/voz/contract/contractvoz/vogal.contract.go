package contractvoz

type Vogal struct {
	Vogal         string
	Caracteristas []string
	GrausCombina  GrausCombina
	Altura        AlturaVogal
}

type AlturaVogal = uint64

const (
	Media AlturaVogal = iota
	Baixa
	Alta
)

type GrausCombina = []Grau
