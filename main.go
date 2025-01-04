package main

import "github.com/alexflint/go-arg"

func main() {
	var args struct {
		AppDir string `arg:"required"`
	}
	arg.MustParse(&args)

	game := newGame(args.AppDir)
	defer game.close()

	if err := game.run(); err != nil {
		panic(err)
	}
}
