package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	"os"
)

func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] directory...\n", os.Args[0])
	fmt.Println("            (multiple directories may be specified)")
	flag.PrintDefaults()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processImages(imagePaths []string) []ImageMetadata {
	var output []ImageMetadata
	for _, imagePath := range imagePaths {
		metadata, err := calculateImageMetadata(imagePath)
		check(err)
		output = append(output, metadata)
	}
	return output
}

func main() {
	recursive := flag.Bool("r", false, "Recursively scan directories")
	deleteFiles := flag.Bool("clean", false, "Delete duplicate images (otherwise only find duplicates)")
	flag.Usage = myUsage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filesToProcess := findImagesInPaths(flag.Args(), *recursive)
	fmt.Printf("Found %d images\n", len(filesToProcess))

	imageMetadata := processImages(filesToProcess)

	byHash := map[float64][]ImageMetadata{}

	for _, img := range imageMetadata {
		if val, ok := byHash[img.hashSum]; ok {
			byHash[img.hashSum] = append(val, img)
		} else {
			byHash[img.hashSum] = []ImageMetadata{img}
		}
	}

	for i := range byHash {
		img := byHash[i]
		fmt.Println(i, img)
	}

	if *deleteFiles {
		fmt.Println("I would delete files")
	} else {
		fmt.Println("I would not delete files")
	}
}
