package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

func createStdlib(state *lua.LState) {
	stdlib := state.NewTable()
	state.SetGlobal("lava", stdlib)

	window := state.NewTable()
	stdlib.RawSetString("window", window)
	window.RawSetString("setFps", state.NewFunction(windowSetFps))
	window.RawSetString("setTitle", state.NewFunction(windowSetTitle))
	window.RawSetString("deltaTime", state.NewFunction(windowDeltaTime))

	draw := state.NewTable()
	stdlib.RawSetString("draw", draw)
	draw.RawSetString("clear", state.NewFunction(drawClear))
	draw.RawSetString("text", state.NewFunction(drawText))
	draw.RawSetString("rect", state.NewFunction(drawRect))

	input := state.NewTable()
	stdlib.RawSetString("input", input)
	input.RawSetString("isKeyPressed", state.NewFunction(isKeyPressed))
	input.RawSetString("isKeyDown", state.NewFunction(isKeyDown))
	input.RawSetString("isKeyReleased", state.NewFunction(isKeyReleased))
	input.RawSetString("isKeyUp", state.NewFunction(isKeyUp))
}

func windowSetFps(state *lua.LState) int {
	rl.SetTargetFPS(int32(state.ToInt(1)))
	return 0
}

func windowSetTitle(state *lua.LState) int {
	rl.SetWindowTitle(state.ToString(1))
	return 0
}

func drawClear(state *lua.LState) int {
	rl.ClearBackground(tableToColor(state.ToTable(1)))
	return 0
}

func drawText(state *lua.LState) int {
	rl.DrawText(
		state.ToString(1),
		int32(state.ToInt(2)),
		int32(state.ToInt(3)),
		int32(state.ToInt(4)),
		tableToColor(state.ToTable(5)),
	)
	return 0
}

func drawRect(state *lua.LState) int {
	rl.DrawRectangle(
		int32(state.ToInt(1)),
		int32(state.ToInt(2)),
		int32(state.ToInt(3)),
		int32(state.ToInt(4)),
		tableToColor(state.ToTable(5)),
	)
	return 0
}

func windowDeltaTime(state *lua.LState) int {
	state.Push(lua.LNumber(rl.GetFrameTime()))
	return 1
}

func isKeyPressed(state *lua.LState) int {
	key := state.ToInt(1)
	state.Push(lua.LBool(rl.IsKeyPressed(int32(key))))
	return 1
}

func isKeyDown(state *lua.LState) int {
	key := state.ToInt(1)
	state.Push(lua.LBool(rl.IsKeyDown(int32(key))))
	return 1
}

func isKeyReleased(state *lua.LState) int {
	key := state.ToInt(1)
	state.Push(lua.LBool(rl.IsKeyReleased(int32(key))))
	return 1
}

func isKeyUp(state *lua.LState) int {
	key := state.ToInt(1)
	state.Push(lua.LBool(rl.IsKeyUp(int32(key))))
	return 1
}

func tableToColor(table *lua.LTable) rl.Color {
	components := []string{"r", "g", "b", "a"}
	values := make([]int, 4)

	for i, component := range components {
		value, err := strconv.Atoi(table.RawGetString(component).String())
		if err != nil {
			return rl.Black
		}
		values[i] = value
	}

	return rl.NewColor(uint8(values[0]), uint8(values[1]), uint8(values[2]), uint8(values[3]))
}
