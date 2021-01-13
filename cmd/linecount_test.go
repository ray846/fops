package cmd

import "testing"

func TestLineCount(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"../test/myfile.txt", 4},
		{"../test/test.js", 5},
		{"../test/test.json", 4},
		{"../test/test.csv", 4},
		{"../test/test.html", 11},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			ans, _ := lineCount(tt.file)
			if ans != tt.want {
				t.Errorf("lineCount(\"%s\") = %d; want %d", tt.file, ans, tt.want)
			}
		})
	}
}
