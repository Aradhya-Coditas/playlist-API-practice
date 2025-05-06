package tests

import (
	"context"
	setupTest "omnenest-backend/src/setupTest"
	"search/commons/constants"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type search struct {
	suite.Suite
	ctrl *gomock.Controller
	ctx  context.Context
}

func (suite *search) SetupSuite() {
	setupTest.InitSuiteConfigs(constants.ServiceName)
}

func (suite *search) BeforeTest(suiteName, testName string) {
	suite.ctx = context.Background()
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *search) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestSearch(t *testing.T) {
	suite.Run(t, new(search))
}
