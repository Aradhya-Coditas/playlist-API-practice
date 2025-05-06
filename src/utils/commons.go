package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	genericConstants "omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"math/rand"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"hermannm.dev/ipfinder"
)

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ContainsArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[string]int)
	for _, item := range a {
		count[item]++
	}
	for _, item := range b {
		if count[item] == 0 {
			return false
		}
		count[item]--
	}
	return true
}

func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	intersection := []string{}
	for _, item := range slice1 {
		m[item] = true
	}
	for _, item := range slice2 {
		if _, ok := m[item]; ok {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

func ParseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		// TODO: handle the error appropriately
		panic(err)
	}
	return t
}

// ConvertDate takes a string representing a date in the format "dd/mm/yyyy" and converts it to the format "mm/dd/yyyy".
func ConvertDate(dateString string) string {
	dateParts := strings.Split(dateString, "/")
	return fmt.Sprintf("%s/%s/%s", dateParts[1], dateParts[0], dateParts[2])
}

func StructToMap(ctx context.Context, input interface{}) map[string][]string {
	_, span := tracer.AddToSpan(ctx, "StructToMap")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	result := make(map[string][]string)
	structValue := reflect.ValueOf(input)

	if structValue.Kind() != reflect.Struct {
		return result
	}

	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := structValue.Field(i)

		if fieldValue.Kind() == reflect.Slice {
			sliceLen := fieldValue.Len()
			if sliceLen > 0 {
				values := make([]string, sliceLen)
				for j := 0; j < sliceLen; j++ {
					values[j] = fieldValue.Index(j).Interface().(string)
				}
				result[fieldName] = values
			}
		}
	}
	return result
}

// FilterResponse filters the given response based on the provided filters.
func FilterResponse[T any](ctx context.Context, response []T, filters map[string][]string) []T {
	_, span := tracer.AddToSpan(ctx, "FilterResponse")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	var filteredResponse []T

	for _, item := range response {
		allFiltersSatisfied := true
		for filterKey, filterValues := range filters {
			if filterValues == nil {
				continue
			}

			itemFieldValue := reflect.ValueOf(item).FieldByName(filterKey).Interface()
			anyFilterValueSatisfied := false
			for _, filterValue := range filterValues {
				if itemFieldValue == filterValue {
					anyFilterValueSatisfied = true
					break
				}
			}
			if !anyFilterValueSatisfied {
				allFiltersSatisfied = false
				break
			}
		}
		if allFiltersSatisfied {
			filteredResponse = append(filteredResponse, item)
		}
	}
	return filteredResponse
}

func ArrayToString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	return strings.Join(arr, ", ")
}

func StringToArray(input string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, ", ")
}

func CreateKeyValueMap(ctx context.Context, key []string, values []string) map[string]int {
	_, span := tracer.AddToSpan(ctx, "CreateKeyValueMap")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	result := make(map[string]int, len(key))
	var keyValuePairs []struct {
		key   string
		value int
	}

	for i, k := range key {
		intValue, err := strconv.Atoi(values[i])
		if err == nil {
			keyValuePairs = append(keyValuePairs, struct {
				key   string
				value int
			}{key: k, value: intValue})
		}
	}

	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].key < keyValuePairs[j].key
	})

	// Populate the result map with sorted key-value pairs
	for _, pair := range keyValuePairs {
		result[pair.key] = pair.value
	}

	return result
}

func SafeDivision(numerator, denominator float32) float32 {
	if denominator == 0 || numerator == 0 {
		return 0
	}
	return numerator / denominator
}

// FilterInstruments filters out instruments from the given list based on a predefined filter.
func FilterInstruments(instruments []string) []string {
	var filtered []string
	for _, instrument := range instruments {
		if !genericConstants.InstrumentFilterOut[instrument] {
			filtered = append(filtered, instrument)
		}
	}
	return filtered
}

func SplitStringWithSeparator(input string, separator string) []string {
	return strings.Split(input, separator)
}

func IsSpreadInstrumentType(instrument string) bool {
	upperCaseInstrument := strings.ToUpper(instrument)
	return strings.HasPrefix(upperCaseInstrument, genericConstants.SpreadContractIdentifier)
}

// SortWithFixedPart sorts the input slice of strings based on a fixed order defined by the fixedArray slice.
func SortWithFixedPart(ctx context.Context, input, fixedArray []string) []string {
	_, span := tracer.AddToSpan(ctx, "SortWithFixedPart")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	// Create a map to store the indices of elements in fixedArray
	fixedIndices := make(map[string]int)
	for i, element := range fixedArray {
		fixedIndices[element] = i
	}

	sort.SliceStable(input, func(i, j int) bool {
		indexI, existsI := fixedIndices[input[i]]
		indexJ, existsJ := fixedIndices[input[j]]

		// If both elements are in fixedArray, compare their indices
		if existsI && existsJ {
			return indexI < indexJ
		}

		// If one element is in fixedArray and the other is not, prioritize the one in fixedOrder
		if existsI {
			return true
		} else if existsJ {
			return false
		}

		// If neither element is in fixedArray, sort them in alphabetical order
		return input[i] < input[j]
	})

	return input
}

