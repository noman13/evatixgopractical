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
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var Copydir string
var Builddir string
var Exe string
// buildexecuteCmd represents the buildexecute command
var buildexecuteCmd = &cobra.Command{
	Use:   "buildexecute",
	Short: "Makes a copy of a folder.",
	Long: `gobuildercli makes copy of a directory according to the passed arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildexecute called")
		fmt.Println(Copydir)
		fmt.Println(Builddir)
		fmt.Println(Exe)
		//

		//copy("./config/aa.txt", "./dest")

		if Copydir != "" {
			source := Copydir
			currentDirectory:= "/"
			fmt.Println(currentDirectory)

			if Builddir != "" {
				destination := Builddir
				if Builddir != Copydir {
					copyDirectory(source, destination)
				} else {
					log.Println("No operation done! Same source and destination.")
				}
			} else {
				destination := currentDirectory
				if Builddir != Copydir {
					copyDirectory(source, destination)
				} else {
					log.Println("No operation done! Same source and current directory.")
				}
			}
		}
		if Exe != "" {

			cmd := exec.Command("go", "build", "-o", Exe)
			pwd, _ := os.Getwd()
			cmd.Dir = pwd
			_, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
		}
	},
}
func copyDirectory(source string , destination string) error {
	trimmedSourcePath, regexPattern := TrimREGEX(source)
	newDestination := verifyDestination(destination)
	fmt.Println(trimmedSourcePath, regexPattern)
	files, err := WalkMatch(trimmedSourcePath, regexPattern)
	if err != nil {
		return err
	}
	for _, filename := range files {
		newSource := trimmedSourcePath + filename
		fmt.Println(newSource)
		copy(newSource, newDestination, filename)
	}
	return err
}
func TrimREGEX(source string) (string, string){
	var trimmedSourcePath string
	var regexPattern string
	var flag int

	for i := len(source)-1; i >= 0; i-- {
		if source[i] == '/' {
			flag = i
			break
		}
	}
	for i:=0; i<len(source); i++{
		if i<= flag {
			trimmedSourcePath += string(source[i])
		} else {
			regexPattern += string(source[i])
		}
	}
	return trimmedSourcePath, regexPattern
}
func verifyDestination(dst string) string {
	if dst[len(dst)-1] != '/' {
		dst += "/"
	}
	return dst
}
func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
func copy(src, dst, filename string) error{
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst+filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func init() {
	rootCmd.AddCommand(buildexecuteCmd)

	buildexecuteCmd.PersistentFlags().StringVarP(&Copydir, "copydir", "c", "", "Passes the source directory.")
	buildexecuteCmd.PersistentFlags().StringVarP(&Builddir, "builddir", "b", "", "Passes the destination directory.")
	buildexecuteCmd.PersistentFlags().StringVarP(&Exe, "exe", "e", "", "Compile the code as binary with given filename.")
}
