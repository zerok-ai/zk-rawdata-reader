package pxl

import (
	"embed"
	"fmt"
)

//go:embed *.pxl
var PxlFileHandler embed.FS

func GetPxlFileContent(fileName string) (string, error) {
	data, err := PxlFileHandler.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error while reading pxl file", err)
		return "", err
	}
	return string(data), nil
}
