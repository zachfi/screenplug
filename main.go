package main

import (
	"fmt"
	"log/slog"

	"github.com/vcraescu/go-xrandr"
)

type mode int

const (
	Unknown mode = iota
	Mobile
	Docked
	WideSingleOutput
	WideDoubleOutput
)

func main() {
	screens, err := xrandr.GetScreens()
	if err != nil {
		panic(err)
	}

	cmd := xrandr.Command()

	switch connectionMode(screens) {
	case Docked:
	case Mobile:
	case WideSingleOutput:
	case WideDoubleOutput:
	default:
		slog.Error("unhandled connectionMode")
	}

	err = cmd.Run()
	if err != nil {
		slog.Error("failed to run command")
		return
	}
}

func docked(screens xrandr.Screens) bool {
	for _, s := range screens {
		for _, m := range s.Monitors {
			fmt.Printf("m: %+v\n", m)
		}
	}

	return false
}

func connectionMode(screens xrandr.Screens) mode {
	for _, s := range screens {
		for _, m := range s.Monitors {
			fmt.Printf("m: %+v\n", m)
			if m.Size.Height == 330 && m.Size.Width == 800 {
			}
		}
	}

	return Docked
}
