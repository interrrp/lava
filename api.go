package main

import (
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

var api = map[string]lua.LGFunction{
	"lava.window.setFps":    windowSetFps,
	"lava.window.setTitle":  windowSetTitle,
	"lava.window.deltaTime": windowDeltaTime,

	"lava.draw.clear": drawClear,
	"lava.draw.text":  drawText,
	"lava.draw.rect":  drawRect,

	"lava.input.isKeyPressed":  isKeyPressed,
	"lava.input.isKeyDown":     isKeyDown,
	"lava.input.isKeyReleased": isKeyReleased,
	"lava.input.isKeyUp":       isKeyUp,
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

func createApi(state *lua.LState) {
	for name, fn := range api {
		parts := strings.Split(name, ".")
		if len(parts) < 2 {
			continue
		}

		// Create tables for each namespace level
		table := state.GetGlobal(parts[0])
		if table == lua.LNil {
			table = state.NewTable()
			state.SetGlobal(parts[0], table)
		}

		// Navigate through intermediate tables
		current := table.(*lua.LTable)
		for i := 1; i < len(parts)-1; i++ {
			next := current.RawGetString(parts[i])
			if next == lua.LNil {
				next = state.NewTable()
				current.RawSetString(parts[i], next)
			}
			current = next.(*lua.LTable)
		}

		// Set the function in the deepest table
		current.RawSetString(parts[len(parts)-1], state.NewFunction(fn))
	}
}
