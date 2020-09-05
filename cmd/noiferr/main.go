package main

import (
	"github.com/Khdbble/noiferr"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noiferr.Analyzer) }

