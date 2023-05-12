package main

type Filme struct {
	tit   string
	dir   string
	ano   int
	notas []float64
}

func main() {
	f := Filme{
		tit:   "",
		dir:   "",
		ano:   2000,
		notas: []float64{},
	}
}

func addNota(f Filme, nota float64) Filme {
	f.notas = append(f.notas, nota)
}

func remNota(f1 Filme) {

}
