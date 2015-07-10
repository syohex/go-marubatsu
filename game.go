package main

import "log"

const (
	MARU = iota
	BATSU
	EMPTY
)

type Result int
const (
	WIN Result  = iota
	LOOSE
	DRAW
	CONTINUE
)

func (r Result) toString() string {
	switch r {
	case WIN:
		return "You win!!"
	case LOOSE:
		return "You loose!!"
	case DRAW:
		return "Draw game!!"
	default:
		return ""
	}
}

type Cell int

type Game struct {
	cells []Cell
	size  int
	turn  int
}

func NewGame(n int) *Game {
	g := &Game{
		size: n,
	}

	g.reset()
	return g
}

func (g *Game) reset() {
	size := g.size * g.size
	g.cells = make([]Cell, size)
	for i := 0; i < size; i++ {
		g.cells[i] = EMPTY
	}
	g.turn = 0
}

func (g *Game) canPut(n int) bool {
	limit := g.size * g.size
	if n > limit {
		log.Printf("[Assert] Invalid position=%d(>= %d)\n", n, limit)
	}
	return g.cells[n] == EMPTY
}

func (g *Game) isGameSet() bool {
	for i := 0; i < g.size*g.size; i += 3 {
		if g.cells[i] == g.cells[i+1] && g.cells[i+1] == g.cells[i+2] {
			return false
		}
	}

	for _, c := range g.cells {
		if c == EMPTY {
			return false
		}
	}

	return true
}

func (g *Game) checkGame(pos int) bool {
	d := pos / g.size
	m := pos % g.size
	c := g.cells[pos]

	total := 0
	for i := 0; i < g.size; i++ {
		if g.cells[d+i] == c {
			total++
		}
	}

	if total == g.size {
		return true
	}

	total = 0
	for i := 0; i < g.size; i++ {
		if g.cells[m+g.size*i] == c {
			total++
		}
	}

	if total == g.size {
		return true
	}

	if pos%(g.size+1) == 0 {
		total = 0
		for i := 0; i < g.size; i++ {
			if g.cells[(g.size+1)*i] == c {
				total++
			}
		}

		if total == g.size {
			return true
		}
	}

	if pos != 0 && pos%(g.size-1) == 0 {
		total = 0
		for i := 1; i <= g.size; i++ {
			if g.cells[(g.size-1)*i] == c {
				total++
			}
		}

		if total == g.size {
			return true
		}
	}

	return false
}

func (g *Game) Put(pos int) Result {
	var c Cell
	if g.turn%2 == 0 {
		c = MARU
	} else {
		c = BATSU
	}

	g.cells[pos] = c

	if g.checkGame(pos) {
		if g.turn % 2 == 0 {
			return WIN
		} else {
			return LOOSE
		}
	}

	g.turn++
	if g.turn == g.size*g.size {
		return DRAW
	}

	return CONTINUE
}
