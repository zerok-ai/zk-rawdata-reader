package models

type ResponseType interface {
	ClusterContext
}

type ResponseModel[T ResponseType] struct {
	Payload T `json:"payload"`
}
