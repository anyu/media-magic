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
	defaultSourceDir = "Downloads"
	defaultOutputDir = "renamed"
	defaultLabel     = "renamed"
)

func main() {

	sourceDir, outputDir, label, err := parseFlagsWithDefaults()
	if err != nil {
		fmt.Println("error parsing flags", err)
		os.Exit(1)
	}

	mediaFiles, err := findMediaFiles(sourceDir, mediaFormats)
	if err != nil {
		fmt.Println("error finding media files", err)
		os.Exit(1)
	}

	err = createDirIfNotExist(outputDir)
	if err != nil {
		fmt.Println("error creating dir", err)
		os.Exit(1)
	}

	p := processor{
		fileNames: mediaFiles,
		metadata: filesMetadata{
			filesFoundCount: len(mediaFiles),
		},
	}
	p.renameFiles(label)
	p.printSummary()
}

func parseFlagsWithDefaults() (string, string, string, error) {

	sourceDir := flag.String("source", "", "Path to source directory with unorganized media files")
	outputDir := flag.String("output", "", "Path to output directory for cmd media files")
	label := flag.String("label", defaultLabel, "Label for renamed files")

	flag.StringVar(sourceDir, "s", "", "Shorthand for source directory")
	flag.StringVar(outputDir, "o", "", "Shorthand for output directory")
	flag.StringVar(label, "l", "", "Shorthand for label")

	flag.Parse()

	sourcePath := *sourceDir

	if sourcePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user's home directory:", err)
			return "", "", "", err
		}
		sourcePath = filepath.Join(homeDir, defaultSourceDir)
		fmt.Printf("No source directory provided, using default: %s\n", sourcePath)
	}

	outputPath := *outputDir
	if outputPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user's home directory:", err)
			return "", "", "", err
		}
		outputPath = filepath.Join(homeDir, "Desktop", defaultOutputDir)
		fmt.Printf("No output directory provided, using default: %s\n", outputPath)
		fmt.Println()
	}

	return sourcePath, outputPath, *label, nil
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
