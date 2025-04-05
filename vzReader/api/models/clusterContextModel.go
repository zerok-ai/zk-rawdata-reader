package models

type ClusterContext struct {
	ClusterKey string `json:"apiKey"`
	CloudAddr  string `json:"cloudAddr"`
	ClusterId  string `json:"clusterId"`
}
