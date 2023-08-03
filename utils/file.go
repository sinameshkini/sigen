package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func Mkdir(path, name string) (dirPath string, err error) {
	dirPath = fmt.Sprintf("%s/%s", path, name)

	exist, err := Exists(dirPath)
	if err != nil {
		return
	}

	if exist {
		logrus.Warnf(name, "already exist in", dirPath)
		return
	}

	if err = os.Mkdir(dirPath, 0775); err != nil {
		return
	}

	return
}

func Touch(path, name string) (filePath string, err error) {
	filePath = fmt.Sprintf("%s/%s", path, name)

	exist, err := Exists(filePath)
	if err != nil {
		return
	}

	if exist {
		logrus.Warnf(name, "already exist in", filePath)
		return
	}

	if _, err = os.Create(filePath); err != nil {
		return
	}

	return
}

func WriteToFile(filePath string, content string) (err error) {
	// open input file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	// write a chunk
	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return
}

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
