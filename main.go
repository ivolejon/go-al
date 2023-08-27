package main

import (
	"go-al/adapters/allsvenskan"
	"go-al/match"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

const MAX_WIDTH = 60

func Filter[T any](s []T, cond func(t T) bool) []T {
	res := []T{}
	for _, v := range s {
		if cond(v) {
			res = append(res, v)
		}
	}
	return res
}

func getViewportWidth() int {
	w, _ := terminal.Width()
	return min(int(w), MAX_WIDTH)
}
func getViewportHeight() int {
	h, _ := terminal.Height()
	return int(h)
}

func Render(list *widgets.List, data []match.Match) {
	finishedMatches := Filter(data, func(m match.Match) bool { return m.IsFinished() })
	upcomingMatches := Filter(data, func(m match.Match) bool { return m.IsNotPlayed() })

	for _, match := range finishedMatches {
		list.Rows = append(list.Rows, match.String(MAX_WIDTH))
	}

	for _, match := range upcomingMatches {
		list.Rows = append(list.Rows, match.String(MAX_WIDTH))
	}
	ui.Render(list)
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize UI: %v", err)
	}
	defer ui.Close()

	list := widgets.NewList()
	list.PaddingLeft = 1
	list.PaddingRight = 1
	list.Title = "Allsvenskan"
	list.WrapText = false
	list.SetRect(0, 0, getViewportWidth(), getViewportHeight())

	allsvenskan := allsvenskan.Adapter{}
	res := <-allsvenskan.Fetch()
	Render(list, res)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			list.SetRect(0, 0, min(payload.Width, MAX_WIDTH), payload.Height)
			ui.Render(list)
		}
	}
}
