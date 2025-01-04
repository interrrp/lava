package main

import (
	"fmt"
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

	if err := g.initializeLuaScript(); err != nil {
		return err
	}

	return g.gameLoop()
}

func (g *game) initializeLuaScript() error {
	scriptFile := path.Join(g.appDir, "script.lua")
	if err := g.lua.DoFile(scriptFile); err != nil {
		return fmt.Errorf("failed to load script.lua: %w", err)
	}

	if err := g.lua.DoString("load()"); err != nil {
		return fmt.Errorf("failed to execute load(): %w", err)
	}

	return nil
}

func (g *game) gameLoop() error {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if err := g.lua.DoString("frame()"); err != nil {
			return fmt.Errorf("failed to execute frame(): %w", err)
		}
		rl.EndDrawing()
	}
	return nil
}
