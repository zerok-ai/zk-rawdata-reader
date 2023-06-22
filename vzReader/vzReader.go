package vzReader

import (
	"fmt"
	zk_operator "github.com/zerok-ai/zk-rawdata-reader/vzReader/api/zk-operator"
	"strings"

	"context"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/utils"
	"px.dev/pixie/src/api/go/pxapi"
)

// VzReader is a struct that contains the configuration for the Vizier client
type VzReader struct {
	CloudAddr   string
	DirectVzId  string
	DirectVzKey string

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

	// Populate CloudAddr, DirectVzId, DirectVzKey from Operator API if not present.
	if config.CloudAddr == "" || config.DirectVzId == "" || config.DirectVzKey == "" {
		fmt.Printf("CloudAddr, DirectVzId, DirectVzKey cannot be empty\n")
		clusterContext, err := zk_operator.GetClusterContext()
		if err != nil {
			return err
		}
		config.DirectVzId = clusterContext.ClusterId
		config.DirectVzKey = clusterContext.ClusterKey
		config.CloudAddr = clusterContext.CloudAddr
	}

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
