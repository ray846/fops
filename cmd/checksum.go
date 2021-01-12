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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"

	"github.com/ray846/fops/fileinfo"
	"github.com/spf13/cobra"
)

var algorithms = []struct {
	name string
	hash hash.Hash
}{
	{"md5", md5.New()},
	{"sha1", sha1.New()},
	{"sha256", sha256.New()},
}

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:     "checksum",
	Short:   "Print the checksum of file",
	Long:    `Print the checksum of file, support multiple algorithms: md5, sha1 and sha256`,
	Example: "  checksum -f [file] --[algorithm]",
	RunE: func(cmd *cobra.Command, args []string) error {
		algo, err := checkAlgorithm(cmd)
		if err != nil {
			return err
		}

		valid, err := fileinfo.IsValidFile(inputFile, true)
		if err != nil {
			return err
		} else if !valid {
			return err
		}

		sum, err := checksum(inputFile, algo)
		if err != nil {
			return err
		}
		fmt.Print(sum)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&inputFile, "file", "f", "", "input file")
	checksumCmd.MarkFlagRequired("file")
	for _, i := range algorithms {
		checksumCmd.Flags().Bool(i.name, false, fmt.Sprintf("checksum type %s", i.name))
	}
}

func checkAlgorithm(cmd *cobra.Command) (hash.Hash, error) {
	count := 0
	var theChosenOne hash.Hash

	//check which one has been chosen
	for _, i := range algorithms {
		choose, err := cmd.Flags().GetBool(i.name)
		if err != nil {
			return nil, nil
		} else if choose {
			count++
			theChosenOne = i.hash
		}
	}

	if count == 1 {
		return theChosenOne, nil
	}
	return nil, errors.New("you can choose 1 and only 1 algorithm")
}

func checksum(file string, algoHash hash.Hash) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", nil
	}
	algoHash.Write(content)
	return fmt.Sprintf("%x", algoHash.Sum(nil)), nil
}
