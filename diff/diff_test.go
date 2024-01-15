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

func TestLCSLines(t *testing.T) {
	type tests struct {
		name           string
		X              []string
		Y              []string
		expectedOutput []string
	}
	ts := []tests{
		{
			name: "Test 1",
			X: []string{
				"This is a test which contains:",
				"this is the lcs",
			},
			Y: []string{
				"this is the lcs",
				"we're testing",
			},
			expectedOutput: []string{"this is the lcs"},
		},
		{
			name: "Test 2",
			X: []string{
				"Coding Challenges helps you become a better software engineer through that build real applications.",
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
			Y: []string{
				"Helping you become a better software engineer through coding challenges that build real applications.",
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
			expectedOutput: []string{
				"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
				"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
			},
		},
	}
	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			diffTool := NewDiffTool()
			got := diffTool.LCSLine(tt.X, tt.Y)
			if !Equal(got, tt.expectedOutput) {
				t.Errorf("got %q, want %q", got, tt.expectedOutput)
			}
		})
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
