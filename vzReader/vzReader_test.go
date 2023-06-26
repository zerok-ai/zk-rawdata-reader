package vzReader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zerok-ai/zk-rawdata-reader/vzReader/api/models"
	operatorMockHandler "github.com/zerok-ai/zk-rawdata-reader/vzReader/api/zk-operator/mocks"

	"testing"
)

type VzReaderTestSuite struct {
	suite.Suite
	operatorMock *operatorMockHandler.OperatorAPIHandlerMock
	vzReader     VzReader
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, &VzReaderTestSuite{})
}

// runs before execution of suite
func (s *VzReaderTestSuite) SetupSuite() {
	p := operatorMockHandler.New()
	s.operatorMock = p
	s.vzReader = VzReader{}
}

func (s *VzReaderTestSuite) TestVzReader_Missing_ClusterContext_Pass() {

	clusterContext := models.ClusterContext{ClusterId: "94711f31-f693-46be-91c3-832c0f64b12f", ClusterKey: "px-api-a3c99b88-3e00-444d-8923-03b143f1cdbb", CloudAddr: "px.avinpx07.getanton.com:443"}
	s.operatorMock.On("GetClusterContext").Return(&clusterContext, nil)

	s.vzReader.ClusterId = clusterContext.ClusterId
	s.vzReader.ClusterKey = clusterContext.ClusterKey
	s.vzReader.CloudAddr = clusterContext.CloudAddr

	err := s.vzReader.Init()
	if err != nil {
		assert.Errorf(s.T(), err, "Init Failed")
	}
	data, err := s.vzReader.GetHTTPRawData([]string{"123"}, "-5m")
	if err != nil {
		assert.Errorf(s.T(), err, "GetHTTPRawData Failed")
	}
	fmt.Printf("data: %v\n", data)

	//resp, err := s.operatorMock.GetClusterContext()
	//assert.Nil(s.T(), err)
	//assert.EqualValues(s.T(), clusterContext, *resp)
}
