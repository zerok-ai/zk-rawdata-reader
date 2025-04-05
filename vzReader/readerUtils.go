package vzReader

import (
	"context"
	"fmt"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/utils"
	"px.dev/pixie/src/api/go/pxapi"
	"strings"
)

// getScriptStr returns the pxl script as string for the given protocol after applying the template
func getScriptStr(scriptFileName string, traceIds []string, startTime string) (string, error) {
	traceStrList := strings.Join(traceIds, "\", \"")
	templateValues := utils.TemplateValues{}
	templateValues.TraceIds = traceStrList
	templateValues.StartTime = startTime
	return utils.ResolveFileDataAsTemplate(scriptFileName, templateValues)
}

// getVzClient returns a new Vizier client
func getVzClient(config *VzReader) (*pxapi.VizierClient, error) {
	client, err := pxapi.NewClient(config.ctx, pxapi.WithAPIKey(config.ClusterKey),
		pxapi.WithCloudAddr(config.CloudAddr))
	if err != nil {
		fmt.Printf("Failed to connect to Cloud, err: %v\n", err)
		return nil, err
	}
	return client.NewVizierClient(config.ctx, config.ClusterId)
}

// executeScript executes the given script and returns the result
func executeScript(vizierClient *pxapi.VizierClient, ctx context.Context, scriptStr string, tm pxapi.TableMuxer) (*pxapi.ScriptResults, error) {
	resultSet, err := vizierClient.ExecuteScript(ctx, scriptStr, tm)
	if err != nil {
		fmt.Printf("Error while executing query, %v\n", err)
		return nil, err
	}

	result, err := utils.GetResult(resultSet)
	if err != nil {
		fmt.Printf("Error while getting results, %v\n", err)
		return nil, err
	}

	return result, nil
}
