package models

import (
	"database/sql"
	"time"
)

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message []ErrorMessage `json:"errors,omitempty"`
	Error   string         `json:"error,omitempty"`
}

type Request struct {
	Type     string      `json:"type,omitempty"`                        // Sender
	Method   string      `json:"method,omitempty"`                      // Recipient
	ClientId string      `json:"clientId,omitempty" logger:"sensitive"` // Content
	UserId   string      `json:"userId,omitempty" logger:"sensitive"`   // Content
	Request  interface{} `json:"request,omitempty"`                     // Content
}

type NestAPIResponse struct {
	ResponseBody interface{}
	StatusCode   int
}

type EncryptResponse struct {
	EncResponse string `json:"encResponse"`
}

type EncryptRequest struct {
	EncRequest string `json:"encRequest" validate:"required"`
}

type EncryptedNestAPIResponse struct {
	EncryptedResponse string `json:"jEncResp,omitempty"`
}

type TokenData struct {
	UserId            uint64       `json:"userId" logger:"sensitive"`
	Username          string       `json:"username" logger:"sensitive"`
	ServerPublicKey   *NestKeyPair `json:"serverKeyPair"`
	UserSessionId     string       `json:"userSessionId"`
	BFFPublicKey      string       `json:"bffPublicKey"`
	BFFPrivateKey     string       `json:"bffPrivateKey"`
	DevicePublicKey   string       `json:"devicePublicKey"`
	AccountId         string       `json:"accountId" logger:"sensitive"`
	BrokerName        string       `json:"brokerName"`
	BranchName        string       `json:"branchName"  logger:"sensitive"`
	ProductAlias      string       `json:"productAlias"`
	CriteriaAttribute []string     `json:"criteriaAttribute"`
	ClearingOrg       string       `json:"clearingOrg"`
	EnabledExchanges  []string     `json:"enabledExchange"`
	GttEnabled        bool         `json:"gttEnabled"`
}

type ChannelResponse struct {
	ApiEndpoint string
	Response    []byte
	Error       error
	Metadata    interface{}
}

type HttpGoRoutineRequest struct {
	ApiEndpoint      string
	Request          interface{}
	Metadata         interface{}
	MicroServiceName string
	Params           []string
}

type NestGetAccountDetailsRequest struct {
	UserId       string `json:"uid" logger:"sensitive"`
	AccountId    string `json:"acctId" logger:"sensitive"`
	ProductAlias string `json:"s_prdt_ali"`
}

