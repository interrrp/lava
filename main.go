package main

import "github.com/alexflint/go-arg"

func main() {
	var args struct {
		AppDir  string `arg:"required,positional"`
		ShowFps bool   `arg:"-f,--show-fps" default:"false"`
	}
	arg.MustParse(&args)

	game := newGame(args.AppDir)
	defer game.close()

	if err := game.run(args.ShowFps); err != nil {
		panic(err)
	}
}
