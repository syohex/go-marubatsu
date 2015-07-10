package main

import (
	"math/rand"
	"time"
)

func comSelectPosition(g *Game) int {
	var possibilities []int
	for i, c := range g.cells {
		if c == EMPTY {
			possibilities = append(possibilities, i)
		}
	}

	n := len(possibilities)
	if n == 0 {
		return possibilities[0]
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	return possibilities[r.Intn(n)]
}