type NestAccountDetailsResponse struct {
	Status                 string        `json:"stat"`
	ErrorMessage           string        `json:"Emsg,omitempty"`
	DematAccountNumber     string        `json:"dpAccountNumber" logger:"sensitive"`
	DematAccountNumber2    string        `json:"dpAccountNumber2" logger:"sensitive"`
	DematAccountNumber3    string        `json:"dpAccountNumber3" logger:"sensitive"`
	DematAccountNumber4    string        `json:"dpAccountNumber4" logger:"sensitive"`
	DematAccountNumber5    string        `json:"dpAccountNumber5" logger:"sensitive"`
	DematAccountNumber6    string        `json:"dpAccountNumber6" logger:"sensitive"`
	StatusIndicator        string        `json:"statusIndicator"`
	ExchangeEnabled        string        `json:"exchEnabled"`
	Name                   string        `json:"accountName" logger:"sensitive"`
	Email                  string        `json:"emailAddr" logger:"sensitive"`
	PhoneNumber            string        `json:"cellAddr" logger:"sensitive"`
	Address                string        `json:"address" logger:"sensitive"`
	AccountType            string        `json:"accountType"`
	AccountStatus          string        `json:"accountStatus"`
	AccountId              string        `json:"accountId" logger:"sensitive"`
	BankId                 string        `json:"bankId" logger:"sensitive"`
	BankName               string        `json:"bankName" logger:"sensitive"`
	BankAccountNumber      string        `json:"bankAccountNo" logger:"sensitive"`
	BankAddress            string        `json:"bankAddres" logger:"sensitive"`
	BankBranchName         string        `json:"bankBranchName" logger:"sensitive"`
	DpId                   string        `json:"dpId" logger:"sensitive"`
	DpName                 string        `json:"dpName"`
	DpType                 string        `json:"dpType"`
	OfficeAddress          string        `json:"officeAddress"`
	CustomerId             string        `json:"customerId" logger:"sensitive"`
	PANNumber              string        `json:"panNo" logger:"sensitive"`
	OfficePhone            string        `json:"officephone" logger:"sensitive"`
	ResidencePhone         string        `json:"residencePhone" logger:"sensitive"`
	BankName2              string        `json:"bankName2" logger:"sensitive"`
	BankAccountNumber2     string        `json:"bankAccountNo2" logger:"sensitive"`
	BankBranchName2        string        `json:"bankBranchName2" logger:"sensitive"`
	BankName3              string        `json:"bankName3" logger:"sensitive"`
	BankAccountNumber3     string        `json:"bankAccountNo3" logger:"sensitive"`
	BankBranchName3        string        `json:"bankBranchName3" logger:"sensitive"`
	BankAccountNumber4     string        `json:"bankAccountNo4" logger:"sensitive"`
	BankBranchName4        string        `json:"bankBranchName4" logger:"sensitive"`
	BankName4              string        `json:"bankName4" logger:"sensitive"`
	BankAccountNumber5     string        `json:"bankAccountNo5" logger:"sensitive"`
	BankName5              string        `json:"bankName5" logger:"sensitive"`
	BankBranchName5        string        `json:"bankBranchName5" logger:"sensitive"`
	BankAccountNumber6     string        `json:"bankAccountNo6" logger:"sensitive"`
	BankName6              string        `json:"bankName6" logger:"sensitive"`
	BankBranchName6        string        `json:"bankBranchName6" logger:"sensitive"`
	NSEIpoPcode            string        `json:"nseipoPcode"`
	BSEIpoPcode            string        `json:"bseipoPcode"`
	BankAddress2           string        `json:"bankAddress2" logger:"sensitive"`
	BankAddress3           string        `json:"bankAddress3" logger:"sensitive"`
	BankAddress4           string        `json:"bankAddress4" logger:"sensitive"`
	BankAddress5           string        `json:"bankAddress5" logger:"sensitive"`
	BankAddress6           string        `json:"bankAddress6" logger:"sensitive"`
	User                   string        `json:"user" logger:"sensitive"`
	DateOfBirth            string        `json:"dobAccount" logger:"sensitive"`
	BrokerName             string        `json:"sBrokerName"`
	Depository             string        `json:"depository"`
	IfscCode1              string        `json:"ifscCode1"`
	IfscCode2              string        `json:"ifscCode2"`
	IfscCode3              string        `json:"ifscCode3"`
	IfscCode4              string        `json:"ifscCode4"`
	IfscCode5              string        `json:"ifscCode5"`
	IfscCode6              string        `json:"ifscCode6"`
	IsipMandateId          string        `json:"isipMandate"`
	MandateId              string        `json:"mandateId"`
	PoaEnabled             string        `json:"poaEnabled"`
	PoaName                string        `json:"poaName" logger:"sensitive"`
	PoaRelation            string        `json:"poaRelation"`
	PoaAddress             string        `json:"poaAddress" logger:"sensitive"`
	PoaDob                 string        `json:"poaDob" logger:"sensitive"`
	PoaCellAddress         string        `json:"poa_CellAddress" logger:"sensitive"`
	PoaEmailAddress        string        `json:"poa_EmailAddress" logger:"sensitive"`
	PoaPanNumber           string        `json:"poaPanNo" logger:"sensitive"`
	PoaType                string        `json:"poaType"`
	PoaDate                string        `json:"poaDate"`
	PoaPhoneNumber         string        `json:"poaPhoneNo" logger:"sensitive"`
	BankDetails            []BankDetails `json:"bankdetails" logger:"sensitive"`
	Product                []string      `json:"product"`
	AccountShortcodeBseIpo string        `json:"accountShortcodeBseIpo"`
	Upi                    []string      `json:"upi" logger:"sensitive"`
	StatusRemarks          string        `json:"statusRemarks"`
	SorEnabledFlag         string        `json:"sorEnabled"`
}

type BankDetails struct {
	BankName          string `json:"bankName" logger:"sensitive"`
	BankBranchName    string `json:"bankBranchName" logger:"sensitive"`
	BankAddress       string `json:"bankAddress"`
	BankAccountNumber string `json:"bankAccountNo" logger:"sensitive"`
}

type NestDefaultLoginRequest struct {
	UserId string `json:"uid" validate:"required" logger:"sensitive"`
}

