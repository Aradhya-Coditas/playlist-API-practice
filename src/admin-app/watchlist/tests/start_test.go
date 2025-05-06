package tests

import (
	"admin-app/watchlist/commons/constants"
	"context"
	setupTest "omnenest-backend/src/setupTest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type watchlist struct {
	suite.Suite
	ctrl *gomock.Controller
	ctx  context.Context
}

func (suite *watchlist) SetupSuite() {
	setupTest.InitSuiteConfigs(constants.ServiceName)
}

func (suite *watchlist) BeforeTest(suiteName, testName string) {
	suite.ctx = context.Background()
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *watchlist) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestWatchlist(t *testing.T) {
	suite.Run(t, new(watchlist))
}
