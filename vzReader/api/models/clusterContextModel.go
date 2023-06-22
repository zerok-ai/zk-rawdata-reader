package models

type ClusterContext struct {
	ClusterKey string `json:"clusterKey"`
	CloudAddr  string `json:"cloudAddr"`
	ClusterId  string `json:"clusterId"`
}
