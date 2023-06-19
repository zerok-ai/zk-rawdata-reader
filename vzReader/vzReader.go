package vzReader

import (
	"fmt"
	"strings"

	"context"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/utils"
	"px.dev/pixie/src/api/go/pxapi"
)

// VzReader is a struct that contains the configuration for the Vizier client
type VzReader struct {
	CloudAddr   string // px.avinpx08.getanton.com:443
	DirectVzId  string // 42f54f46-6e0f-46d7-bdab-6f5f66e0c7ff
	DirectVzKey string // px-api-ef2df548-1b59-44be-8e48-ab098809694d

	vzClient *pxapi.VizierClient
	ctx      context.Context
}

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
	client, err := pxapi.NewClient(config.ctx, pxapi.WithAPIKey(config.DirectVzKey),
		pxapi.WithCloudAddr(config.CloudAddr))
	if err != nil {
		fmt.Printf("Failed to connect to Cloud, err: %v\n", err)
		return nil, err
	}
	return client.NewVizierClient(config.ctx, config.DirectVzId)
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

// Init initializes the VzReader struct
func (config *VzReader) Init() error {

	// TODO: Populate CloudAddr, DirectVzId, DirectVzKey from Operator API if not present.

	// Create context and VizierClient for the given config
	config.ctx = context.Background()
	vizierClient, err := getVzClient(config)
	if err != nil {
		fmt.Printf("Failed to connect to Vizier, err: %v\n", err)
		return err
	}
	config.vzClient = vizierClient
	return nil
}
