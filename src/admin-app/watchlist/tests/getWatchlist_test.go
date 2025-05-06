package tests

import (
	"admin-app/watchlist/business"
	// "admin-app/watchlist/commons/constants"
	"admin-app/watchlist/handlers"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	"errors"
	"net/http"
	"net/http/httptest"
	genericConstants "omnenest-backend/src/constants"
	setupTest "omnenest-backend/src/setupTest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bytedance/sonic"
	// "gorm.io/gorm"
)

func (suite *watchlist) prepareGetWatchlistSetup(request interface{}, expectedResponseBytes []byte, expectedStatusCode int) {
	w := httptest.NewRecorder()
	ctx := setupTest.GetTestGinContext(w)
	setupTest.MockJsonPost(ctx, request)

	getWatchlistRepository := repositories.NewGetWatchlistRepository()
	getWatchlistService := business.NewGetWatchlistService(getWatchlistRepository)
	getWatchlistController := handlers.NewGetWatchlistController(getWatchlistService)

	getWatchlistController.HandleGetWatchlist(ctx)

	suite.EqualValues(expectedStatusCode, w.Code)
	suite.EqualValues(string(expectedResponseBytes), w.Body.String())
}

func (suite *watchlist) TestMockGetWatchlist400MismatchedDataType() {
	requestData := map[string]interface{}{
		"userID":   "123",
		"brokerID": 2,
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message: []models.ErrorMessage{
			{Key: "userId", ErrorMessage: genericConstants.JsonBindingFieldError},
		},
	}
	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareGetWatchlistSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

func (suite *watchlist) TestMockGetWatchlist400MissingFields() {
	requestData := models.BFFGetWatchlistRequest{
		UserID:   nil,
		BrokerID: nil,
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: "userId: required, brokerId: required"},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareGetWatchlistSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

// NOT WORKING
func (suite *watchlist) TestMockGetWatchlist204NoContent() {
	userID := uint64(8)
	brokerID := uint64(3)
	requestData := models.BFFGetWatchlistRequest{
		UserID:   &userID,
		BrokerID: &brokerID,
	}

	// Mock CheckUserIdExists
	userIdQuery := regexp.QuoteMeta(`SELECT "id" FROM "users_info" WHERE "id" = $1 ORDER BY "users_info"."id" LIMIT $2`)
	setupTest.GetSqlMock().ExpectQuery(userIdQuery).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	// Mock CheckBrokerIdExists
	brokerIdQuery := regexp.QuoteMeta(`SELECT count(*) FROM "users_info" JOIN brokers ON users_info.broker_id = brokers.id WHERE users_info.id = $1 AND brokers.id = $2`)
	setupTest.GetSqlMock().ExpectQuery(brokerIdQuery).
		WithArgs(userID, brokerID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Mock GetUserWatchlist to return no records
	userWatchlistQuery := regexp.QuoteMeta(`SELECT "id","watchlist_name" FROM "watchlists" WHERE "user_id" = $1`)
	setupTest.GetSqlMock().ExpectQuery(userWatchlistQuery).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "watchlist_name"}))

	// Mock GetBrokerWatchlist to return no records
	brokerWatchlistQuery := regexp.QuoteMeta(`SELECT "id","watchlist_name" FROM "broker_watchlists" WHERE "broker_id" = $1`)
	setupTest.GetSqlMock().ExpectQuery(brokerWatchlistQuery).
		WithArgs(brokerID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "watchlist_name"}))

	// Expect an empty response body for 204 No Content
	suite.prepareGetWatchlistSetup(requestData, []byte(""), http.StatusNoContent)

	// Validate mock expectations
	err := setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockGetWatchlist500DBError() {
	userID := uint64(5)
	brokerID := uint64(1)
	requestData := models.BFFGetWatchlistRequest{
		UserID:   &userID,
		BrokerID: &brokerID,
	}

	// Mock CheckUserIdExists to return a database error
	userIdQuery := regexp.QuoteMeta(`SELECT "id" FROM "users_info" WHERE "id" = $1 ORDER BY "users_info"."id" LIMIT $2`)
	setupTest.GetSqlMock().ExpectQuery(userIdQuery).
		WithArgs(userID, 1).
		WillReturnError(errors.New(genericConstants.MockDBError))

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusInternalServerError),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.InternalServerError},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareGetWatchlistSetup(requestData, expectedResponseBytes, http.StatusInternalServerError)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockGetWatchlist200Successful() {
	userID := uint64(5)
	brokerID := uint64(1)
	requestData := models.BFFGetWatchlistRequest{
		UserID:   &userID,
		BrokerID: &brokerID,
	}

	// Mock CheckUserIdExists
	userIdQuery := regexp.QuoteMeta(`SELECT "id" FROM "users_info" WHERE "id" = $1 ORDER BY "users_info"."id" LIMIT $2`)
	setupTest.GetSqlMock().ExpectQuery(userIdQuery).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	// Mock CheckBrokerIdExists
	brokerIdQuery := regexp.QuoteMeta(`SELECT count(*) FROM "users_info" JOIN brokers ON users_info.broker_id = brokers.id WHERE users_info.id = $1 AND brokers.id = $2`)
	setupTest.GetSqlMock().ExpectQuery(brokerIdQuery).
		WithArgs(userID, brokerID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Mock GetUserWatchlist
	userWatchlistQuery := regexp.QuoteMeta(`SELECT "id","watchlist_name" FROM "watchlists" WHERE "user_id" = $1`)
	setupTest.GetSqlMock().ExpectQuery(userWatchlistQuery).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "watchlist_name"}).AddRow(1, "User Watchlist 1"))

	// Mock GetBrokerWatchlist count query
	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "users_info" JOIN brokers ON users_info.broker_id = brokers.id WHERE users_info.id = $1 AND brokers.id = $2`)
	setupTest.GetSqlMock().ExpectQuery(countQuery).
		WithArgs(userID, brokerID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Mock GetBrokerWatchlist fetch query
	brokerWatchlistQuery := regexp.QuoteMeta(`SELECT "id","watchlist_name" FROM "broker_watchlists" WHERE "broker_id" = $1 AND "user_id" = $2`)
	setupTest.GetSqlMock().ExpectQuery(brokerWatchlistQuery).
		WithArgs(brokerID, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "watchlist_name"}).AddRow(1, "Broker Watchlist 1"))

	expectedResponse := models.BFFWatchlistResponse{
		Userdefine: []models.BFFUserdefine{
			{Id: 1, WatchlistName: "User Watchlist 1"},
		},
		Predefine: []models.BFFPredefine{
			{Id: 1, WatchlistName: "Broker Watchlist 1"},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(map[string]interface{}{
		"message": genericConstants.BFFResponseSuccessMessage,
		"data":    expectedResponse,
	})
	suite.NoError(err)
	suite.prepareGetWatchlistSetup(requestData, expectedResponseBytes, http.StatusOK)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}
