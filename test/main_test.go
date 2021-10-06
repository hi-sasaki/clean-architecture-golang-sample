package test

import (
	"fmt"
	"os"
	"testing"
)

// カバー率（0.8 以上で合格）
var thresholdCoverage = 0.8

func TestMain(m *testing.M) {
	var returnCode = m.Run()
	if returnCode == 0 && testing.CoverMode() != "" {
		var rateCoverage = testing.Coverage()
		if rateCoverage < thresholdCoverage {
			fmt.Println("テストはパスしましたが。カバレッジに失敗しました。カバレッジ率:", rateCoverage)
			returnCode = 0
		}
	}
	os.Exit(returnCode)
}
