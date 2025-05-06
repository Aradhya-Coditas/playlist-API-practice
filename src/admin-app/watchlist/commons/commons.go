package commons

import (
	"admin-app/watchlist/models"
	"strings"
)

func PrepareBFFResponseScrips(scrips string, scripGroupSeparator string, scripSubgroupSeparator string) []models.BFFScrip {
	if scrips == "" {
		return []models.BFFScrip{}
	}
	scripGroups := strings.Split(scrips, scripGroupSeparator)
	var preparedScrips []models.BFFScrip
	for _, scripGroup := range scripGroups {
		scripParts := strings.Split(scripGroup, scripSubgroupSeparator)
		exchangeName := scripParts[0]
		scripToken := scripParts[1]

		scrip := models.BFFScrip{ExchangeName: exchangeName, ScripToken: scripToken}
		preparedScrips = append(preparedScrips, scrip)
	}

	return preparedScrips
}
