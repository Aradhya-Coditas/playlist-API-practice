package constants

// Watchlist NEST API URL Keys
const (
	ServiceName      = "watchlist"
	PortDefaultValue = 9098
)

// Watchlist SQL Query Entry Keys
const (
	WatchlistName       = "watchlist_name"
	ExchangeSegement    = "exchange_segment"
	TradingSymbol       = "trading_symbol"
	ScripToken          = "scrip_token"
	Exchange            = "exchange"
	SymbolName          = "symbol_name"
	UniqueKey           = "unique_key"
	DisplayExpiryDate   = "display_expiry_date"
	ExpiryDate          = "expiry_date"
	StrikePrice         = "strike_price"
	DecimalPrecision    = "decimal_precision"
	InstrumentType      = "instrument_type"
	WatchlistId         = "watchlist_id"
	UserId              = "user_id"
	BrokerId            = "Broker_id"
	LastUpdatedAt       = "last_updated_at"
	ScripCount          = "scrip_count"
	WatchlistScripLimit = "watchlist_scrip_limit"
	Order               = "order"
	NewWatchlistNameKey = "newWatchlistName"
	ScripId             = "scrip_id"
	Id                  = "id"
)

// Watchlist SQL Queries
const (
	ScripTokenFilterQuery                   = "(exchange_segment = ? AND scrip_token IN (?))"
	ORConditionOperator                     = " OR "
	WatchlistScripsInnerJoinQuery           = "JOIN watchlist_scrips ON watchlist_scrips.scrip_id = scrip_master.id"
	WatchlistsInnerJoinQuery                = "JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id"
	WatchlistIdAndUserIdCondition           = "watchlists.user_id = ? AND watchlists.id = ?"
	WatchlistIDAndUserIdCondition           = "watchlists.user_id = ? AND watchlists.id = ?"
	ValidScripsCondition                    = "id IN ? AND id NOT IN (?)"
	CountAndWatchlistExistenceQuery         = "COUNT(*) as count, EXISTS (SELECT 1 FROM watchlists WHERE watchlists.id = ?) as watchlist_exists"
	WatchlistScripOrderParameter            = "watchlist_scrips.order"
	GetMultiWatchlistsInnerJoinQuery        = "INNER JOIN watchlists ON watchlists.id = watchlist_scrips.watchlist_id AND watchlists.user_id = ?"
	GetScripCountPerWatchlistSelectQuery    = "watchlists.id AS watchlist_id, COUNT(watchlist_scrips.id) AS count"
	GetMaxOrderPerWatchlistSelectQuery      = "watchlist_scrips.watchlist_id AS watchlist_id, COALESCE(MAX(watchlist_scrips.order), 0) AS max_order"
	GetMaxOrderPerWatchlistInnerJoinQuery   = "INNER JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id"
	GetMaxOrderByWatchlistCondition         = "watchlists.user_id = ? AND watchlists.id IN ?"
	GetScripCountPerWatchlistInnerJoinQuery = "LEFT JOIN watchlist_scrips ON watchlists.id = watchlist_scrips.watchlist_id"
	QueryWithMultipleIds                    = "id IN ?"
	GetScripCountByWatchlistCondition       = "watchlists.user_id = ? AND watchlists.id IN ?"
	DeleteWatchlistScripsCondition          = "watchlist_scrips.watchlist_id = ?"
	UpdateScripCountCaseQuery               = "scrip_count = CASE id"
	UpdateLastUpdatedAtCaseQuery            = "last_updated_at = CASE id"
	UpdateScripCountCaseConditionQuery      = " WHEN %d THEN %d"
	UpdateLastUpdatedAtCaseConditionQuery   = " WHEN %d THEN CURRENT_TIMESTAMP"
	UpdateWatchlistCondition                = "UPDATE watchlists SET %s, %s WHERE id IN ?"
	UpdateWatchlistCaseEndQuery             = " END"
	GetWatchlistsContainingScripCondition   = "watchlist_scrips.scrip_id = ?"
	WatchlistScripScripIdCondition          = "watchlist_scrips.scrip_id"
	WatchlistUserIdCondition                = "watchlists.user_id"
	WatchlistIdParameter                    = "watchlists.id"
	GetWatchlistScripsInnerJoinQuery        = "JOIN scrip_master ON scrip_master.id = watchlist_scrips.scrip_id"
)

