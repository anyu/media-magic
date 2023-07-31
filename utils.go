package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func hasExtension(fileName string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(fileName), ext) {
			return true
		}
	}
	return false
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

	oldFile := filepath.Base(oldFilePath)
	newFile := filepath.Base(newFilePath)

	fmt.Printf("Renaming file '%s' to '%s'...\n", oldFile, newFile)
	err := os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}
