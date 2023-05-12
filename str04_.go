package main

type musica struct {
	tit string
	art string
	dur float64
}

type Playlist struct {
	nome    string
	musicas []struct{}
}

func main() {
	m1 := musica{
		tit: "Come as You Are",
		art: "Nirvana",
		dur: 220,
	}
	m2 := musica{
		tit: "Rock You Like a Hurricane",
		art: "Scorpions",
		dur: 255,
	}
	m3 := musica{
		tit: "Start Me Up",
		art: "The Rolling Stones",
		dur: 215,
	}
	m4 := musica{
		tit: "Toxic",
		art: "Britney Spears",
		dur: 210,
	}
	m5 := musica{
		tit: "Beat It",
		art: "Michael Jackson",
		dur: 260,
	}
	m6 := musica{
		tit: "Take On Me",
		art: "A-ha",
		dur: 225,
	}

	p1 := Playlist{nome: "Rock", musicas: }
}
