package main

import (
	"testing"
)

func TestSplitWhiteSpaces(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []string
	}{
		{
			name:   "Empty string",
			input:  "",
			output: []string{},
		},
		{
			name:   "Single word",
			input:  "hello",
			output: []string{"hello"},
		},
		{
			name:   "Multiple words with spaces",
			input:  "This is a test string",
			output: []string{"This", "is", "a", "test", "string"},
		},
		{
			name:   "Trailing spaces",
			input:  "This has trailing spaces  ",
			output: []string{"This", "has", "trailing", "spaces"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := SplitWhiteSpaces(tc.input)
			if len(output) != len(tc.output) {
				t.Errorf("Expected %d words, got %d", len(tc.output), len(output))
				return
			}
			for i := range output {
				if output[i] != tc.output[i] {
					t.Errorf("Expected word %d to be %s, got %s", i+1, tc.output[i], output[i])
					return
				}
			}
		})
	}
}

func TestPuncTuation(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "No punctuation",
			input:  []string{"This", "is", "a", "test", "string"},
			output: []string{"This", "is", "a", "test", "string"},
		},
		{
			name:   "Punctuation at the end",
			input:  []string{"This", "is", "a", "test", "string", "."},
			output: []string{"This", "is", "a", "test", "string."},
		},
		{
			name:   "Punctuation requiring change an to a",
			input:  []string{"This", "is", "an", "pest", "string"},
			output: []string{"This", "is", "a", "pest", "string"},
		},
		{
			name:   "Punctuation requiring a change (a)",
			input:  []string{"A", "apple", "a", "day"},
			output: []string{"An", "apple", "a", "day"},
		},
		{
			name:   "Punctuation requiring no change (the)",
			input:  []string{"The", "apple", "a", "day"},
			output: []string{"The", "apple", "a", "day"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := puncTuation(tc.input)
			if len(output) != len(tc.output) {
				t.Errorf("Expected %d words, got %d", len(tc.output), len(output))
				return
			}
			for i := range output {
				if output[i] != tc.output[i] {
					t.Errorf("Expected word %d to be %s, got %s", i+1, tc.output[i], output[i])
					return
				}
			}
		})
	}
}

func TestCheckQuotation(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "No qunctuation",
			input:  []string{"This", "is", "a", "test", "string"},
			output: []string{"This", "is", "a", "test", "string"},
		},
		{
			name:   "Punctuation requiring change ",
			input:  []string{"This", "is", "an", "pest", "'", "string.", "'"},
			output: []string{"This", "is", "a", "pest", "'", "string.", "'"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := puncTuation(tc.input)
			if len(output) != len(tc.output) {
				t.Errorf("Expected %d words, got %d", len(tc.output), len(output))
				return
			}
			for i := range output {
				if output[i] != tc.output[i] {
					t.Errorf("Expected word %d to be %s, got %s", i+1, tc.output[i], output[i])
					return
				}
			}
		})
	}
}

func TestfinalPunc(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "No qunctuation",
			input:  []string{"This", "is", "a", "test", "string"},
			output: []string{"This", "is", "a", "test", "string"},
		},
		{
			name:   "Punctuation requiring change ",
			input:  []string{"This", "is", "an", "pest", "'", "string.", "'"},
			output: []string{"This", "is", "a", "pest", "'string.'"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := puncTuation(tc.input)
			if len(output) != len(tc.output) {
				t.Errorf("Expected %d words, got %d", len(tc.output), len(output))
				return
			}
			for i := range output {
				if output[i] != tc.output[i] {
					t.Errorf("Expected word %d to be %s, got %s", i+1, tc.output[i], output[i])
					return
				}
			}
		})
	}
}

func TestHexConv2Dec(t *testing.T) {
	input := "FF"
	expected := "255"
	input = "1E"
	expected = "30"
	actual := HexConv2Dec(input)
	if actual != expected {
		t.Errorf("HexConv2Dec failed on valid hex test. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestBinConv2Dec(t *testing.T) {
	input := "10"
	expected := "2"
	input = "101"
	expected = "5"
	actual := BinConv2Dec(input)
	if actual != expected {
		t.Errorf("HexConv2Dec failed on valid hex test. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestToUpper(t *testing.T) {
	input := "function"
	expected := "FUNCTION"
	input = "Alphabet"
	expected = "ALPHABET"
	actual := ToUpper(input)
	if actual != expected {
		t.Errorf("HexConv2Dec failed on valid hex test. Expected: %v, Actual: %v", expected, actual)
	}
}

func TestToLower(t *testing.T) {
	input := "FunctioN"
	expected := "function"
	input = "AlphAbet"
	expected = "alphabet"
	actual := ToLower(input)
	if actual != expected {
		t.Errorf("HexConv2Dec failed on valid hex test. Expected: %v, Actual: %v", expected, actual)
	}
}