type NestDefaultLoginResponse struct {
	Status                       string                    `json:"stat"`
	EnabledExchanges             []string                  `json:"exarr"`
	EnabledOrderType             []string                  `json:"orarr"`
	EnabledProductType           []string                  `json:"prarr"`
	ProductAlias                 string                    `json:"s_prdt_ali"`
	TransactionFlag              string                    `json:"sTransFlg"`
	DefaultMarketWatchlistName   string                    `json:"dwm"`
	BrokerName                   string                    `json:"brkname" logger:"sensitive"`
	BranchId                     string                    `json:"brnchid" logger:"sensitive"`
	MarketWatchCount             string                    `json:"MaxMWCount"`
	Email                        string                    `json:"email" logger:"sensitive"`
	AccountId                    string                    `json:"sAccountId" logger:"sensitive"`
	PasswordSpecialCharacterFlag string                    `json:"pwdSplChar" logger:"sensitive"`
	AccountName                  string                    `json:"accountName" logger:"sensitive"`
	UserPrivileges               string                    `json:"userPrivileges"`
	YSXExchangeFlag              string                    `json:"YSXorderEntry"`
	CriteriaAttribute            []string                  `json:"criteriaAttribute"`
	GtcGtdNLMFlag                string                    `json:"GTCGTDNLMflag"`
	ErrorMessage                 string                    `json:"Emsg,omitempty"`
	ExchangeDetail               map[string][]ExchangeInfo `json:"exchDeatil"`
}

type ExchangeInfo struct {
	ExchangeSegment string   `json:"exchseg"`
	ProductCode     []string `json:"product"`
	Exchange        string   `json:"exch"`
}

type OpenSearchResponse struct {
	ScripId             string          `json:"scripId" example:"NSE_4567654"`
	ScripToken          string          `json:"scripToken" example:"8075"`
	Group               string          `json:"group" example:"EQ"`
	ExchangeSegment     string          `json:"exchangeSegment" example:"nse_cm"`
	InstrumentType      string          `json:"instrumentType" example:"USDINR"`
	SymbolName          string          `json:"symbolName" example:"TCS"`
	TradingSymbol       string          `json:"tradingSymbol" example:"TCS-EQ"`
	OptionType          string          `json:"optionType" example:"CE"`
	UniqueKey           string          `json:"uniqueKey" example:"TCS-EQ"`
	ISIN                string          `json:"isin" example:"IN0000000000"`
	AssetCode           string          `json:"assetCode" example:"INR"`
	LotSize             uint32          `json:"lotSize" example:"1"`
	TickSize            sql.NullFloat64 `json:"tickSize" example:"0.05"`
	ExpiryDate          sql.NullString  `json:"expiryDate" example:"2021-01-01"`
	Multiplier          uint32          `json:"multiplier" example:"1"`
	DecimalPrecision    uint16          `json:"decimalPrecision" example:"2"`
	StrikePrice         sql.NullFloat64 `json:"strikePrice" example:"100"`
	DisplayStrikePrice  sql.NullFloat64 `json:"displayStrikePrice" example:"100"`
	Description         string          `json:"description" example:"TCS-EQ"`
	SubGroup            string          `json:"subGroup" example:"EQ"`
	AmcCode             string          `json:"amcCode" example:"HDFCMutualFund_EQ"`
	ContractID          string          `json:"contractID" example:"TCS-EQ-2021-01-01-CE-100"`
	CombinedSymbol      string          `json:"combinedSymbol" example:"NA"`
	ScripReserved       string          `json:"scripReserved" example:"NA"`
	ExchangeSymbolName  string          `json:"exchangeSymbolName" example:"NA"`
	RmsMarketProtection float32         `json:"rmsMarketProtection" example:"0"`
	ContractType        string          `json:"contractType" example:"Option"`
	ExchangeName        string          `json:"exchangeName" example:"NSE"`
	LastUpdatedAt       time.Time       `json:"lastUpdatedAt" example:"2021-01-01T00:00:00Z"`
	CompanyCode         uint32          `json:"companyCode" example:"1"`
	DisplayExpiryDate   sql.NullString  `json:"displayExpiryDate" example:"2021-01-01"`
	SymbolLeg2          string          `json:"symbolLeg2" example:"NA"`
	RelevanceScore      float64         `json:"relevanceScore,omitempty"`
}
