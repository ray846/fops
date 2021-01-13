package fileinfo

import (
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

//isExist checks if a file exisits.
func isExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false, nil
	}
	//check permission denied
	return err == nil, err
}

//isDirectory checks if the fileInfo is a directory
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

//isBinaryFile checks if file's mime type is not text/plain
func isBinaryFile(file string) (bool, error) {
	detectedMIME, err := mimetype.DetectFile(file)
	if err != nil {
		return false, err
	}
	for mime := detectedMIME; mime != nil; mime = mime.Parent() {
		if mime.Is("text/plain") {
			return false, nil
		}
	}
	return true, nil
}

//IsValidFile checks non-exist-file, got directory, binary file
func IsValidFile(file string, acceptBinary bool) (bool, error) {
	exist, err := isExist(file)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, fmt.Errorf("No such file '%s'", file)
	}

	isDir, err := isDirectory(file)
	if err != nil {
		return false, err
	}
	if isDir {
		return false, fmt.Errorf("Expected file got directory '%s'", file)
	}

	if !acceptBinary {
		isBinary, err := isBinaryFile(file)
		if err != nil {
			return false, err
		}
		if isBinary {
			return false, fmt.Errorf("Cannot do linecount for binary file '%s'", file)
		}
	}
	return err == nil, err
}
