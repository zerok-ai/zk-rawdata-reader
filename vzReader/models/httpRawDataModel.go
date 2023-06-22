package models

type HttpRawDataModel struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Latency     string `json:"latency"`
	ReqPath     string `json:"req_path"`
	ReqMethod   string `json:"req_method"`
	ReqHeaders  string `json:"req_headers"`
	ReqBody     string `json:"req_body"`
	RespStatus  string `json:"resp_status"`
	RespMessage string `json:"resp_message"`
	RespHeaders string `json:"resp_headers"`
	RespBody    string `json:"resp_body"`
	TraceId     string `json:"trace_id"`
	SpanId      string `json:"span_id"`
}
