package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	imageFormats = []string{".jpg", ".jpeg", ".png", ".heic"}
	videoFormats = []string{".mov", ".mp4"}
	mediaFormats = append(imageFormats, videoFormats...)
)

const DefaultOutputDirName = "renamed"

func main() {

	mediaDir := flag.String("file", "", "Path to directory with unorganized media files")
	flag.Parse()

	outputDir := flag.String("output", "", "Path to output directory for renamed meda files")
	flag.Parse()

	if *mediaDir == "" {
		fmt.Println("File path is missing")
		os.Exit(1)
	}

	mediaFiles, err := findMediaFiles(*mediaDir, mediaFormats)
	if err != nil {
		fmt.Println("error finding media files", err)
		os.Exit(1)
	}

	fmt.Println(mediaFiles)

	for _, f := range mediaFiles {
		fileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Println(fmt.Sprintf("error getting file info for file: %s", f), err)
			return
		}
		modifiedTime := fileInfo.ModTime()
		formattedModTime := modifiedTime.Format("2006-01-02")

		fmt.Println("File modified time:", formattedModTime)

		randomStr := randomString(10)

		err = renameFileWithExtension(f, fmt.Sprintf("%s-test-%s", formattedModTime, randomStr))
		if err != nil {
			fmt.Println("error renaming file:", err)
			return
		}

	}

	outputDirName := DefaultOutputDirName
	if *outputDir != "" {
		outputDirName = *outputDir
	}
	err = createDirIfNotExist(outputDirName)
	if err != nil {
		return
	}
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

func hasExtension(fileName string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(fileName), ext) {
			return true
		}
	}
	return false
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func createDirIfNotExist(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
		fmt.Println("Directory created successfully.")
	} else if err != nil {
		return err
	}

	return nil
}

func renameFileWithExtension(oldFilePath, newFileName string) error {
	extension := filepath.Ext(oldFilePath)
	newFilePath := filepath.Join(filepath.Dir(oldFilePath), newFileName+extension)

	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}
