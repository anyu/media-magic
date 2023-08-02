package main

import (
	"fmt"
	"os"
)

const timeFmt = "2006-01-02"

type filesMetadata struct {
	filesFoundCount     int
	filesProcessedCount int
	filesSkippedCount   int
}

type processor struct {
	fileNames []string
	outputDir string
	metadata  filesMetadata
}

func (p *processor) renameAndMoveFiles(label string) {
	filesProcessed := 0
	filesSkipped := 0

	fmt.Println()
	fmt.Println("Renaming files and moving to output directory:", p.outputDir)
	for _, f := range p.fileNames {
		// Get file info for current file
		fileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Println(fmt.Sprintf("error getting file info for file: %s", f), err)
			filesSkipped++
			continue
		}

		// Generate new file name by combining the file modified time with label and random string.
		formattedModTime := fileInfo.ModTime().Format(timeFmt)
		newName := formattedModTime + "_" + label + "-" + randomString(10)

		err = renameFileWithExtension(f, p.outputDir, newName)
		if err != nil {
			fmt.Println("error renaming file:", err)
			filesSkipped++
			continue
		}
		filesProcessed++
	}
	p.metadata.filesProcessedCount = filesProcessed
	p.metadata.filesSkippedCount = filesSkipped
}

func (p *processor) printSummary() {
	fmt.Println()
	fmt.Printf("Number of files found: %d\n", p.metadata.filesFoundCount)
	fmt.Printf("Number of files renamed: %d\n", p.metadata.filesProcessedCount)
	fmt.Printf("Number of files skipped: %d\n", p.metadata.filesSkippedCount)
}
