package restAPI

import (
	"context"
	"fmt"
	"io"
	"net/url"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"

	"github.com/bytedance/sonic"
)

// HTTPGetRequest is a function that sends a HTTP GET request to a specified API endpoint and returns the response data.
func HTTPGetRequest(ctx context.Context, baseURL string, requestPayload *models.Request, queryParams url.Values, appName string) (interface{}, error) {
	var data interface{}

	// Uncomment in future if caching to be done at BFF
	// client := redis.GetRedisClient()
	// if cachedData, err := client.Get(ctx, commons.GenerateRedisKey(requestPayload)); err == nil {
	// 	_ = sonic.Unmarshal([]byte(cachedData), &data)
	// 	log.Info(constants.WatchListRedisRetrieved)
	// 	return data, nil
	// }

	apiURL := baseURL + "?" + EncodeQueryParams(queryParams)

	response, err := HttpRequest(ctx, apiURL, requestPayload, nil, appName)
	if err != nil {
		return data, fmt.Errorf(genericConstants.HitRestAPIError, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ResponseBodyReadError, err)
	}

	err = sonic.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.HttpGetResponseDecodeError, err)
	}
	// Uncomment in future if caching to be done at BFF
	// jsonData, err := sonic.Marshal(data)
	// if err != nil {
	// 	return data, err
	// }
	// if err := client.SetWithExpiry(ctx, commons.GenerateRedisKey(requestPayload), jsonData, time.Duration(50)*time.Second); err != nil {
	// 	return data, err
	// }

	return data, nil
}
