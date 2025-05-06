package configs

import (
	"omnenest-backend/src/constants"
	"regexp"
)

var compiledRegexMap map[string]*regexp.Regexp

var regexPatterns = map[string]string{
	constants.NonDigitSequenceKey:               constants.NonDigitSequence,
	constants.DigitKey:                          constants.Digit,
	constants.InstrumentTypeScripTokenFormatKey: constants.SpreadInstrumentTypeScripTokenFormat,
	constants.DecimalZeroOrCommaKey:             constants.DecimalZeroOrComma,
	constants.InitialPriceTagKey:                constants.InitialPriceTag + constants.BrokerRecommendationPattern,
	constants.StopLossTagKey:                    constants.StopLossTag + constants.BrokerRecommendationPattern,
	constants.TargetTagKey:                      constants.TargetTag + constants.BrokerRecommendationPattern,
	constants.PriceAtCallTagKey:                 constants.PriceAtCallTag + constants.BrokerRecommendationPattern,
	constants.DateFormatMatchKey:                constants.DateFormatMatch,
	constants.WatchListLimitErrorKey:            constants.WatchListLimitError,
	constants.WatchListMaxScripsMatchKey:        constants.WatchListMaxScripsMatch,
	constants.MatchHttpUrlKey:                   constants.MatchHttpUrl,
	constants.ReplaceSpecialCharsWithSpaceKey:   constants.ReplaceSpecialCharsWithSpace,
	constants.PanCardValidationKey:              constants.PANFormat,
}

// InitRegexPatterns initializes the compiledRegexMap with regular expression patterns.
func InitRegexPatterns() {
	// Extracting the password validation regex from the yml file as can vary for different brokers.
	applicationConfig := GetApplicationConfig()
	authenticationConfig := applicationConfig.Authentication
	dataLoadConfig := applicationConfig.DataLoad
	regexPatterns[constants.PasswordValidationRegex] = authenticationConfig.PasswordValidationRegex
	regexPatterns[constants.CmotsYearMonthRegex] = dataLoadConfig.CmotsYearMonthRegex

	// Orders Regex validations
	ordersConfig := applicationConfig.Orders
	regexPatterns[constants.BasketNameValidationRegex] = ordersConfig.BasketNameValidationRegex

	// Watchlist Regex validations
	watchlistConfig := applicationConfig.Watchlist
	regexPatterns[constants.WatchlistNameValidationRegex] = watchlistConfig.WatchlistNameValidationRegex

	compiledRegexMap = make(map[string]*regexp.Regexp)
	for key, pattern := range regexPatterns {
		compiledRegexMap[key] = regexp.MustCompile(pattern)
	}
}

// GetRegexPatternImpl allows for easier mocking in tests
var GetRegexPatternImpl = func(key string) *regexp.Regexp {
	return compiledRegexMap[key]
}

// GetRegexPattern returns the compiled regular expression pattern associated with the given key
func GetRegexPattern(key string) *regexp.Regexp {
	return GetRegexPatternImpl(key)
}
