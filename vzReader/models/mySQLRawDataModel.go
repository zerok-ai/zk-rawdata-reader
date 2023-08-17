package models

type MySQLRawDataModel struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Latency     float32 `json:"latency"`
	RemotePort  int     `json:"remote_port"`
	ReqCmd      int     `json:"req_cmd"`
	ReqBody     string  `json:"req_body"`
	RespStatus  int     `json:"resp_status"`
	RespBody    string  `json:"resp_body"`
	TraceId     string  `json:"trace_id"`
	SpanId      string  `json:"span_id"`
	WorkloadIds string  `json:"workload_ids"`
	IsTruncated bool    `json:"is_truncated"`
	Rows        int     `json:"rows"`
	Time        uint64  `json:"time"`
}