// Scrip Group and Subgroup Separator
const (
	PredefinedScripsSubGroupSeparator  = "|"
	PredefinedScripGroupSeparator      = ","
	QueryToCheckAllScripsExistInMaster = "id IN ?"
)

const (
	WatchlistScripIdExistWarningText = "few scrips were not added in the watchlist"
	PGUniqueKeyErrorCode             = "23505"
	PGForeignKeyErrorCode            = "23503"
)

const (
	CreateWatchlistCountOnUserIdQuery                       = `SELECT count(*) FROM "watchlists" WHERE "user_id" = $1`
	CreateWatchlistInsertWatchlistDetailsQuery              = `INSERT INTO "watchlists" ("user_id","watchlist_name","last_updated_at","scrip_count") VALUES ($1,$2,$3,$4) RETURNING "id"`
	CreateWatchlistGetWatchlistDetailsQuery                 = `SELECT * FROM "watchlists" WHERE "id" = $1 AND "user_id" = $2 ORDER BY "watchlists"."id" LIMIT $3`
	CreateWatchlistGetTotalCountOfScripsQuery               = `SELECT COUNT(*) as count, EXISTS (SELECT 1 FROM watchlists WHERE watchlists.id = $1) as watchlist_exists FROM "watchlist_scrips" WHERE "watchlist_id" = $2`
	CreateWatchlistInsertWatchlistScripDetailsQuery         = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3) RETURNING "id"`
	CreateWatchlistUpdateScripCountQuery                    = `UPDATE "watchlists" SET "last_updated_at"=$1,"scrip_count"=$2 WHERE "id" = $3`
	CreateWatchlistBulkInsertScripsQuery                    = `SELECT * FROM "scrip_master" WHERE id IN ($1,$2)`
	WatchlistCreatedSavepoint                               = `SAVEPOINT watchlistCreatedSavePoint`
	CreateWatchlistCountAllInWatchlistQuery                 = `SELECT count\(\*\) FROM "watchlists"`
	WatchlistCreatedRollBack                                = `ROLLBACK TO SAVEPOINT watchlistCreatedSavePoint`
	CreateWatchlistInsertWatchlistScripMultipleDetailsQuery = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3),($4,$5,$6) RETURNING "id"`
	CreateWatchlistGetValidScripsQuery                      = `SELECT * FROM "scrip_master" LEFT JOIN (SELECT "scrip_id" FROM "watchlist_scrips" JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id WHERE "user_id" = $1 AND "watchlist_id" = $2) temp ON scrip_master.id = temp.scrip_id WHERE scrip_master.id IN ($3) AND temp.scrip_id IS NULL`
	CreateWatchSingleSelectScripQuery                       = `SELECT * FROM "scrip_master" WHERE id IN ($1)`
	ScripMasterLeftJoinSubqueryQuery                        = "LEFT JOIN (?) temp ON scrip_master.id = temp.scrip_id"
	ScripMasterNotInWatchlistCondition                      = "temp.scrip_id IS NULL"
	ScripMasterValidScripsCondition                         = "scrip_master.id IN (?)"
)

