package views

import (
	"math"
	"strings"

	"github.com/johanmcquillan/twf/internal/config"
	"github.com/johanmcquillan/twf/internal/filetree"
	"github.com/johanmcquillan/twf/internal/state"
	term "github.com/johanmcquillan/twf/internal/terminal"
)

type treeView struct {
	config     *config.TwfConfig
	state      *state.State
	lineByPath map[string]int
	rows       int
	scroll     int
}

func NewTreeView(config *config.TwfConfig, state *state.State) term.View {
	return &treeView{
		config: config,
		state:  state,
	}
}

func (v *treeView) Position(totalRows int, totalCols int) term.Position {
	if v.config.Preview.Enabled && !v.state.Cursor.IsDir() {
		return term.Position{
			Top:  1,
			Left: 1,
			Rows: totalRows - 1,
			Cols: int(math.Ceil(float64(totalCols) / 2.0)),
		}
	} else {
		return term.Position{
			Top:  1,
			Left: 1,
			Rows: totalRows - 1,
			Cols: totalCols,
		}
	}
}

func (v *treeView) HasBorder() bool {
	return false
}

func (v *treeView) ShouldRender() bool {
	return true
}

func (v *treeView) renderNode(
	node *filetree.FileTree,
	indentation int,
	maxLength int,
) term.Line {
	line := term.NewLine(&term.Graphics{}, maxLength)
	line.Append(strings.Repeat("  ", indentation), nil)

	graphics := term.Graphics{}
	if node.IsDir() {
		if g, ok := v.config.Graphics["tree:dir"]; ok {
			graphics.Merge(g)
		}
	}
	if node == v.state.Cursor {
		if g, ok := v.config.Graphics["tree:cursor"]; ok {
			graphics.Merge(g)
		}
	}

	if node.IsDir() {
		if node.Expanded() {
			line.Append("▼ ", &graphics)
		} else {
			line.Append("▶ ", &graphics)
		}
	}
	line.Append(node.Name(), &graphics)
	return line
}

func (v *treeView) Render(p term.Position) []term.Line {
	lines := []term.Line{}
	v.rows = p.Rows
	v.lineByPath = make(map[string]int)
	v.state.Root.Traverse(true, filetree.ByTypeAndName, func(tree *filetree.FileTree, depth int) error {
		line := v.renderNode(tree, depth, p.Cols)
		v.lineByPath[tree.AbsPath] = len(lines)
		lines = append(lines, line)
		return nil
	})
	v.scroll = v.scrollForPath(v.state.Cursor.AbsPath)
	return lines[v.scroll:]
}

func (v *treeView) scrollForPath(path string) int {
	targetLine := v.lineByPath[path]
	if targetLine < v.scroll {
		return targetLine
	} else if targetLine >= v.scroll+v.rows {
		return targetLine - v.rows + 1
	} else {
		return v.scroll
	}
}

func (v *treeView) GetCommands() map[string]term.Command {
	return map[string]term.Command{
		"tree:prev":           v.prev,
		"tree:next":           v.next,
		"tree:open":           v.open,
		"tree:close":          v.close,
		"tree:toggle":         v.toggle,
		"tree:toggleAll":      v.toggleAll,
		"tree:openAll":        v.openAll,
		"tree:closeAll":       v.closeAll,
		"tree:parent":         v.parent,
		"tree:locateExternal": v.locateExternal,
		"tree:selectPath":     v.selectPath,
	}
}

func (v *treeView) selectPath(helper term.TerminalHelper, args ...interface{}) error {
	v.state.Selection = append(v.state.Selection, v.state.Cursor)
	return nil
}

func (v *treeView) prev(helper term.TerminalHelper, args ...interface{}) error {
	prev, err := v.state.Cursor.Prev(true, filetree.ByTypeAndName)
	if err != nil {
		return err
	}
	if prev != nil {
		v.state.Cursor = prev
	}
	return nil
}

func (v *treeView) next(helper term.TerminalHelper, args ...interface{}) error {
	next, err := v.state.Cursor.Next(true, filetree.ByTypeAndName)
	if err != nil {
		return err
	}
	if next != nil {
		v.state.Cursor = next
	}
	return nil
}

func (v *treeView) open(helper term.TerminalHelper, args ...interface{}) error {
	return v.state.Cursor.Expand()
}

func (v *treeView) close(helper term.TerminalHelper, args ...interface{}) error {
	v.state.Cursor.Collapse()
	return nil
}

func (v *treeView) toggle(helper term.TerminalHelper, args ...interface{}) error {
	if v.state.Cursor.Expanded() {
		v.state.Cursor.Collapse()
		return nil
	} else {
		return v.state.Cursor.Expand()
	}
}

func (v *treeView) toggleAll(helper term.TerminalHelper, args ...interface{}) error {
	expanded := v.state.Cursor.Expanded()
	return v.state.Cursor.Traverse(false, nil, func(tree *filetree.FileTree, _ int) error {
		if expanded {
			tree.Collapse()
			return nil
		} else {
			return tree.Expand()
		}
	})
}

func (v *treeView) openAll(helper term.TerminalHelper, args ...interface{}) error {
	return v.state.Cursor.Traverse(false, nil, func(tree *filetree.FileTree, _ int) error {
		return tree.Expand()
	})
}

func (v *treeView) closeAll(helper term.TerminalHelper, args ...interface{}) error {
	return v.state.Cursor.Traverse(false, nil, func(tree *filetree.FileTree, _ int) error {
		tree.Collapse()
		return nil
	})
}

func (v *treeView) parent(helper term.TerminalHelper, args ...interface{}) error {
	parent := v.state.Cursor.Parent()
	if parent != nil {
		v.state.Cursor = parent
	}
	return nil
}

func (v *treeView) locateExternal(helper term.TerminalHelper, args ...interface{}) error {
	content, err := helper.ExecuteInTerminal(v.config.TreeView.LocateCommand)
	if err != nil {
		return err
	}
	v.state.LocatePath(strings.TrimSpace(content))
	return nil
}
