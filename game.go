package main

import (
	"fmt"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type game struct {
	appDir string
	lua    *lua.LState
	api    *lua.LTable
}

func newGame(appDir string) game {
	luaState := lua.NewState()

	api := createApi(luaState)
	luaState.SetGlobal("lava", api)

	return game{
		appDir: appDir,
		lua:    luaState,
		api:    api,
	}
}

func (g *game) close() {
	g.lua.Close()
}

func (g *game) run(showFps bool) error {
	// rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(640, 360, "Lava App")
	defer rl.CloseWindow()

	if err := g.initializeLuaScript(); err != nil {
		return err
	}

	return g.gameLoop(showFps)
}

func (g *game) initializeLuaScript() error {
	scriptFile := path.Join(g.appDir, "script.lua")
	if err := g.lua.DoFile(scriptFile); err != nil {
		return fmt.Errorf("failed to load script.lua: %w", err)
	}

	loadFn := g.lua.GetField(g.api, "load")
	if loadFn == lua.LNil {
		return fmt.Errorf("load function not found")
	}
	if err := g.lua.CallByParam(lua.P{Fn: loadFn}); err != nil {
		return fmt.Errorf("failed to execute load(): %w", err)
	}

	return nil
}

func (g *game) gameLoop(showFps bool) error {
	frameFn := g.lua.GetField(g.api, "frame")
	if frameFn == lua.LNil {
		return fmt.Errorf("frame function not found")
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if err := g.lua.CallByParam(lua.P{Fn: frameFn}); err != nil {
			return fmt.Errorf("failed to execute frame(): %w", err)
		}

		if showFps {
			rl.DrawText(fmt.Sprintf("%d", rl.GetFPS()), 16, 16, 20, rl.Green)
		}

		rl.EndDrawing()
	}
	return nil
}
