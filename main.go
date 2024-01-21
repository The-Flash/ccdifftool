package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/The-Flash/ccdifftool/diff"
)

func main() {
	flag.Parse()

	diffTool := diff.NewDiffTool()

	filepath1 := flag.Arg(0)
	filepath2 := flag.Arg(1)

	content1, err := os.ReadFile(filepath1)
	if err != nil {
		panic(fmt.Sprintf("cannot open %s", filepath1))
	}
	content2, err := os.ReadFile(filepath2)
	if err != nil {
		panic(fmt.Sprintf("cannot open %s", filepath2))
	}
	X := splitLines(string(content1))
	Y := splitLines(string(content2))
	diffTool.Diff(X, Y)
}

func splitLines(x string) []string {
	return strings.Split(x, "\n")
}
