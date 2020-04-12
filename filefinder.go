package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func isImage(filename string) bool {
	lower := strings.ToLower(filename)

	return strings.HasSuffix(lower, "png") ||
		strings.HasSuffix(lower, "jpg") ||
		strings.HasSuffix(lower, "gif")
}

func findFilesRecursive(root string) []string {
	var output []string
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if isImage(info.Name()) {
				output = append(output, path)
			}
			return nil
		})
	check(err)
	return output
}

func findFiles(root string) []string {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	var output []string
	for _, f := range files {
		if isImage(f.Name()) {
			output = append(output, filepath.Clean(root+"/"+f.Name()))
		}
	}
	return output
}

// Returns slice containing unique paths to all the images found in each given path
func findImagesInPaths(paths []string, recursive bool) []string {
	filesToProcess := map[string]bool{}
	for _, path := range paths {
		fmt.Println("Finding images in", path)
		if recursive {
			for _, imagePath := range findFilesRecursive(path) {
				filesToProcess[imagePath] = true
			}
		} else {
			for _, imagePath := range findFiles(path) {
				filesToProcess[imagePath] = true
			}
		}
	}
	output := make([]string, 0, len(filesToProcess))
	for imagePath := range filesToProcess {
		output = append(output, imagePath)
	}

	return output
}
