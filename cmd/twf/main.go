package main

import (
	"fmt"
	"regexp"

	"github.com/johanmcquillan/twf/internal/config"
	"github.com/johanmcquillan/twf/internal/filetree"
	"github.com/johanmcquillan/twf/internal/state"
	"github.com/johanmcquillan/twf/internal/terminal"
	"github.com/johanmcquillan/twf/internal/views"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := config.GetConfig()

	if config.LogLevel != "" {
		var level zapcore.Level
		if err := level.UnmarshalText([]byte(config.LogLevel)); err != nil {
			panic(err)
		}
		logger, err := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Encoding:    "console",
			OutputPaths: []string{"/tmp/twf.log"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:  "message",
				LevelKey:    "level",
				TimeKey:     "time",
				EncodeLevel: zapcore.CapitalColorLevelEncoder,
				EncodeTime:  zapcore.RFC3339TimeEncoder,
			},
		}.Build()
		if err != nil {
			panic(err)
		}
		defer logger.Sync()
		zap.ReplaceGlobals(logger)
	} else {
		zap.ReplaceGlobals(zap.NewNop())
	}

	zap.L().Info("Starting twf.")

	tree, err := filetree.InitFileTree(config.Dir)
	if err != nil {
		panic(err)
	}
	state := state.State{
		Root:   tree,
		Cursor: tree,
	}

	var ignore *regexp.Regexp
	if config.AutoexpandIgnore != "" {
		ignore, err = regexp.Compile(config.AutoexpandIgnore)
		if err != nil {
			panic(err)
		}
	}
	if err := state.AutoExpand(config.AutoexpandDepth, ignore); err != nil {
		panic(err)
	}
	if config.LocatePath != "" {
		err = state.LocatePath(config.LocatePath)
		if err != nil {
			panic(err)
		}
	}
	views := []terminal.View{
		views.NewTreeView(config, &state),
		views.NewPreviewView(config, &state),
		views.NewStatusView(config, &state),
	}

	t, err := terminal.OpenTerm(&config.Terminal)
	if err != nil {
		panic(err)
	}
	err = t.StartLoop(config.Keybindings, views)
	t.Close()
	if err != nil {
		panic(err)
	}

	for _, node := range state.Selection {
		fmt.Println(node.AbsPath)
	}
	zap.L().Info("Stopping twf.")
}
