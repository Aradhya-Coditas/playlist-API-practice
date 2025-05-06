package tests

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/handlers"
	"admin-app/watchlist/models"
	"admin-app/watchlist/repositories"
	"errors"
	"net/http"
	"net/http/httptest"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/setupTest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bytedance/sonic"
	"gorm.io/gorm"
)

func (suite *watchlist) prepareDeleteWatchlistSetup(request interface{}, expectedResponseBytes []byte, expectedStatusCode int) {
	w := httptest.NewRecorder()
	ctx := setupTest.GetTestGinContext(w)
	setupTest.MockJsonPost(ctx, request)

	deleteWatchlistRepository := repositories.GetDeleteWatchlistRepository(false)
	deleteWatchlistService := business.NewDeleteWatchlistService(deleteWatchlistRepository)
	deleteWatchlistController := handlers.NewDeleteWatchlistController(deleteWatchlistService)

	deleteWatchlistController.HandleDeleteWatchlist(ctx)

	suite.EqualValues(expectedStatusCode, w.Code)
	suite.EqualValues(expectedResponseBytes, w.Body.Bytes())
	suite.EqualValues(string(expectedResponseBytes), w.Body.String())
}

func (suite *watchlist) TestMockDeleteWatchlist200Successful() {
	watchlistId := uint64(1030336487419445251)
	userID := uint64(1)
	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: &watchlistId,
	}

	setupTest.GetSqlMock().ExpectBegin()
	deleteScripsQuery := regexp.QuoteMeta(constants.DeleteWatchlistScripsQuery)
	setupTest.GetSqlMock().ExpectExec(deleteScripsQuery).
		WithArgs(userID, watchlistId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	deleteQuery := regexp.QuoteMeta(constants.DeleteWatchlistQuery)
	setupTest.GetSqlMock().ExpectExec(deleteQuery).
		WithArgs(watchlistId, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	setupTest.GetSqlMock().ExpectCommit()

	expectedResponse := models.BFFDeleteWatchlistResponse{
		Message: genericConstants.BFFResponseSuccessMessage,
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusOK)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockGetDeleteWatchlist400MismatchedDataType() {
	requestData := map[string]interface{}{
		"watchlistId": "120",
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message:      []models.ErrorMessage{{Key: "watchlistId", ErrorMessage: genericConstants.JsonBindingFieldError}},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)

	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

func (suite *watchlist) TestDeleteWatchlist400MissingFields() {
	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: nil,
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusBadRequest),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: "watchlistId: required"},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusBadRequest)
}

func (suite *watchlist) TestMockDeleteWatchlist204NoContent() {
	watchlistId := uint64(1030336487419445251)
	userID := uint64(1)
	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: &watchlistId,
	}

	setupTest.GetSqlMock().ExpectBegin()

	deleteScripsQuery := regexp.QuoteMeta(constants.DeleteWatchlistScripsQuery)
	setupTest.GetSqlMock().ExpectExec(deleteScripsQuery).
		WithArgs(userID, watchlistId).
		WillReturnResult(sqlmock.NewResult(0, 0))

	deleteQuery := regexp.QuoteMeta(constants.DeleteWatchlistQuery)
	setupTest.GetSqlMock().ExpectExec(deleteQuery).
		WithArgs(watchlistId, userID).
		WillReturnError(gorm.ErrRecordNotFound)

	// Expect transaction rollback instead of commit
	setupTest.GetSqlMock().ExpectRollback()
	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusNotFound),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.NoDataFoundError},
		},
	}
	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusNotFound)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockDeleteWatchlist500DBError() {
	watchlistId := uint64(1030336487419445251)
	userID := uint64(1)
	setupTest.GetSqlMock().ExpectBegin()

	deleteScripsQuery := regexp.QuoteMeta(constants.DeleteWatchlistScripsQuery)

	setupTest.GetSqlMock().ExpectExec(deleteScripsQuery).
		WithArgs(userID, watchlistId).
		WillReturnError(errors.New(genericConstants.MockDBError))

	setupTest.GetSqlMock().ExpectRollback()

	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: &watchlistId,
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusInternalServerError),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.InternalServerError},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusInternalServerError)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockDeleteWatchlist500TransactionBeginError() {
	watchlistId := uint64(1030336487419445251)

	// transaction begin error by using a mock that always returns an error
	setupTest.GetSqlMock().ExpectBegin().WillReturnError(errors.New("transaction begin failed"))
	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: &watchlistId,
	}

	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusInternalServerError),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.InternalServerError},
		},
	}

	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusInternalServerError)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *watchlist) TestMockDeleteWatchlist500TransactionCommitError() {
	watchlistId := uint64(1030336487419445251)
	userID := uint64(1)
	setupTest.GetSqlMock().ExpectBegin()

	deleteScripsQuery := regexp.QuoteMeta(constants.DeleteWatchlistScripsQuery)
	setupTest.GetSqlMock().ExpectExec(deleteScripsQuery).
		WithArgs(userID, watchlistId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	deleteQuery := regexp.QuoteMeta(constants.DeleteWatchlistQuery)
	setupTest.GetSqlMock().ExpectExec(deleteQuery).
		WithArgs(watchlistId, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	//Expect transaction commit to fail
	setupTest.GetSqlMock().ExpectCommit().WillReturnError(errors.New("transaction commit failed"))

	requestData := models.BFFDeleteWatchlistRequest{
		WatchlistId: &watchlistId,
	}
	expectedResponse := models.ErrorAPIResponse{
		ErrorMessage: http.StatusText(http.StatusInternalServerError),
		Message: []models.ErrorMessage{
			{Key: genericConstants.GenericErrorKey, ErrorMessage: genericConstants.InternalServerError},
		},
	}
	expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
	suite.NoError(err)
	suite.prepareDeleteWatchlistSetup(requestData, expectedResponseBytes, http.StatusInternalServerError)

	// Validate mock expectations
	err = setupTest.GetSqlMock().ExpectationsWereMet()
	suite.NoError(err)
}
