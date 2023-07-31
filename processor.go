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
	metadata  filesMetadata
}

func (p *processor) renameFiles(pattern string) {
	filesProcessed := 0
	for _, f := range p.fileNames {
		fileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Println(fmt.Sprintf("error getting file info for file: %s", f), err)
			return
		}
		modifiedTime := fileInfo.ModTime()
		formattedModTime := modifiedTime.Format(timeFmt)

		randomStr := randomString(10)

		err = renameFileWithExtension(f, fmt.Sprintf("%s-%s-%s", formattedModTime, pattern, randomStr))
		if err != nil {
			fmt.Println("error renaming file:", err)
			return
		}
		filesProcessed++
	}
	p.metadata.filesProcessedCount = filesProcessed
}

func (p *processor) printSummary() {
	fmt.Println()
	fmt.Printf("Number of files found: %d\n", p.metadata.filesFoundCount)
	fmt.Printf("Number of files renamed: %d\n", p.metadata.filesProcessedCount)
	fmt.Printf("Number of files skipped: %d\n", p.metadata.filesSkippedCount)
}
