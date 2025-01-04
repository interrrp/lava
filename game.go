package main

import (
	"path"

	lua "github.com/yuin/gopher-lua"
)

type game struct {
	lua    *lua.LState
	appDir string
}

func newGame(appDir string) game {
	return game{
		lua:    lua.NewState(),
		appDir: appDir,
	}
}

func (g *game) close() {
	g.lua.Close()
}

func (g *game) run() error {
	scriptFile := path.Join(g.appDir, "script.lua")
	if err := g.lua.DoFile(scriptFile); err != nil {
		return err
	}

	return g.lua.DoString("load()")
}
