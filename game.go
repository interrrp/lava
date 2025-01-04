package main

import (
	"path"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type game struct {
	lua    *lua.LState
	appDir string
}

func createDrawFunctions(state *lua.LState) {
	draw := state.NewTable()

	draw.RawSetString("clear", state.NewFunction(func(state *lua.LState) int {
		rl.ClearBackground(tableToColor(state.ToTable(1)))
		return 0
	}))

	draw.RawSetString("text", state.NewFunction(func(state *lua.LState) int {
		rl.DrawText(
			state.ToString(1),
			int32(state.ToInt(2)),
			int32(state.ToInt(3)),
			int32(state.ToInt(4)),
			tableToColor(state.ToTable(5)),
		)
		return 0
	}))

	state.SetGlobal("draw", draw)
}

func gatherErrors(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

func tableToColor(table *lua.LTable) rl.Color {
	r, rErr := strconv.Atoi(table.RawGetString("r").String())
	g, gErr := strconv.Atoi(table.RawGetString("g").String())
	b, bErr := strconv.Atoi(table.RawGetString("b").String())
	a, aErr := strconv.Atoi(table.RawGetString("a").String())

	if err := gatherErrors(rErr, gErr, bErr, aErr); err != nil {
		return rl.Black
	}

	return rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

func newGame(appDir string) game {
	luaState := lua.NewState()
	createDrawFunctions(luaState)

	return game{lua: luaState, appDir: appDir}
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

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if err := g.lua.DoString("frame()"); err != nil {
			return err
		}
		rl.EndDrawing()
	}

	return g.lua.DoString("load()")
}
