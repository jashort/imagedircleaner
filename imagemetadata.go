package main

import (
	"fmt"
	"github.com/vitali-fedulov/images"
)

type ImageMetadata struct {
	path    string
	width   int
	height  int
	hash    []float32
	hashSum float64
}

func (i ImageMetadata) String() string {
	return fmt.Sprintf("%s: %f", i.path, i.hashSum)
}

func calculateImageMetadata(filename string) (ImageMetadata, error) {
	imgA, err := images.Open(filename)
	if err != nil {
		return ImageMetadata{}, err
	}

	imgHash, imgSize := images.Hash(imgA)
	var metadata = ImageMetadata{path: filename, width: imgSize.X, height: imgSize.Y, hash: imgHash}
	imgHashSum := float64(0)
	for i := 0; i < len(imgHash); i++ {
		imgHashSum = imgHashSum + float64(imgHash[i])
	}
	metadata.hashSum = imgHashSum
	return metadata, nil
}
