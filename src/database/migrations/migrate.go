package migrations

import (
	"encoding/csv"
	"fmt"
	"omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"os"
	"strconv"
)

// AutoMigrate performs automatic table migrations for models.
func AutoMigrate(dbConnectionClient *genericModels.DBConnectionClient) error {

	// Add your models
	//We will uncomment the specific model for which we want to migrate the data in the database
	models := []interface{}{
		// &genericModels.UsersInfo{},
		// &genericModels.ScripMaster{},
		// &genericModels.Watchlists{},
		// &genericModels.BrokerWatchlists{},
		// &genericModels.Brokers{},
		// &genericModels.WatchlistScrips{},
		// &genericModels.Devices{},
		// &genericModels.User{},
		// &genericModels.Broker{},
		// &genericModels.StaticExchangeConfig{},
		// &genericModels.CompanyMaster{},
		// &genericModels.CashFlow{},
		// &genericModels.BalanceSheet{},
		// &genericModels.ProfitAndLoss{},
		// &genericModels.CompanyYearlyRatios{},
		// &genericModels.CmotsEtlTracker{},
		// &genericModels.CompanyProfile{},
		// &genericModels.ShareHoldingPattern{},
		// &genericModels.Basket{},
		// &genericModels.BasketOrder{},
	}

	// Auto-migrate models to postgres
	if err := dbConnectionClient.GormDb.AutoMigrate(models...); err != nil {
		return err
	}

	// err := genericModels.InitDBConstraints(cockroachDBClient.GormDb)
	// if err != nil {
	// 	return err
	// }

	// err := genericModels.InitDBIndexes(cockroachDBClient.GormDb)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// LoadExchangeConfigs loads the exchange configurations from a CSV file and migrates the data to the database.
// This needs to be done when we move to any new environment.
func LoadExchangeConfigs(dbConnectionClient *genericModels.DBConnectionClient) error {
	// Open the CSV file
	csvPath := constants.StaticExchangeConfigFilePath
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf(constants.OpenCsvFileError, err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf(constants.ReadCsvFileError, err)
	}

	// Prepare a slice to hold the data
	var configs []genericModels.StaticExchangeConfig

	// Iterate over the records (skip header row)
	for i, record := range records {
		if i == 0 {
			continue
		}

		advRetTypeLicensed, _ := strconv.ParseBool(record[4])
		spreadIndicator, _ := strconv.ParseBool(record[5])
		auctionIndicator, _ := strconv.ParseBool(record[6])

		// Append the record to the configs slice
		config := genericModels.StaticExchangeConfig{
			Exchange:                      record[1],
			PriceType:                     record[2],
			RetentionType:                 record[3],
			AdvancedRetentionTypeLicensed: advRetTypeLicensed,
			SpreadIndicator:               spreadIndicator,
			AuctionIndicator:              auctionIndicator,
		}
		configs = append(configs, config)
	}

	query := fmt.Sprintf("TRUNCATE TABLE %s", constants.StaticExchangeConfigTables)

	// Truncate the table before inserting new records
	if err := dbConnectionClient.GormDb.Exec(query).Error; err != nil {
		return fmt.Errorf(constants.PostgresFailedToTruncateTable)
	}

	// Insert the records into the database
	if err := dbConnectionClient.GormDb.Create(&configs).Error; err != nil {
		return fmt.Errorf(constants.PostgresExchangeConfigDataMigrationError)
	}

	return nil
}
