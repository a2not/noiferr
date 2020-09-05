package noiferr_test

import (
	"testing"

	"github.com/Khdbble/noiferr"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, noiferr.Analyzer, "a")
}

