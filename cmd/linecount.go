/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/ray846/fops/fileinfo"
	"github.com/spf13/cobra"
)

// linecountCmd represents the linecount command
var linecountCmd = &cobra.Command{
	Use:     "linecount",
	Short:   "Print the line count of file",
	Example: "  linecount -f [file]",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fileinfo.IsValidFile(inputFile, false)
		if err != nil {
			return err
		}
		count, err := lineCount(inputFile)
		if err != nil {
			return err
		}
		fmt.Print(count)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&inputFile, "file", "f", "", "input file")
	linecountCmd.MarkFlagRequired("file")
}

func lineCount(file string) (int, error) {
	fileContent, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	const lineBreak = '\n'
	count := 0 // don't count first line?
	buf := make([]byte, bufio.MaxScanTokenSize)
	for {
		bufferSize, err := fileContent.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	return count, nil
}
