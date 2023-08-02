package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	// Default source directory to look for media files.
	defaultSourceDir = "Downloads"
	// Default output directory for renamed and moved files.
	defaultOutputDir = "renamed"
	// Default label to use in renamed files.
	defaultLabel = "pic"
)

var (
	imageFormats = []string{".jpg", ".jpeg", ".png", ".heic"}
	videoFormats = []string{".mov", ".mp4"}
	mediaFormats = append(imageFormats, videoFormats...)
)

func main() {
	sourceDir, outputDir, label, err := parseFlagsWithDefaults()
	if err != nil {
		log.Fatal("error parsing flags", err)
	}

	mediaFiles, err := findMediaFiles(sourceDir, mediaFormats)
	if err != nil {
		log.Fatal("error finding media files", err)
	}

	err = createOutputDirIfNotExists(outputDir)
	if err != nil {
		log.Fatal("error creating output dir", err)
	}

	p := processor{
		fileNames: mediaFiles,
		outputDir: outputDir,
		metadata: filesMetadata{
			filesFoundCount: len(mediaFiles),
		},
	}
	p.renameAndMoveFiles(label)
	p.printSummary()
}

func parseFlagsWithDefaults() (string, string, string, error) {
	sourceDir := flag.String("source", "", "Path to source directory with unorganized media files")
	outputDir := flag.String("output", "", "Path to output directory for renamed media files")
	label := flag.String("label", "", "Label for renamed files")

	flag.StringVar(sourceDir, "s", "", "Shorthand for source directory")
	flag.StringVar(outputDir, "o", "", "Shorthand for output directory")
	flag.StringVar(label, "l", defaultLabel, "Shorthand for label")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ./media-magic [options]\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	sourcePath := *sourceDir
	if sourcePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("error getting user's home directory:", err)
			return "", "", "", err
		}
		sourcePath = filepath.Join(homeDir, defaultSourceDir)
		fmt.Println("No source directory provided, using default:", sourcePath)
	}

	outputPath := *outputDir
	if outputPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("error getting user's home directory:", err)
			return "", "", "", err
		}
		currentDate := time.Now().Format(timeFmt)
		outputPath = filepath.Join(homeDir, "Desktop", currentDate+"_"+defaultOutputDir)
		fmt.Println("No output directory provided, using default:", outputPath)
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
