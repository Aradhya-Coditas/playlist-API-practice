 # Postgres Package

This package provides functionalities for connecting to a PostgreSQL database, including both actual and mock connections. Below is a detailed overview of the functions and their usage. 

## Functions

1. `InitPostgresDBConfig(ctx context.Context) error`

   - **Description**: Initializes the PostgreSQL database configuration by reading environment variables and establishing a connection.
   - **Parameters**:
     - `ctx`: The context object.
   - **Returns**:
     - `error`: An error if initialization fails.

2. `ConnectPostgresDatabase(ctx context.Context, postgresConfig models.PostgresConfig, log logger.Logger) error`

   - **Description**: Establishes a connection to the PostgreSQL database.
   - **Parameters**:
     - `ctx`: The context object.
     - `postgresConfig`: An instance of `models.PostgresConfig` containing the database configuration.
     - `log`: The logger object for logging messages.
   - **Returns**:
     - `error`: An error if connection establishment fails.

3. `ConnectMockDatabase(ctx context.Context, postgresConfig models.PostgresConfig, log logger.Logger) error`

   - **Description**: Establishes a mock connection to the PostgreSQL database when mock is set to true.
   - **Parameters**:
     - `ctx`: The context object.
     - `postgresConfig`: An instance of `models.PostgresConfig` containing the database configuration.
     - `log`: The logger object for logging messages.
   - **Returns**:
     - `error`: An error if mock connection establishment fails.

4. `SetPostgresClient(db *gorm.DB, sqlDB *sql.DB)`

   - **Description**: Sets the PostgreSQL client with the provided database and SQL connection parameters.
   - **Parameters**:
     - `db`: The GORM database object.
     - `sqlDB`: The SQL database object.
   - **Returns**: None.

5. `ClosePostgres(ctx context.Context)`

   - **Description**: Closes the PostgreSQL client connection.
   - **Parameters**:
     - `ctx`: The context object.
   - **Returns**: None.

6. `GetPostGresClient() *PostGresConClient`

   - **Description**: Retrieves the PostgreSQL client.
   - **Parameters**: None.
   - **Returns**:
     - `*PostGresConClient`: The PostgreSQL client object.

7. `getEnv(key, defaultValue string) string`

   - **Description**: Retrieves the environment variable value with the provided key and fallback to the default value if not found.
   - **Parameters**:
     - `key`: The key of the environment variable.
     - `defaultValue`: The default value if the environment variable is not found.
   - **Returns**:
     - `string`: The value of the environment variable or the default value.

## Initialize Postgres connection

To initialize the postgres connection in main file.
1.  Import postgres library 
2.  Call InitPostgresDBConfig() function

```go 
import (
    	"context"
        "omnenest-backend/src/postgres"
)

func main() {
	ctx := context.Background()
	err := postgres.InitPostgresDBConfig(ctx)
	if err != nil {
		panic(err)
	}
	defer postgres.ClosePostgres(ctx)
}
```

# helper functions 

In helper.go file we have defined various db operations provided by the PostgreSQL package and used them for database operations.

Below are few helper functions for performing various database operations provided by the PostgreSQL package.

```go
// ReadAllRecords retrieves all records from the PostgreSQL database.
func (client *PostGresConClient) ReadAllRecords(record interface{}) *gorm.DB

// UpdateRecordWithCondition updates an existing record in the PostgreSQL database with conditions.
func (client *PostGresConClient) UpdateRecordWithCondition(record interface{}, conditions map[string]interface{}, updates map[string]interface{}) *gorm.DB

// DeleteRecordByID deletes a record from the PostgreSQL database by its ID.
func (client *PostGresConClient) DeleteRecordByID(record interface{}, id uint) *gorm.DB
```

## Use helper functions

```go 
//call GetPostGresClient to perform database operations 
db := postgres.GetPostGresClient().GormDb
// columns you want to update in database table 
updateData := map[string]interface{}{
   database_columnname1: updatedvalue,
   database_columnname2: true,
}
// conditions to be satisfied
conditions := map[string]interface{}{
    database_columname2: condition_to_be_matched_colum, // example: userid  (where in sql)
}
// call the helped functions to perform the operation with parameters and conditons 
err := service.unblockUserRegistry.UpdateRecordWithCondition(db, &user, conditions, updateData)
```

