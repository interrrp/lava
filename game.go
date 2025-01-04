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

	luaState.SetGlobal("drawText", luaState.NewFunction(func(lua *lua.LState) int {
		rl.DrawText(
			lua.ToString(1),
			int32(lua.ToInt(2)),
			int32(lua.ToInt(3)),
			int32(lua.ToInt(4)),
			rl.White,
		)
		return 0
	}))

	luaState.SetGlobal("clear", luaState.NewFunction(func(lua *lua.LState) int {
		rl.ClearBackground(rl.Black)
		return 0
	}))

	return game{lua: luaState, appDir: appDir}
}

func (g *game) close() {
	g.lua.Close()
}

func (g *game) run() error {
	rl.InitWindow(640, 360, "Lava App")
	defer rl.CloseWindow()

	scriptFile := path.Join(g.appDir, "script.lua")
	if err := g.lua.DoFile(scriptFile); err != nil {
		return err
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if err := g.lua.DoString("frame()"); err != nil {
			return err
		}
		rl.EndDrawing()
	}

	return g.lua.DoString("load()")
}
