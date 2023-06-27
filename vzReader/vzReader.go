package vzReader

import (
	"context"
	"fmt"
	zk_operator "github.com/zerok-ai/zk-rawdata-reader/vzReader/api/zk-operator"
	"px.dev/pixie/src/api/go/pxapi"
)

// VzReader is a struct that contains the configuration for the Vizier client
type VzReader struct {
	CloudAddr  string
	ClusterId  string
	ClusterKey string

	vzClient *pxapi.VizierClient
	ctx      context.Context
}

// Init initializes the VzReader struct
func (config *VzReader) Init() error {
	// Populate CloudAddr, ClusterId, ClusterKey from Operator API if not present.
	if config.CloudAddr == "" || config.ClusterId == "" || config.ClusterKey == "" {
		fmt.Printf("CloudAddr, ClusterId, ClusterKey cannot be empty\n")
		clusterContext, err := zk_operator.GetClusterContext()
		if err != nil {
			return err
		}
		config.ClusterId = clusterContext.ClusterId
		config.ClusterKey = clusterContext.ClusterKey
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

func (config *VzReader) Close() {
	config.vzClient = nil
}