// Watchlist SQL Queries for test
const (
	DeleteWatchlistScripsQuery            = `DELETE FROM watchlist_scrips USING watchlists WHERE watchlist_scrips.watchlist_id = watchlists.id AND watchlists.user_id = $1 AND watchlist_scrips.watchlist_id = $2`
	DeleteWatchlistQuery                  = `DELETE FROM "watchlists" WHERE "id" = $1 AND "user_id" = $2`
	UpdateWatchlistQuery                  = `UPDATE "watchlists" SET "last_updated_at"=$1,"watchlist_name"=$2 WHERE "id" = $3 AND "user_id" = $4`
	UpdateScripCountQuery                 = `UPDATE "watchlists" SET "last_updated_at"=$1,"scrip_count"=$2 WHERE "id" = $3 AND "user_id" = $4`
	InsertWatchlistScripQuery             = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3) RETURNING "id"`
	InsertMultipleWatchlistScripsQuery    = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3),($4,$5,$6) RETURNING "id"`
	ValidateScripQuery                    = `SELECT * FROM "scrip_master" WHERE id IN ($1)`
	ValidateMultipleScripQuery            = `SELECT * FROM "scrip_master" WHERE id IN ($1,$2)`
	SelectWatchlistQuery                  = `SELECT "id","watchlist_name","last_updated_at" FROM "watchlists" WHERE "user_id" = $1`
	DeleteScripFromMultiWatchlistsQuery   = `DELETE FROM watchlist_scrips ws USING watchlists w WHERE ws.watchlist_id = w.id AND w.user_id = ? AND w.id = ANY(?) AND ws.scrip_id = ?`
	DeleteWatchlistAndScripsQuery         = `DELETE FROM watchlist_scrips USING watchlists WHERE watchlist_scrips.watchlist_id = watchlists.id AND watchlists.user_id = ? AND watchlist_scrips.watchlist_id = ?`
	AdgToMultiWatchlistSelectQuery        = `SELECT watchlists.id AS watchlist_id, COUNT(watchlist_scrips.id) AS count    FROM "watchlists"  LEFT JOIN watchlist_scrips ON watchlists.id = watchlist_scrips.watchlist_id  WHERE watchlists.user_id = $1 AND watchlists.id IN ($2,$3)  GROUP BY "watchlists"."id"`
	AdgToMultiWatchlistOneIdSelectQuery   = `SELECT watchlists.id AS watchlist_id, COUNT(watchlist_scrips.id) AS count FROM "watchlists" LEFT JOIN watchlist_scrips ON watchlists.id = watchlist_scrips.watchlist_id WHERE watchlists.user_id = $1 AND watchlists.id IN ($2) GROUP BY "watchlists"."id"`
	AdgToMultiWatchlistOneIdGetQuery      = `SELECT watchlists.id,watchlists.watchlist_name,watchlists.last_updated_at   FROM "watchlist_scrips"   INNER JOIN watchlists   ON watchlists.id = watchlist_scrips.watchlist_id AND watchlists.user_id = $1   WHERE watchlist_scrips.scrip_id = $2`
	AdgToMultiWatchlistMaxOrderQuery      = `SELECT watchlist_scrips.watchlist_id AS watchlist_id, COALESCE(MAX(watchlist_scrips.order), 0) AS max_order FROM "watchlist_scrips" INNER JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id WHERE watchlists.user_id = $1 AND watchlists.id IN ($2,$3) GROUP BY "watchlist_id"`
	AdgToMultiWatchlistInsertQuery        = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3),($4,$5,$6) RETURNING "id"`
	AdgToMultiWatchlistInsertSingleQuery  = `INSERT INTO "watchlist_scrips" ("watchlist_id","scrip_id","order") VALUES ($1,$2,$3) RETURNING "id"`
	AdgToMultiWatchlistUpdateQuery        = `UPDATE watchlists SET scrip_count = CASE id   WHEN %[1]d THEN 2  WHEN %[2]d THEN 2  END, last_updated_at = CASE id WHEN %[1]d THEN CURRENT_TIMESTAMP    WHEN %[2]d THEN CURRENT_TIMESTAMP  END   WHERE id IN ($1,$2)`
	AdgToMultiWatchlistInnerJoinQuery     = `SELECT watchlists.id,watchlists.watchlist_name,watchlists.last_updated_at FROM "watchlist_scrips" INNER JOIN watchlists ON watchlists.id = watchlist_scrips.watchlist_id AND watchlists.user_id = $1 WHERE watchlist_scrips.scrip_id = $2`
	AdgToMultiWatchlistDeleteQuery        = `DELETE FROM watchlist_scrips ws USING watchlists w WHERE ws.watchlist_id = w.id AND w.user_id = $1 AND w.id = ANY($2) AND ws.scrip_id = $3`
	AdgToMultiWatchlistIdDeleteQuery      = `DELETE FROM watchlist_scrips ws USING watchlists w WHERE ws.watchlist_id = w.id AND w.user_id = $1 AND w.id = ANY($2) AND ws.scrip_id = $3`
	AdgToMultiWatchlistIdUpdateQuery      = `UPDATE watchlists SET scrip_count = CASE id WHEN %d THEN 0 END, last_updated_at = CASE id WHEN %d THEN CURRENT_TIMESTAMP END WHERE id IN ($1)`
	AdgToMultiWatchlistOneIdUpdateQuery   = `UPDATE watchlists SET scrip_count = CASE id WHEN %[1]d THEN 2 END, last_updated_at = CASE id WHEN %[1]d THEN CURRENT_TIMESTAMP END WHERE id IN ($1)`
	AdgToMultiWatchlistOneUpdateQuery     = `UPDATE watchlists SET scrip_count = CASE id WHEN %[1]d THEN 1 END, last_updated_at = CASE id WHEN %[1]d THEN CURRENT_TIMESTAMP END WHERE id IN ($1)`
	AdgToMultiWatchlistOneIdMaxOrderQuery = `SELECT watchlist_scrips.watchlist_id AS watchlist_id, COALESCE(MAX(watchlist_scrips.order), 0) AS max_order FROM "watchlist_scrips" INNER JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id WHERE watchlists.user_id = $1 AND watchlists.id IN ($2) GROUP BY "watchlist_id"`
	GetScripMasterDetailsQuery            = `SELECT "exchange_segment","trading_symbol","scrip_token","exchange","instrument_type","symbol_name","unique_key","display_expiry_date","strike_price","decimal_precision" FROM "scrip_master" JOIN watchlist_scrips ON watchlist_scrips.scrip_id = scrip_master.id JOIN watchlists ON watchlist_scrips.watchlist_id = watchlists.id WHERE watchlists.user_id = $1 AND watchlists.id = $2 ORDER BY watchlist_scrips.order`
	GetScripMasterDetailsByGroupQuery     = `SELECT "exchange_segment","trading_symbol","scrip_token","exchange","instrument_type","symbol_name","unique_key","display_expiry_date","strike_price","decimal_precision" FROM "scrip_master" WHERE (exchange_segment = $1 AND scrip_token IN ($2,$3,$4,$5,$6))`
)

var GetWatchlistsContainingScripColumns = []string{
	"watchlists.id",
	"watchlists.watchlist_name",
	"watchlists.last_updated_at",
}

const (
	BrokerIdJoinQuery  = "JOIN brokers ON users_info.broker_id = brokers.id"
	BrokerIdWhereQuery = "users_info.id = ? AND brokers.id = ?"
)

const (
	UserIDKey   = "user_id"
	BrokerIDKey = "broker_id"
)

const (
	Request = "request"
)

// Success messages
const (
	GetWatchlistScripsSuccessMessage = "GetWatchlistScripsQuery executed successfully"
)

const (
	CheckUserIdAPILog     = "CheckUserIdAPI"
	CheckBrokerIdAPILog   = "CheckBrokerIdAPI"
	UserWatchlistAPILog   = "UserWatchlistAPI"
	BrokerWatchlistAPILog = "BrokerWatchlistAPI"
	WatchlistAPILog       = "WatchlistAPI"
)

const (
	User_id       = "user_id"
	Song_ids       = "song_ids"
	SongCondition = "song_condition"
	PlaylistData  = "playlist_data"
	PlaylistSongs = "playlist_songs"
	Name          = "name"
	Song          = "song"
	Playlist_id   = "playlist_id"
	Song_id       = "song_id"
	ID_IN         = "id IN"
	Description   = "description"
	
)

const (
	SuccessfullyFetchedWatchlists = "Successfully fetched watchlists"
)

const (
	SuccessfullyCreatedPlaylist = "Playlist created successfully"
)

const (
	PlaylistName          = "playlist Name"
	PlaylistDescription   = "playlist Description"
	PlaylistCreationError = "playlist Creation Error"
)

const (
	PlaylistAPILog = "PlaylistAPILog"
)

const (
	UsersTable      = "users"
	PlaylistsTable = "playlists"
	SongsTable = "songs"
)

const (
	SuccessfullyModifiedPlaylist = "successfully modified playlist"
	SuccessfullyAddedSongInPlaylist = "successfully Added songs in playlist"
	SuccessfullyDeletedSongFromPlaylist = "successfully Deleted songs from playlist"
)

const (
	ActionAdd = "add"
	ActionDelete = "delete"
)
