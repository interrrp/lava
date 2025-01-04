package main

import (
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type game struct {
	lua    *lua.LState
	appDir string
}

func newGame(appDir string) game {
	luaState := lua.NewState()
	createDrawFunctions(luaState)
	createInputFunctions(luaState)

	return game{
		lua:    luaState,
		appDir: appDir,
	}
}

func (g *game) close() {
	g.lua.Close()
}

func (g *game) run() error {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(640, 360, "Lava App")
	defer rl.CloseWindow()

	scriptFile := path.Join(g.appDir, "script.lua")
	if err := g.lua.DoFile(scriptFile); err != nil {
		return err
	}

	if err := g.lua.DoString("load()"); err != nil {
		return err
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if err := g.lua.DoString("frame()"); err != nil {
			return err
		}
		rl.EndDrawing()
	}

	return nil
}
