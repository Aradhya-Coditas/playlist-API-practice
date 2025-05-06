package tests

import (
	"admin-app/search/business"
	"admin-app/search/commons/constants"
	"admin-app/search/handlers"
	"admin-app/search/models"
	"admin-app/search/repositories"
	"fmt"
	"net/http"
	"net/http/httptest"
	genericConstants "omnenest-backend/src/constants"
	setupTest "omnenest-backend/src/setupTest"
	"regexp"

	"github.com/bytedance/sonic"
)

func (suite *search) prepareSearchEquityGroupSetup(request interface{}, expectedResponseBytes []byte, expectedStatusCode int) {
	w := httptest.NewRecorder()
	ctx := setupTest.GetTestGinContext(w)
	setupTest.MockJsonPost(ctx, request)

	searchEquityGroupRepository := repositories.GetSearchEquityGroupRepository(false)
	searchEquityGroupService := business.NewSearchEquityGroupService(searchEquityGroupRepository)
	searchEquityGroupController := handlers.NewSearchEquityGroupController(searchEquityGroupService)

	searchEquityGroupController.HandleSearchEquityGroup(ctx)

	suite.EqualValues(expectedStatusCode, w.Code)
	suite.EqualValues(expectedResponseBytes, w.Body.Bytes())
	suite.EqualValues(string(expectedResponseBytes), w.Body.String())
}

func (suite *search) TestSearchEquityGroup400MismatchedDataType() {
	requestData := map[string]interface{}{
		"exchange": 12345,
	}
	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message:      []models.ErrorMessage{{Key: "exchange", ErrorMessage: genericConstants.JsonBindingFieldError}},
	}
	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareSearchEquityGroupSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

func (suite *search) TestSearchEquityGroup400MissingField() {
	requestData := models.BFFSearchEquityGroupRequest{}
	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: "exchange: required"},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareSearchEquityGroupSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

func (suite *search) TestSearchEquityGroup204NoDataFound() {
	// Define the SELECT DISTINCT query
	selectQuery := regexp.QuoteMeta(constants.GetDistinctGroupsByExchangeQuery)

	// Create empty rows to simulate no data found
	rows := setupTest.GetSqlMock().NewRows([]string{"group"})

	// Simulate the repository returning 0 rows
	setupTest.GetSqlMock().ExpectQuery(selectQuery).
		WithArgs("NSE").
		WillReturnRows(rows)

	// Request payload
	requestData := models.BFFSearchEquityGroupRequest{
		ExchangeName: "NSE",
	}

	// Invoke the test setup for the function being tested
	suite.prepareSearchEquityGroupSetup(requestData, nil, http.StatusNoContent)
}

func (suite *search) TestSearchEquityGroup204OnlyEmptyGroups() {
	// Define the SELECT DISTINCT query
	selectQuery := regexp.QuoteMeta(constants.GetDistinctGroupsByExchangeQuery)

	// Create rows with groups containing only whitespace or empty strings
	rows := setupTest.GetSqlMock().NewRows([]string{"group"}).
		AddRow(" ").
		AddRow("").
		AddRow("   "). // Multiple spaces
		AddRow(genericConstants.EmptySpace)

	// Simulate the repository returning these rows
	setupTest.GetSqlMock().ExpectQuery(selectQuery).WithArgs("NSE").WillReturnRows(rows)

	// Request payload
	requestData := models.BFFSearchEquityGroupRequest{
		ExchangeName: "NSE",
	}

	// Invoke the test setup for the function being tested
	suite.prepareSearchEquityGroupSetup(requestData, nil, http.StatusNoContent)
}

func (suite *search) TestSearchEquityGroup500InternalServerError() {
	// Mocking the SELECT DISTINCT query for the group column
	selectQuery := regexp.QuoteMeta(constants.GetDistinctGroupsByExchangeQuery)

	// Simulate database error
	setupTest.GetSqlMock().ExpectQuery(selectQuery).WithArgs("NSE").
		WillReturnError(fmt.Errorf("database error"))

	requestData := models.BFFSearchEquityGroupRequest{
		ExchangeName: "NSE",
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusInternalServerError),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.InternalServerError},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)

	// Call the function and assert the response
	suite.prepareSearchEquityGroupSetup(requestData, expectedResponseBytes, http.StatusInternalServerError)

	// Verify that all expectations were met
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *search) TestSearchEquityGroup200Success() {
	// Mocking the SELECT DISTINCT query
	selectQuery := regexp.QuoteMeta(constants.GetDistinctGroupsByExchangeQuery)

	// Create rows with sample data
	rows := setupTest.GetSqlMock().NewRows([]string{"group"}).
		AddRow("BL").
		AddRow("EQ").
		AddRow("XX")

	// Expect the query with specific arguments
	setupTest.GetSqlMock().ExpectQuery(selectQuery).
		WithArgs("NSE").
		WillReturnRows(rows)

	requestData := models.BFFSearchEquityGroupRequest{
		ExchangeName: "NSE",
	}

	expectedResponse := models.BFFSearchEquityGroupResponse{
		Groups: []string{"BL", "EQ", "XX"},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)

	// Call the function and assert the response
	suite.prepareSearchEquityGroupSetup(requestData, expectedResponseBytes, http.StatusOK)

	// Verify that all expectations were met
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}
