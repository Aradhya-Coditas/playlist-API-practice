# Backend Test Setup

This README provides detailed instructions for setting up tests in a backend project. It outlines the necessary steps and considerations for running tests effectively.

## Steps to Setup Tests

Follow these steps to set up tests in your backend project:

1. **Add Endpoint Configuration**:
   Ensure that the endpoints are properly configured. This might involve updating configuration files, such as `apis.yml`.

2. **Create Test Initialization File**:
   In your project, create a file dedicated to test initialization. This file should set up any necessary configurations or dependencies required for testing.

    ```go
    // Example start_test.go file
    package tests

    import (
        "context"
        setupTest "path/to/setupTest"
        "testing"

        "github.com/golang/mock/gomock"
        "github.com/stretchr/testify/suite"
    )

    type TestSuite struct {
        suite.Suite
        ctrl *gomock.Controller
        ctx  context.Context
    }

    func (suite *TestSuite) SetupSuite() {
        // Initialize test configurations
        setupTest.InitSuiteConfigs()
    }

    func (suite *TestSuite) BeforeTest(suiteName, testName string) {
        // Set up before each test
        suite.ctx = context.Background()
        suite.ctrl = gomock.NewController(suite.T())
    }

    func (suite *TestSuite) AfterTest(suiteName, testName string) {
        // Clean up after each test
        suite.ctrl.Finish()
    }

    func TestAll(t *testing.T) {
        // Run tests in the suite
        suite.Run(t, new(TestSuite))
    }
    ```


4. **Create Test Files**:
   Create test files corresponding to specific functionalities.

    ```go

     func (suite *TestSuite) sampleTestSetup(request interface{}, expectedResponseBytes []byte, expectedStatusCode int) {
         // Create a new HTTP recorder for capturing response.
         w := httptest.NewRecorder()

         // Set up a test Gin context.
         ctx := setupTest.GetTestGinContext(w)

         // Mock JSON POST request.
         setupTest.MockJsonPost(ctx, request)

         // Initialize necessary dependencies.
         nestApiWrapper := nestIntegration.GetNestWrapper(true)
         sampleService := business.NewSampleService(nestApiWrapper)
         sampleController := handlers.NewSampleController(sampleService)

         // Execute the handler function being tested.
         sampleController.HandleSample(ctx)

         suite.EqualValues(expectedStatusCode, w.Code)
         suite.EqualValues(expectedResponseBytes, w.Body.Bytes())
         suite.EqualValues(string(expectedResponseBytes), w.Body.String())
     }

     // TestSample200 is a test function for testing a specific scenario.
     func (suite *TestSuite) TestSample200() {
         requestData := models.SampleRequest{}

         expectedResponse := models.SampleResponse{
             Message: "success",
         }

         expectedResponseBytes, err := sonic.Marshal(&expectedResponse)
         suite.NoError(err)
         suite.sampleTestSetup(requestData, expectedResponseBytes, http.StatusOK)
     }

    ```
5. **Run Tests**:
   Execute the tests using the following command:
    go test ./...

5. **Review Results**:
   Analyze the test results to identify any failures or unexpected behavior.

By following these steps, you can effectively set up and run tests in your backend project to ensure the reliability and correctness of your codebase.
