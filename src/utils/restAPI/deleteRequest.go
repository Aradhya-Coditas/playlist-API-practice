package restAPI

import (
	"context"
	"fmt"
	"io"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"

	"github.com/bytedance/sonic"
)

// HttpDELETERequest sends an HTTP DELETE request to the specified baseURL with the provided requestPayload.
func HttpDeleteRequest(ctx context.Context, baseURL string, requestPayload *models.Request, appName string) (interface{}, error) {
	var data interface{}
	response, err := HttpRequest(ctx, baseURL, requestPayload, nil, appName)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	// Decode the API response into the struct
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ResponseBodyReadError, err)
	}

	err = sonic.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.HttpDeleteResponseDecodeError, err)
	}

	return data, nil
}
