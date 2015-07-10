package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nsf/termbox-go"
)

func askRepeat(eventCh <-chan termbox.Event) bool {
	for {
		select {
		case ev := <-eventCh:
			if ev.Type == termbox.EventKey {
				switch ev.Ch {
				case 'y':
					return true
				case 'n':
					return false
				}
			}
		}
	}

	return false
}

func continueMessage(r Result) string {
	return fmt.Sprintf("%s Replay(y/n)? ", r.toString())
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer termbox.Close()

	eventCh := make(chan termbox.Event)
	go func() {
		for {
			eventCh <- termbox.PollEvent()
		}
	}()

	var size int
	if len(os.Args) > 1 {
		size, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid size %s\n", size)
			os.Exit(1)
		}
	} else {
		size = 3
	}

	g := NewGame(size)
	Render(g)
LOOP:
	for {
		select {
		case ev := <-eventCh:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyEsc:
					return
				case ev.Ch >= '1' && ev.Ch <= '9':
					pos := int(ev.Ch) - '1'
					if !g.canPut(pos) {
						msg := fmt.Sprintf("Can't put: '%c'", ev.Ch)
						RenderWarning(g, msg)
						continue LOOP
					}

					result := g.Put(pos)
					if result != CONTINUE {
						RenderPrompt(g, continueMessage(result))
						if askRepeat(eventCh) {
							g.reset()
							Render(g)
							continue LOOP
						} else {
							return
						}
					}
				default:
					msg := fmt.Sprintf("InvalidKey: '%c'", ev.Ch)
					RenderWarning(g, msg)
					continue
				}
			}
		}

		Render(g)

		pos := comSelectPosition(g)
		result := g.Put(pos)
		Render(g)
		if result != CONTINUE {
			RenderPrompt(g, continueMessage(result))
			if askRepeat(eventCh) {
				g.reset()
			} else {
				return
			}
		}
	}
}
