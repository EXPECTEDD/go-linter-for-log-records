package checkspecialcharactersoremoji

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckIsLower_Good(t *testing.T) {
	testDataGood, err := filepath.Abs("testdata/good")
	if err != nil {
		panic(err)
	}
	testDataBad, err := filepath.Abs("testdata/bad")
	if err != nil {
		panic(err)
	}

	analysistest.Run(t, testDataGood, Analyzer)
	analysistest.Run(t, testDataBad, Analyzer)
}
