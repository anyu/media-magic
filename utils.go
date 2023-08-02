package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// hasExtension checks if the provided file ends with particular file extensions.
func hasExtension(fileName string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(fileName), ext) {
			return true
		}
	}
	return false
}

// createOutputDirIfNotExists creates the output directory at the provided path if it doesn't already exist.
func createOutputDirIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
		fmt.Println("Output directory created successfully.")
	} else if err != nil {
		return err
	}

	return nil
}

// renameFileWithExtension renames the provided file to the output directory with the provided file name.
func renameFileWithExtension(oldFilePath, outputDir, newFileName string) error {
	extension := filepath.Ext(oldFilePath)
	newFilePath := filepath.Join(outputDir, newFileName+extension)

	oldFile := filepath.Base(oldFilePath)
	newFile := filepath.Base(newFilePath)

	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("- Renamed '%s' to '%s'\n", oldFile, newFile)

	return nil
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}
