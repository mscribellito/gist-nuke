package main

import (
	"os"

	"github.com/mscribellito/gist-nuke/cmd"
)

func main() {

	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}

}
