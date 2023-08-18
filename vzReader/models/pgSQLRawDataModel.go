package models

type PgSQLRawDataModel struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	RemotePort  int     `json:"remote_port"`
	ReqCmd      string  `json:"req_cmd"`
	Req         string  `json:"req"`
	Resp        string  `json:"resp"`
	Latency     float32 `json:"latency"`
	TraceId     string  `json:"trace_id"`
	SpanId      string  `json:"span_id"`
	WorkloadIds string  `json:"workload_ids"`
	IsTruncated bool    `json:"is_truncated"`
	Time        uint64  `json:"time"`
}
