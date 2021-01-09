package cmd

import "testing"

func TestLineCount(t *testing.T) {
	file := "../test/myfile.txt"
	ans, _ := lineCount(file)
	if ans != 4 {
		t.Errorf("lineCount(\"%s\") = %d; want 4", file, ans)
	}
}
