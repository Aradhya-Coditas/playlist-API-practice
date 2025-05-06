package configs

import "omnenest-backend/src/constants"

// GetCMOTSAPIUrl returns the API URL for the given endpoint.
func GetCMOTSAPIUrl(endpoint string) string {
	return BaseUrls[constants.CMOTSKey][endpoint]
}
