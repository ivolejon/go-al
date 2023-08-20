package main

import (
	"go-al/adapters/allsvenskan"
	"go-al/match"

	"log"
	// "strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

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
	return int(w)
}
func getViewportHeight() int {
	h, _ := terminal.Height()
	return int(h)
}

func Render(list *widgets.List, data []match.Match) {
	finishedMatches := Filter(data, func(m match.Match) bool { return m.IsFinished() })
	upcomingMatches := Filter(data, func(m match.Match) bool { return m.IsNotPlayed() })

	for _, match := range finishedMatches {
		list.Rows = append(list.Rows, match.String())
	}

	for _, match := range upcomingMatches {
		list.Rows = append(list.Rows, match.String())
	}
	ui.Render(list)
}

func main() {
	// allsvenskan := allsvenskan.Adapter{}
	// res := <-allsvenskan.Fetch()
	// fmt.Print(res)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	list := widgets.NewList()
	list.WrapText = false
	list.SetRect(0, 0, getViewportWidth(), getViewportHeight())

	allsvenskan := allsvenskan.Adapter{}
	res := <-allsvenskan.Fetch()
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	Render(list, res)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			list.SetRect(0, 0, payload.Width, payload.Height)
			ui.Render(list)
		}
	}
}
