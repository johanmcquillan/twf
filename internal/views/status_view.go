package views

import (
	"github.com/johanmcquillan/twf/internal/config"
	"github.com/johanmcquillan/twf/internal/state"
	term "github.com/johanmcquillan/twf/internal/terminal"
)

type statusView struct {
	config *config.TwfConfig
	state  *state.State
}

func NewStatusView(config *config.TwfConfig, state *state.State) term.View {
	return &statusView{
		config: config,
		state:  state,
	}
}

func (v *statusView) Position(totalRows int, totalCols int) term.Position {
	return term.Position{
		Top:  totalRows,
		Left: 1,
		Rows: 1,
		Cols: totalCols,
	}
}

func (v *statusView) HasBorder() bool {
	return false
}

func (v *statusView) ShouldRender() bool {
	return true
}

func (v *statusView) Render(p term.Position) []term.Line {
	line := term.NewLine(&term.Graphics{}, p.Cols)
	line.Append("", &term.Graphics{})
	return []term.Line{line}
}

func (v *statusView) GetCommands() map[string]term.Command {
	return map[string]term.Command{}
}