func GetUintReferenceValue(uintValue uint) *uint {
	return &uintValue
}

func GetUint16ReferenceValue(uint16Value uint16) *uint16 {
	return &uint16Value
}

func GetUint32ReferenceValue(uint32Value uint32) *uint32 {
	return &uint32Value
}

func GetUint64ReferenceValue(uint64Value uint64) *uint64 {
	return &uint64Value
}

func GetFloat32ReferenceValue(float32Value float32) *float32 {
	return &float32Value
}

func GetFloat64ReferenceValue(float64Value float64) *float64 {
	return &float64Value
}

func GetInt32ReferenceValue(int32Value int32) *int32 {
	return &int32Value
}

func GetInt64ReferenceValue(int64Value int64) *int64 {
	return &int64Value
}

func StringToUint64(s string) (uint64, error) {
	uint64Value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64Value, nil
}

func StringToUint16(s string) (uint16, error) {
	uint64Value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint16(uint64Value), nil
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func InitApiUrls() {
	// Get a Apis config from yaml
	ApisConfig, err := configs.Get(genericConstants.ApiConfig)
	if err != nil {
		panic(err)
	}
	// Unmarshal the YAML data into the baseUrls map
	err = ApisConfig.Unmarshal(&configs.BaseUrls)
	if err != nil {
		panic(fmt.Errorf(genericConstants.APIYamlFileUnmarshalError, err))
	}
	configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIRestBaseURL] = GetEnv(genericConstants.RestBaseUrlKey, configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIRestBaseURL])
	configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIScannerBaseURL] = GetEnv(genericConstants.ScannerBaseUrlKey, configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIScannerBaseURL])
	configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIGlobalSearchBaseURL] = GetEnv(genericConstants.GlobalSearchBaseUrlKey, configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIGlobalSearchBaseURL])
	configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIIPOBaseURL] = GetEnv(genericConstants.IPOBaseUrlKey, configs.BaseUrls[genericConstants.NestAPITypeToURLMapping][genericConstants.NestAPIIPOBaseURL])
	configs.BaseUrls[genericConstants.CMOTSAPITypeToURLmapping][genericConstants.CMOTSAPIBaseURL] = GetEnv(genericConstants.CMOTSBaseURLKey, configs.BaseUrls[genericConstants.CMOTSAPITypeToURLmapping][genericConstants.CMOTSAPIBaseURL])
}

func ConvertPaisaToRupees(priceInPaise float32, multiplier uint, decimalPrecision uint16) float32 {
	var denominator float32
	denominator = genericConstants.DefaultDenominatorFloatValue
	if multiplier != 0 && decimalPrecision != 0 {
		denominator = float32(multiplier) * float32(math.Pow(10, float64(decimalPrecision)))
	}
	return SafeDivision(priceInPaise, denominator)
}

func ConvertRupeeToPaisa(rupees float32, multiplier uint, decimalPrecision uint16) int64 {
	var denominator float32
	denominator = genericConstants.DefaultDenominatorFloatValue
	if multiplier != 0 && decimalPrecision != 0 {
		denominator = float32(multiplier) * float32(math.Pow(10, float64(decimalPrecision)))
	}
	return int64(rupees * denominator)
}

func ConvertDecimalRupeeToPaisa(rupees decimal.Decimal, multiplier decimal.Decimal, decimalPrecision uint16) int64 {
	denominator := decimal.NewFromFloat(genericConstants.DefaultDenominatorFloatValue)
	if multiplier.Cmp(decimal.Zero) != 0 && decimalPrecision != 0 {
		denominator = multiplier.Mul(decimal.NewFromFloat(math.Pow(10, float64(decimalPrecision))))
	}
	return int64(rupees.Mul(denominator).Round(0).IntPart())
}

func ConvertDecimalPaisaToRupee(paisa decimal.Decimal, multiplier decimal.Decimal, decimalPrecision uint16) float32 {
	denominator := decimal.NewFromFloat(genericConstants.DefaultDenominatorFloatValue)
	if multiplier.Cmp(decimal.Zero) != 0 && decimalPrecision != 0 {
		denominator = multiplier.Mul(decimal.NewFromFloat(math.Pow(10, float64(decimalPrecision))))
	}

	res, _ := paisa.Div(denominator).Round(int32(decimalPrecision)).Float64()
	return float32(res)
}

func GetMultiplier(exchangeSegment string, scripMasterDetails *genericModels.ScripMaster) decimal.Decimal {
	// TODO: need to change this function when we get price multiplier from the nest
	if genericConstants.MultiplierExchangeMapping[exchangeSegment] == genericConstants.MultiplierTag {
		return decimal.NewFromUint64(uint64(scripMasterDetails.Multiplier))
	}
	return decimal.NewFromInt(genericConstants.DefaultMultiplierValue)
}

