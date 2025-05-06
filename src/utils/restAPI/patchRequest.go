package restAPI

import (
	"context"
	"fmt"
	"io"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"

	"github.com/bytedance/sonic"
)

// HttpPatchRequest sends an HTTP PATCH request to the specified baseURL with the provided requestPayload.
func HttpPatchRequest(ctx context.Context, baseURL string, requestPayload *models.Request, appName string) (interface{}, error) {
	var data interface{}
	response, err := HttpRequest(ctx, baseURL, requestPayload, nil, appName)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ResponseBodyReadError, err)
	}

	err = sonic.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.HttpPatchResponseDecodeError, err)
	}

	return data, nil
}
