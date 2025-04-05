package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/api/models"
)

// The mock store contains additonal methods for inspection
type OperatorAPIHandlerMock struct {
	mock.Mock
}

func (m *OperatorAPIHandlerMock) GetClusterContext() (*models.ClusterContext, error) {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called()
	return rets.Get(0).(*models.ClusterContext), rets.Error(1)
}

func New() *OperatorAPIHandlerMock {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new OperatorAPIHandlerMock instance to it, instead of an actual store
	*/
	s := new(OperatorAPIHandlerMock)
	return s
}