func CreateNestOrderSource(ctx *gin.Context) string {
	return genericConstants.NestHeaderSourcePrefixValue + strings.ToUpper(ctx.GetHeader(genericConstants.Source))
}

func ReplaceSpecialCharsWithSpaceRegex(input string) string {
	regex := configs.GetRegexPattern(genericConstants.ReplaceSpecialCharsWithSpaceKey)
	return regex.ReplaceAllString(strings.TrimSpace(input), " ")
}

func ConvertStringToTime(timeInput string, timeFormat string) (time.Time, error) {
	parsedTime, err := time.Parse(timeFormat, timeInput)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// Add your common functionalities here

func GetIPAndMAC() (ip *string, mac *string, appVersion *string, loginSource *string) {
	localIPs, err := ipfinder.FindLocalIPs()
	if err != nil {
		return nil, nil, nil, nil
	}

	ipStr := localIPs[0].Address.String()
	macStr := localIPs[0].NetworkInterface.HardwareAddr.String()
	appVersionStr := "1.0.0"
	loginSourceStr := "TWS"

	ip = &ipStr
	mac = &macStr
	appVersion = &appVersionStr
	loginSource = &loginSourceStr

	return
}

func ConstructCMOTSURL(endpoint string, params []string) string {
	var url strings.Builder
	url.WriteString(endpoint)
	for _, param := range params {
		url.WriteString("/")
		url.WriteString(param)
	}
	return url.String()
}

// createResponse is a function that takes in data and converts it to JSON format and then creates a new HTTP response with the provided status code and JSON body.
func CreateResponse(data interface{}) *http.Response {
	// Convert the data map to JSON
	jsonData, err := sonic.Marshal(data)
	if err != nil {
		// Handle the error if JSON marshaling fails
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader(err.Error())),
		}
	}

	// Create a new response with the provided status code and JSON body
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(jsonData)),
		Header:     http.Header{genericConstants.ContentType: []string{genericConstants.ApplicationJSONTypeConfig}},
	}
}

func JoinStringsWithSeparator(stringArr []string, separator string) string {
	return strings.Join(stringArr, separator)
}

// PadKey ensures the key is 16, 24, or 32 bytes long by padding or truncating it.
func PadKey(key []byte) []byte {
	if len(key) == 16 || len(key) == 24 || len(key) == 32 {
		return key
	}
	if len(key) < 16 {
		return append(key, make([]byte, 16-len(key))...)
	} else if len(key) < 24 {
		return append(key, make([]byte, 24-len(key))...)
	} else if len(key) < 32 {
		return append(key, make([]byte, 32-len(key))...)
	}

	return key[:32]
}

func IsGTCGTDExchange(exchangeToCheck string) bool {
	for _, exchange := range genericConstants.Exchanges {
		if exchangeToCheck == exchange {
			return true
		}
	}
	return false
}

func ConvertDateToEpoch(inputDate string) (string, error) {
	date, err := time.Parse(genericConstants.DateOldLayout2, inputDate)
	if err != nil {
		return "", err
	}
	epoch := fmt.Sprintf("%d", date.Unix())
	return epoch, nil
}

func ConvertEpochToDate(expiryDate int64, dateLayout string) (string, string) {
	timestamp := time.Unix(expiryDate, 0)

	ist, err := time.LoadLocation(genericConstants.IST)
	if err != nil {
		return "", ""
	}
	istTimestamp := timestamp.In(ist)

	display := istTimestamp.Format(dateLayout)
	value := strconv.FormatInt(expiryDate, 10)
	return display, value
}

// ConvertToScripMasterMap converts a slice of ScripMaster to a map with scrip token as key
func ConvertToScripMasterMap(scripMasterDetails []genericModels.ScripMaster) map[string]*genericModels.ScripMaster {
	scripMasterMap := make(map[string]*genericModels.ScripMaster, len(scripMasterDetails)) // Pre-allocate map capacity
	for i := range scripMasterDetails {
		scripMasterMap[PrepareScripMasterMapKey(scripMasterDetails[i].ExchangeSegment, scripMasterDetails[i].ScripToken)] = &scripMasterDetails[i]
		if scripMasterDetails[i].CombinedScripToken != "" {
			scripMasterMap[PrepareScripMasterMapKey(scripMasterDetails[i].ExchangeSegment, scripMasterDetails[i].CombinedScripToken)] = &scripMasterDetails[i]
		}
	}
	return scripMasterMap
}

func GetGuiId(username string) string {
	currentTime := time.Now()
	seconds := currentTime.Unix()
	randomNum := 10000 + rand.Intn(90000)
	return fmt.Sprintf("%s%d-%d", username, seconds, randomNum)
}

func PrepareRedisKey(exchangeSegment string, scripToken string) string {
	var market string = "MKT"
	return strings.ToUpper(market + ":" + exchangeSegment + ":" + scripToken)
}

func PrepareScripMasterMapKey(exchangeSegment string, scripToken string) string {
	return strings.ToUpper(exchangeSegment + ":" + scripToken)
}
