package models

type HttpRawDataModel struct {
	ReqHeaders   interface{} `json:"req_headers"`
	ReqBody      interface{} `json:"req_body"`
	RespHeaders  interface{} `json:"resp_headers"`
	RespBody     interface{} `json:"resp_body"`
	ReqBodySize  uint64      `json:"req_body_size"`
	RespBodySize uint64      `json:"resp_body_size"`
	WorkloadIds  string      `json:"workload_ids"`
	TraceId      string      `json:"trace_id"`
	SpanId       string      `json:"span_id"`
}
