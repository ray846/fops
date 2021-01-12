package fileinfo

import (
	"testing"
)

func TestIsExist(t *testing.T) {
	var tests = []struct {
		testName string
		file     string
		want     bool
	}{
		{"simple exists file", "../test/myfile.txt", true},
		{"simple not exists file", "../test/myfile2.txt", false},
		{"simple exists file with wrong extension", "../test/myfile.jpg", false},
		{"simple exists path", "../test", true},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, _ := isExist(tt.file)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestIsDictionary(t *testing.T) {
	var tests = []struct {
		testName string
		file     string
		want     bool
	}{
		{"simple file", "../test/myfile.txt", false},
		{"simple path", "../test", true},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, _ := isDirectory(tt.file)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestIsBinaryFile(t *testing.T) {
	var tests = []struct {
		testName string
		file     string
		want     bool
	}{
		{"txt file", "../test/myfile.txt", false},
		{"js file", "../test/test.js", false},
		{"json file", "../test/test.json", false},
		{"csv file", "../test/test.csv", false},
		{"html file", "../test/test.html", false},
		{"jpg file", "../test/test.jpg", true},
		{"exe file", "../test/test.exe", true},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, _ := isBinaryFile(tt.file)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
