## zk-rawdata-reader
Go Package to read raw data from vizier.

### Exposed APIs
```go   
import "github.com/zerok-ai/zk-rawdata-reader/vzReader"

// Vizier client configuration
vzReader.VzReader{
	// Optional, if not provided, will be read from Operator API.
    CloudAddr string
    DirectVzId string
    DirectVzKey string
}

// Protocol supported for raw data fetch 
Protocol 
    HTTP

// Init initializes the reader
func (r *VzReader) Init() error

// GetHTTPRawData returns raw data for given traceIds and startTime
func (r *VzReader) GetHTTPRawData(traceIds []string, startTime string) ([]byte, error)
```

### Usage
```go
import (
    "github.com/zerok-ai/zk-rawdata-reader/vzReader"
)

func main() {
	reader := vzReader.VzReader{}
	reader.CloudAddr = "<Cloud-address>>"
	reader.DirectVzId = "<Cluster-ID>"
	reader.DirectVzKey = "<Cluster-key>"

    _ := reader.Init()

	startTime := "<Duration>" // -5m, -10m, -1h etc
	traceIds := []string{"<trace-id-1>", "<trace-id-2>"}
	data, _ := reader.GetHTTPRawData(traceIds, startTime)

}
```
