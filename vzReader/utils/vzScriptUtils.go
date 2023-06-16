package utils

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"px.dev/pixie/src/api/go/pxapi"
	"px.dev/pixie/src/api/go/pxapi/errdefs"
)

type TemplateValues struct {
	TraceIds  string
	StartTime string
}

func ResolveFileDataAsTemplate(filePath string, tx TemplateValues) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		println("failed to get working dir")
		return "", nil
	}

	dat, err := os.ReadFile(path + filePath)
	if err != nil {
		fmt.Printf("failed to get script file content %s\n", path+filePath)
		return "", nil
	}
	templateStr := template.New("Template")
	templateStr, _ = templateStr.Parse(string(dat))

	var doc bytes.Buffer
	err = templateStr.Execute(&doc, tx)
	if err != nil {
		println("failed to Parse script template")
		return "", nil
	}
	scriptStr := doc.String()
	return scriptStr, nil
}

func GetResult(resultSet *pxapi.ScriptResults) (*pxapi.ScriptResults, error) {
	// Receive the PxL script results.
	defer func(resultSet *pxapi.ScriptResults) {
		err := resultSet.Close()
		if err != nil {
			println("Error while closing resultSet, err: %v\n", err)
		}
	}(resultSet)

	if err := resultSet.Stream(); err != nil {
		if errdefs.IsCompilationError(err) {
			fmt.Printf("Got compiler error: \n %v\n", err)
			return nil, err
		} else {
			fmt.Printf("Got error : %+v, while streaming\n", err)
		}
		if err.Error() == "rpc error: code = Internal desc = Auth middleware failed: failed to fetch token - unauthenticated" {
			return nil, err
		}
		return nil, err
	}

	return resultSet, nil
}
