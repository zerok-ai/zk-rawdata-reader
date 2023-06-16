package vzReader

// Protocol is an enum for the different protocols that can be used to get raw data
type Protocol string

const (
	HTTP  Protocol = "HTTP"
	MySQL Protocol = "MySQL"
)

// mapping for pxl files location relative to package root.
var filePathPrefix = "/vzReader/pxl"

var protocolMapping = map[Protocol]string{
	HTTP: filePathPrefix + "/getHTTPRawDataForTraces.pxl",
}
