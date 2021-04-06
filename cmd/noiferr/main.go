package main

import (
	"github.com/a2not/noiferr"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(noiferr.Analyzer) }
