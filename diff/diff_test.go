package diff

import (
	"testing"
)

func TestLCSSingleString(t *testing.T) {
	type tests struct {
		name           string
		X              string
		Y              string
		expectedOutput string
	}

	ts := []tests{
		{
			name:           "LCS(AABCXY, XYZ)",
			X:              "AABCXY",
			Y:              "XYZ",
			expectedOutput: "XY",
		},
		{
			name:           "LCS(ABCDEF, ABCDEF)",
			X:              "ABCDEF",
			Y:              "ABCDEF",
			expectedOutput: "ABCDEF",
		},
		{
			name:           "LCS(AABCXY, XYZ)",
			X:              "AABCXY",
			Y:              "XYZ",
			expectedOutput: "XY",
		},
		{
			name:           "LCS('', '')",
			X:              "",
			Y:              "",
			expectedOutput: "",
		},
		{
			name:           "LCS(ABCD, AC)",
			X:              "ABCD",
			Y:              "AC",
			expectedOutput: "AC",
		},
	}

	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			diffTool := NewDiffTool()
			got := diffTool.LCS(tt.X, tt.Y)
			if got != tt.expectedOutput {
				t.Errorf("got %q, want %q", got, tt.expectedOutput)
			}
		})
	}
}
