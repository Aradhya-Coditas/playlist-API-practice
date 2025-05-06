package setupTest

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils"

	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/metrics"
	postgresClient "omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var mock sqlmock.Sqlmock
var mockDB *sql.DB

func InitSuiteConfigs(serviceName string) {
	ctx := context.Background()
	// Init Test Configs
	workingDir, _ := os.Getwd()
	testConfig := workingDir + constants.TestConfig
	mockDir := workingDir + constants.MockResponseConfigPath
	configs.Init([]string{mockDir, testConfig})
	logger.StartLogger(ctx, constants.InfoLevel)
	err := configs.InitApplicationConfigs(ctx)
	if err != nil {
		panic(fmt.Errorf(constants.ConfigBindingFailedError))
	}
	traceProvider, _ := initTrace(ctx, serviceName)
	defer traceProvider.Shutdown(ctx)
	metrics.Init()
	MockPostgresDB()
	MockCockroachDB()

	utils.InitApiUrls()
	configs.InitRegexPatterns()
}

func MockPostgresDB() {
	mockDB, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: constants.PostgresDriverName,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	postgresClient.SetPostgresClient(db, mockDB)
}

// Creates and configures a mock CockroachDB instance using sqlmock and GORM for testing database interactions
func MockCockroachDB() {
	mockDB, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: constants.CockroachDBDriverName,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	postgresClient.SetPostgresClient(db, mockDB)
}

func initTrace(ctx context.Context, serviceName string) (*trace.TracerProvider, error) {
	traceProvider, err := tracer.InitTracer(ctx, true, serviceName, constants.OLTPHttpEndpointUrl)
	return traceProvider, err
}

func MockJsonPost(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = constants.PostMethod
	setContext(ctx)

	jsonBytes, err := sonic.Marshal(content)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func MockJsonGet(ctx *gin.Context, params gin.Params, u url.Values) {
	ctx.Request.Method = constants.GetMethod
	setContext(ctx)

	ctx.Params = params
	ctx.Request.URL.RawQuery = u.Encode()
}

func MockJsonDelete(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = constants.DeleteMethod
	setContext(ctx)

	jsonBytes, err := sonic.Marshal(content)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func setContext(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set(constants.RequestIDHeader, "1")
	ctx.Request.Header.Set(constants.DeviceIdHeader, "6")
	ctx.Request.Header.Set(constants.AppVersion, "0.0.1")
	ctx.Request.Header.Set(constants.Source, "MOB")
	ctx.Set(constants.UserID, uint64(1))
	ctx.Set(constants.CriteriaAttributeKey, []string{"Y|Y|Y"})
	ctx.Set(constants.Username, "PROTEST3")
	ctx.Set(constants.AccountID, "PROTEST3")
	ctx.Set(constants.DeviceIdHeader, "6")
	ctx.Set(constants.AccountID, "PROTEST3-OTPUAT")
	ctx.Set(constants.BrokerName, "OTPUAT")
	ctx.Set(constants.BranchName, "HO")
	ctx.Set(constants.ProductAlias, "NRML:NRML")
	ctx.Set(constants.Source, "MOB")
	ctx.Set(constants.UserSessionToken, "121212131")
	ctx.Set(constants.APIRequestTime, "2402061903")
	ctx.Set(constants.RequestIDHeader, "1")
	ctx.Set(constants.EnabledExchangesKey, []string{"BSE", "MCX", "CDS", "NSE", "NFO"})
	ctx.Set(constants.ClearingOrg, "[{\"exchArray\":[\"BSE\",\"NSE\"],\"segment\":\"CASH\"}]")
	ctx.Set(constants.GttEnabledKey, true)

	nestKeyPair := &genericModels.NestKeyPair{
		PrivateKey:       nil,
		PublicKey:        nil,
		PublicHashedKey:  "",
		PrivateHashedKey: "",
	}
	ctx.Set(constants.ServerKeyPair, nestKeyPair)
}

func GetSqlMock() sqlmock.Sqlmock {
	return mock
}

func MockDBClose() {
	mockDB.Close()
}
