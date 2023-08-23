package models

type HttpRawDataModel struct {
	Source       string      `json:"source"`
	Destination  string      `json:"destination"`
	SourceIp     string      `json:"source_ip"`
	DestIp       string      `json:"dest_ip"`
	Latency      float32     `json:"latency"`
	ReqPath      string      `json:"req_path"`
	ReqMethod    string      `json:"req_method"`
	ReqHeaders   interface{} `json:"req_headers"`
	ReqBody      interface{} `json:"req_body"`
	RespStatus   int         `json:"resp_status"`
	RespMessage  string      `json:"resp_message"`
	RespHeaders  interface{} `json:"resp_headers"`
	RespBody     interface{} `json:"resp_body"`
	ReqBodySize  uint64      `json:"req_body_size"`
	RespBodySize uint64      `json:"resp_body_size"`
	TraceId      string      `json:"trace_id"`
	SpanId       string      `json:"span_id"`
	WorkloadIds  string      `json:"workload_ids"`
	IsTruncated  bool        `json:"is_truncated"`
	Time         uint64      `json:"time"`
}
