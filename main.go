package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	imageFormats = []string{".jpg", ".jpeg", ".png", ".heic"}
	videoFormats = []string{".mov", ".mp4"}
	mediaFormats = append(imageFormats, videoFormats...)
)

const (
	DefaultTargetDirName = "~/Downloads"
	DefaultOutputDirName = "renamed"
)

func main() {

	mediaDir := flag.String("source", "", "Path to source directory with unorganized media files")
	outputDir := flag.String("output", "", "Path to output directory for cmd media files")
	namingPattern := flag.String("naming-pattern", "", "Naming pattern for cmd files")

	flag.Parse()

	targetDir := DefaultTargetDirName
	if *mediaDir == "" {
		fmt.Println("File path is missing. Defaulting to looking in Downloads")
	} else {
		targetDir = *mediaDir
	}

	mediaFiles, err := findMediaFiles(targetDir, mediaFormats)
	if err != nil {
		fmt.Println("error finding media files", err)
		os.Exit(1)
	}

	err = setup(outputDir, err)
	if err != nil {
		fmt.Println("error with setup", err)
		os.Exit(1)
	}

	p := processor{
		fileNames: mediaFiles,
		metadata: filesMetadata{
			filesFoundCount: len(mediaFiles),
		},
	}
	p.renameFiles(*namingPattern)
	p.printSummary()
}

func setup(outputDir *string, err error) error {
	outputDirName := DefaultOutputDirName
	if *outputDir != "" {
		outputDirName = *outputDir
	}
	err = createDirIfNotExist(outputDirName)
	if err != nil {
		return err
	}
	return nil
}

func findMediaFiles(dirPath string, extensions []string) ([]string, error) {
	var matchedFiles []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && hasExtension(info.Name(), extensions) {
			matchedFiles = append(matchedFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matchedFiles, nil
}
