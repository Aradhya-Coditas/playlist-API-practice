package mockResources

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils/openSearch"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

func Do(ctx *gin.Context, client *openSearch.OpenSearchClient) (*opensearchapi.Response, error) {

	query := ctx.Query("symbol")
	size := ctx.Query("size")
	pageNumber := ctx.Query("pageNumber")

	if query == "TESTCONNECTION" {
		return nil, fmt.Errorf("mock error: connection failed")
	}

	if query == "ERROR_400" {
		// Return a 400 Bad Request error with an invalid JSON body
		invalidJSON := `{"error": "parsing_exception", "reason": "Invalid JSON structure"` // Missing closing brace

		return &opensearchapi.Response{
			StatusCode: 400,
			Body:       io.NopCloser(bytes.NewReader([]byte(invalidJSON))), // Invalid JSON
		}, nil
	}

	// Create mock response data
	type Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source models.OpenSearchResponse `json:"_source"`
		} `json:"hits"`
	}

	if query == "INVALID_JSON" {
		malformedJSON := `{
			"hits": {
				"total": "invalid_type_should_be_object", 
				"hits": [
					{
						"_source": {
							"field1": "value1"
						}
					}
				]
			}
		}`

		return &opensearchapi.Response{
			Body:       io.NopCloser(strings.NewReader(malformedJSON)),
			StatusCode: 200,
		}, nil
	}

	if query != "USDINR" {
		// Create an empty mock response
		mockResponse := struct {
			Hits Hits `json:"hits"`
		}{
			Hits: Hits{
				Total: struct {
					Value int `json:"value"`
				}{
					Value: 0, // Set to 0 for empty response
				},
				Hits: []struct {
					Source models.OpenSearchResponse `json:"_source"`
				}{},
			},
		}

		// Marshal the mock response to JSON
		responseBody, err := json.Marshal(mockResponse)
		if err != nil {
			return nil, err
		}

		// Create a mock response with the JSON body
		return &opensearchapi.Response{
			Body:       io.NopCloser(strings.NewReader(string(responseBody))),
			StatusCode: 204, // No Content
		}, nil
	}

	if query == "USDINR" && size == "3" && pageNumber == "1" {
		mockResponse := struct {
			Hits Hits `json:"hits"`
		}{
			Hits: Hits{
				Total: struct {
					Value int `json:"value"`
				}{
					Value: 6, // Total number of hits updated to 6
				},
				Hits: []struct {
					Source models.OpenSearchResponse `json:"_source"`
				}{
					{
						Source: models.OpenSearchResponse{
							ScripId:             "NSE_1",
							ScripToken:          "XX",
							Group:               "",
							ExchangeSegment:     "nse",
							InstrumentType:      "UNDCUR",
							SymbolName:          "USDINR",
							TradingSymbol:       "USDINR-NSE",
							OptionType:          "XX",
							UniqueKey:           "0",
							ISIN:                "",
							AssetCode:           "",
							LotSize:             1,
							TickSize:            sql.NullFloat64{Float64: 25000, Valid: true},
							ExpiryDate:          sql.NullString{Valid: false},
							Multiplier:          1000,
							DecimalPrecision:    4,
							StrikePrice:         sql.NullFloat64{Valid: false},
							DisplayStrikePrice:  sql.NullFloat64{Valid: false},
							Description:         "",
							SubGroup:            "",
							AmcCode:             "",
							ContractID:          "",
							CombinedSymbol:      "",
							ScripReserved:       "",
							ExchangeSymbolName:  "",
							RmsMarketProtection: 0,
							ContractType:        "EQUITY",
							ExchangeName:        "",
							LastUpdatedAt:       time.Date(2025, time.January, 10, 12, 30, 39, 445214000, time.FixedZone("IST", 5*60*60+30*60)),
							CompanyCode:         0,
							DisplayExpiryDate:   sql.NullString{Valid: false},
							SymbolLeg2:          "",
							RelevanceScore:      0.0,
						},
					},
					// Case 2: No symbol match + in fixed order (nfo)
					{
						Source: models.OpenSearchResponse{
							ScripId:             "NFO_2",
							ScripToken:          "1000012",
							Group:               "",
							ExchangeSegment:     "nfo",
							InstrumentType:      "SP-OPTCUR",
							SymbolName:          "EURINR",
							TradingSymbol:       "EURINR-NFO",
							OptionType:          "CE",
							UniqueKey:           "",
							ISIN:                "",
							AssetCode:           "",
							LotSize:             1,
							TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
							ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
							Multiplier:          1,
							DecimalPrecision:    4,
							StrikePrice:         sql.NullFloat64{Valid: true, Float64: 850000},
							DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 85},
							Description:         "",
							SubGroup:            "",
							AmcCode:             "",
							ContractID:          "",
							CombinedSymbol:      "",
							ScripReserved:       "",
							ExchangeSymbolName:  "",
							RmsMarketProtection: 0,
							ContractType:        "FUTURES",
							ExchangeName:        "",
							LastUpdatedAt:       time.Time{},
							CompanyCode:         0,
							DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
							SymbolLeg2:          "2117236",
							RelevanceScore:      0.0,
						},
					},
					// Case 3: Symbol match + not in fixed order
					{
						Source: models.OpenSearchResponse{
							ScripId:             "custom_exchange_3",
							ScripToken:          "1000012",
							Group:               "",
							ExchangeSegment:     "custom_exchange",
							InstrumentType:      "SP-OPTCUR",
							SymbolName:          "USDINR",
							TradingSymbol:       "USDINR-CUSTOM",
							OptionType:          "CE",
							UniqueKey:           "",
							ISIN:                "",
							AssetCode:           "",
							LotSize:             1,
							TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
							ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
							Multiplier:          1,
							DecimalPrecision:    4,
							StrikePrice:         sql.NullFloat64{Valid: true, Float64: 850000},
							DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 85},
							Description:         "",
							SubGroup:            "",
							AmcCode:             "",
							ContractID:          "",
							CombinedSymbol:      "",
							ScripReserved:       "",
							ExchangeSymbolName:  "",
							RmsMarketProtection: 0,
							ContractType:        "EQUITY",
							ExchangeName:        "BCD",
							LastUpdatedAt:       time.Time{},
							CompanyCode:         0,
							DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
							SymbolLeg2:          "2117236",
							RelevanceScore:      0.0,
						},
					},
				},
			},
		}

		// Convert mock response to JSON bytes
		responseBytes, err := json.Marshal(mockResponse)
		if err != nil {
			return nil, err
		}

		// Create mock opensearchapi.Response
		return &opensearchapi.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(responseBytes)),
		}, nil
	}

	mockResponse := struct {
		Hits Hits `json:"hits"`
	}{
		Hits: Hits{
			Total: struct {
				Value int `json:"value"`
			}{
				Value: 6, // Total number of hits updated to 6
			},
			Hits: []struct {
				Source models.OpenSearchResponse `json:"_source"`
			}{
				// Case 1: Symbol match with search text "USDINR" + in fixed order (nse)
				{
					Source: models.OpenSearchResponse{
						ScripId:             "NSE_1",
						ScripToken:          "XX",
						Group:               "",
						ExchangeSegment:     "nse",
						InstrumentType:      "UNDCUR",
						SymbolName:          "USDINR",
						TradingSymbol:       "USDINR-NSE",
						OptionType:          "XX",
						UniqueKey:           "0",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25000, Valid: true},
						ExpiryDate:          sql.NullString{Valid: false},
						Multiplier:          1000,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: false},
						DisplayStrikePrice:  sql.NullFloat64{Valid: false},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "EQUITY",
						ExchangeName:        "",
						LastUpdatedAt:       time.Date(2025, time.January, 10, 12, 30, 39, 445214000, time.FixedZone("IST", 5*60*60+30*60)),
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{Valid: false},
						SymbolLeg2:          "",
						RelevanceScore:      0.0,
					},
				},
				// Case 2: No symbol match + in fixed order (nfo)
				{
					Source: models.OpenSearchResponse{
						ScripId:             "NFO_2",
						ScripToken:          "1000012",
						Group:               "",
						ExchangeSegment:     "nfo",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "EURINR",
						TradingSymbol:       "EURINR-NFO",
						OptionType:          "CE",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: true, Float64: 850000},
						DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 85},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "FUTURES",
						ExchangeName:        "",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
						SymbolLeg2:          "2117236",
						RelevanceScore:      0.0,
					},
				},
				// Case 3: Symbol match + not in fixed order
				{
					Source: models.OpenSearchResponse{
						ScripId:             "custom_exchange_3",
						ScripToken:          "1000012",
						Group:               "",
						ExchangeSegment:     "custom_exchange",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "USDINR",
						TradingSymbol:       "USDINR-CUSTOM",
						OptionType:          "CE",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: true, Float64: 850000},
						DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 85},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "EQUITY",
						ExchangeName:        "BCD",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
						SymbolLeg2:          "2117236",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "other_exchange_4",
						ScripToken:          "1000013",
						Group:               "",
						ExchangeSegment:     "other_exchange",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "GBPINR",
						TradingSymbol:       "GBPINR-OTHER",
						OptionType:          "PE",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 50, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: true, Float64: 750000},
						DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 75},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "FUTURES",
						ExchangeName:        "",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
						SymbolLeg2:          "2117237",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "bse_5",
						ScripToken:          "1000014",
						Group:               "",
						ExchangeSegment:     "bse",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "JPYINR",
						TradingSymbol:       "JPYINR-BSE",
						OptionType:          "XX",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: false},
						DisplayStrikePrice:  sql.NullFloat64{Valid: false},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "EQUITY",
						ExchangeName:        "BSE",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{Valid: false},
						SymbolLeg2:          "",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "custom_2_6",
						ScripToken:          "1000015",
						Group:               "",
						ExchangeSegment:     "custom_2",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "AUDINR",
						TradingSymbol:       "AUDINR-CUSTOM",
						OptionType:          "PE",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 50, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: true, Float64: 650000},
						DisplayStrikePrice:  sql.NullFloat64{Valid: true, Float64: 65},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "FUTURES",
						ExchangeName:        "Custom2",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{String: "17Jan17Jan", Valid: true},
						SymbolLeg2:          "2117238",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "nse_eq_7",
						ScripToken:          "1000016",
						Group:               "",
						ExchangeSegment:     "nse_eq",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "CHFINR",
						TradingSymbol:       "CHFINR-NSE-EQ",
						OptionType:          "XX",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: false},
						DisplayStrikePrice:  sql.NullFloat64{Valid: false},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "EQUITY",
						ExchangeName:        "NSE_EQ",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{Valid: false},
						SymbolLeg2:          "",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "bse_a_8",
						ScripToken:          "1000017",
						Group:               "",
						ExchangeSegment:     "bse_a",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "CHFINR",
						TradingSymbol:       "CHFINR-BSE-A",
						OptionType:          "XX",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: false},
						DisplayStrikePrice:  sql.NullFloat64{Valid: false},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "EQUITY",
						ExchangeName:        "BSE_A",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{Valid: false},
						SymbolLeg2:          "",
						RelevanceScore:      0.0,
					},
				},
				{
					Source: models.OpenSearchResponse{
						ScripId:             "mcx_9",
						ScripToken:          "1000018",
						Group:               "XYZ",
						ExchangeSegment:     "mcx",
						InstrumentType:      "SP-OPTCUR",
						SymbolName:          "CHFINR",
						TradingSymbol:       "CHFINR-MCX",
						OptionType:          "XX",
						UniqueKey:           "",
						ISIN:                "",
						AssetCode:           "",
						LotSize:             1,
						TickSize:            sql.NullFloat64{Float64: 25, Valid: true},
						ExpiryDate:          sql.NullString{String: "1737138599|1737138599", Valid: true},
						Multiplier:          1,
						DecimalPrecision:    4,
						StrikePrice:         sql.NullFloat64{Valid: false},
						DisplayStrikePrice:  sql.NullFloat64{Valid: false},
						Description:         "",
						SubGroup:            "",
						AmcCode:             "",
						ContractID:          "",
						CombinedSymbol:      "",
						ScripReserved:       "",
						ExchangeSymbolName:  "",
						RmsMarketProtection: 0,
						ContractType:        "FUTURES",
						ExchangeName:        "MCX",
						LastUpdatedAt:       time.Time{},
						CompanyCode:         0,
						DisplayExpiryDate:   sql.NullString{Valid: false},
						SymbolLeg2:          "",
						RelevanceScore:      0.0,
					},
				},
			},
		},
	}

	// Convert mock response to JSON bytes
	responseBytes, err := json.Marshal(mockResponse)
	if err != nil {
		return nil, err
	}

	// Create mock opensearchapi.Response
	return &opensearchapi.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(responseBytes)),
	}, nil
}
