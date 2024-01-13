package diff

import (
	"fmt"
	"testing"
)

func TestLCSSingleString(t *testing.T) {
	string1 := "AABCXY"
	string2 := "XYZ"
	want := "XY"
	diffTool := NewDiffTool()
	fmt.Println(string1, string2)
	got := diffTool.LCS(string1, string2)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func TestLCSSingleString2(t *testing.T) {
	string1 := "AGCAT"
	string2 := "GAC"
	want := "GC"
	diffTool := NewDiffTool()
	fmt.Println(string1, string2)
	got := diffTool.LCS(string1, string2)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
