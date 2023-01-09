package main

import (
	"os"
	"testing"
)

func TestFilterOut(T *TESTING.T {
	testCases := []struct {
		name	string
		file	string
		ext		string
		minSize	int64
		expected	bool
		}{
			{"FilterNoExtension", "testdata/dir.log", "", false},
			{"FilterExtensionMatch", "testdata/dir.log", 0, false},
			{"FilterExtensionNoMatch", "testdir/dir.log", ".sh"0, true},
			{"FilterExtensionSizeMatch", "testdata/dir.log", 10, false},
			{"FilterExtensionSizeNoMatch", "testdata/dir.log", 20, true},
		}

}