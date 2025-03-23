package fileapp

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CreateFile(name string, dirSave string, file multipart.File, header *multipart.FileHeader) (path string, ext string, err error) {
	fileExtension := filepath.Ext(header.Filename)

	outputFileName := fmt.Sprintf("%s/%s%s", dirSave, name, fileExtension)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return "", "", err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		return "", "", err
	}

	return outputFileName, fileExtension, nil
}
