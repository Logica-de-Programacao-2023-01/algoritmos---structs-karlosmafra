package main

import "fmt"

type musica struct {
	tit string
	art string
	dur float64
}

type Playlist struct {
	nome    string
	musicas []musica
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

	p1 := Playlist{nome: "Rock", musicas: []musica{m1, m2, m3}}
	p2 := Playlist{nome: "Pop", musicas: []musica{m4, m5, m6}}

	imprimir(p1, p2)
}

func imprimir(p1, p2 Playlist) {
	fmt.Print("Escolha uma playlist (1 ou 2): ")
	var p int
	fmt.Scan(&p)

	if p == 1 {
		fmt.Println("Nome da playlist:", p1.nome)

		var temp float64
		for m := range p1.musicas {
			temp += p1.musicas[m].dur
		}
		fmt.Println("Duração da playlist:", temp, "minutos")

		for m := range p1.musicas {
			fmt.Println(p1.musicas[m])
		}

	} else if p == 2 {
		fmt.Println("Nome da playlist", p2.nome)

		var temp float64
		for m := range p1.musicas {
			temp += p1.musicas[m].dur
		}
		fmt.Println("Duração da playlist:", temp, "minutos")

		for m := range p1.musicas {
			fmt.Println(p1.musicas[m])
		}
	}
}
