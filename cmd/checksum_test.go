package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"testing"
)

func TestChecksum(t *testing.T) {
	var tests = []struct {
		testName string
		file     string
		hash     hash.Hash
		want     string
	}{
		{"md5", "../test/myfile.txt", md5.New(), "a8c5d553ed101646036a811772ffbdd8"},
		{"sha1", "../test/myfile.txt", sha1.New(), "a656582ca3143a5f48718f4a15e7df018d286521"},
		{"sha256", "../test/myfile.txt", sha256.New(), "495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans, _ := checksum(tt.file, tt.hash)
			if ans != tt.want {
				t.Errorf("\ngot  %s\nwant %s", ans, tt.want)
			}
		})
	}
}
