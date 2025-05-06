package configs

import "omnenest-backend/src/constants"

var BaseUrls map[string]map[string]string

// GetNestAPIUrl returns the API URL for the given endpoint.
func GetNestAPIUrl(endpoint string) string {
	return BaseUrls[constants.OmnenestKey][endpoint]
}
