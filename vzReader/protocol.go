package vzReader

// Protocol is an enum for the different protocols that can be used to get raw data
type Protocol string

const (
	HTTP  Protocol = "HTTP"
	MySQL Protocol = "MySQL"
	PgSQL Protocol = "PgSQL"
)

var protocolMapping = map[Protocol]string{
	HTTP:  "getHTTPRawDataForTraces.pxl",
	MySQL: "getMySQLRawDataForTraces.pxl",
	PgSQL: "getPgSQLRawDataForTraces.pxl",
}
