package main

import "github.com/nsf/termbox-go"

const lineColor = termbox.ColorWhite
const cellWidth = 5
const promptY = 8
const messageBoxY = 12

func putString(msg string, x int, y int, fg termbox.Attribute) {
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, fg, termbox.ColorBlack)
	}
}

func renderBoard(size int) {
	bg := termbox.ColorBlack
	yLimit := size * 2
	xLimit := size*cellWidth - 1
	for y := 0; y < yLimit; y += 2 {
		for x := 0; x < xLimit; x++ {
			if x%cellWidth == (cellWidth - 1) {
				termbox.SetCell(x, y, '|', lineColor, bg)
			}
		}
	}

	for y := 1; y < yLimit-1; y += 2 {
		for x := 0; x < xLimit; x++ {
			var c rune
			if x%cellWidth == (cellWidth - 1) {
				c = '+'
			} else {
				c = '-'
			}
			termbox.SetCell(x, y, c, lineColor, bg)
		}
	}
}

func renderStones(g *Game) {
	for i := range g.cells {
		if g.cells[i] == EMPTY {
			continue
		}

		d := i / g.size
		m := i % g.size

		x := (m * cellWidth) + 1
		y := 2 * d

		var c rune
		if g.cells[i] == MARU {
			c = 'o'
		} else {
			c = 'x'
		}
		termbox.SetCell(x, y, c, termbox.ColorYellow, termbox.ColorBlack)
	}
}

func renderHint(g *Game) {
	hintX := (g.size*cellWidth - 1) + 3

	for i, c := range g.cells {
		if c != EMPTY {
			continue
		}

		d := i / g.size
		m := i % g.size

		x := (m * cellWidth) + 1
		y := 2 * d

		c := '1' + i
		termbox.SetCell(x+hintX, y, rune(c), termbox.ColorYellow, termbox.ColorBlack)
	}
}

func Render(g *Game) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	renderBoard(g.size)
	renderStones(g)
	renderHint(g)

	termbox.Flush()
}

func RenderPrompt(g *Game, msg string) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	renderBoard(g.size)
	renderStones(g)
	renderHint(g)

	putString(msg, 0, messageBoxY, termbox.ColorYellow)
	termbox.Flush()
}

func RenderWarning(g *Game, msg string) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	renderBoard(g.size)
	renderStones(g)
	renderHint(g)

	putString(msg, 0, messageBoxY, termbox.ColorMagenta)
	termbox.Flush()
}
