package main

import (
	"os"
	"flag"

	"golang.org/x/mod/modfile"
)

var (
	showIndirect = flag.Bool("indirect", false, "show indirect dependencies")
	showVersion = flag.Bool("version", false, "show dependencies version")
)

func main() {
	flag.Parse()

	content, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	file, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		panic(err)
	}

	for _, module := range file.Require {
		if module == nil {
			panic("got nil module, check the file")
		}
		if module.Indirect && showIndirect != nil && !*showIndirect {
			continue
		}
		if showVersion != nil && *showVersion {
			os.Stdout.WriteString(module.Mod.Path + " " + module.Mod.Version + "\n")
		} else {
			os.Stdout.WriteString(module.Mod.Path + "\n")
		}
	}
}
