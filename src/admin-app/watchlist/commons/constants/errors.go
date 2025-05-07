package constants

// BFF Watchlists Errors
const (
	GetPredefinedWatchlistResponseError         = "failed to fetch predefined watchlist"
	GetPredefinedWatchlistResponseTimeoutError  = "failed to retrieve predefined watchlist due to timeout"
	AdgToMultiWatchlistLengthError              = "maximum number of operations allowed are 2"
	AdgToMultiWatchlistActionTypeError          = "in case of multiple operations, only allowed operations are ADD and DEL"
	AdgToMultiWatchlistDuplicateActionTypeError = "duplicate action type %s found"
	AdgToMultiWatchlistScripIdWatchlistError    = "WatchlistIds must be different for 'ADD' and 'DEL'"
)

// Watchlist Database Errors
const (
	WatchlistExistError                     = "watchlist already exist with name: %s, please try another name"
	PredefinedScripsNotFoundError           = "predefined scrips not found"
	NoScripsFoundError                      = "no scrips found in watchlist"
	InvalidScripsError                      = "all provided scrips are invalid "
	FailedToAddScripsError                  = "failed while adding scrips to watchlist: %w"
	InvalidWatchlistIdError                 = "invaild watchlist id provided"
	FailedToCheckWatchlistIdError           = "failed to check watchlist ID: %w"
	ScripForeignKeyError                    = "invalid scrips provided"
	ScripUniqueKeyError                     = "all the given scrips already exists in watchlist"
	InvalidScripError                       = "invalid scrip provided"
	ScripLimitExceededForAllWatchlistsError = "scrip limit exceeded for all watchlists"
	ScripLimitExceededForFewWatchlistsError = "scrip limit exceeded for few watchlists"
	AllWatchlistIdsInvalidError             = "all the watchlists are invalid"
	FewWatchlistIdsInvalidError             = "few watchlist IDs are invalid"
)

// Watchlist BFF Error Messages
const (
	WatchlistScripLimitError        = "watchlist scrip limit exceeded, can not add scrips/stocks more than %d"
	WatchlistLimitError             = "watchlist limit exceeded, can not add more than %d watchlists"
	BFFWatchlistNameValidationError = "watchlist name must ne alphanumeric, should not start with special caracter or numeric value"
	WatchlistNameRegexNotFound      = "watchlist name regex not found"
)

const (
	JSONBindingFailedError = "json binding failed"
	ValidationFailedError  = "validation failed"
)

const (
	NoUserIdFoundError   = "user not found"
	NoBrokerIdFoundError = "broker not found"
)

const (
	BothWatchlistError = "both watchlist empty"
)

//Datafetching Error
const (
	UserIdFetchingError   = "error fetching user-defined watchlists: %w"
	BrokerIdFetchingError = "error fetching pre-defined watchlists: %w"
)

const (
	SongsFetchingError               = "error while fetching songs"
	InvalidSongIDsError              = "invalid Song Id"
	DuplicatePlaylistError           = "playlist Already Created"
	DuplicateKeyViolationError       = "duplicate key value violates unique constraint"
	InvalidInputError                = "invalid input error"
	InvalidUserID                    = "invalid user id"
	CreatePlaylistRepositoryNilError = "createPlaylistRepository cannot be nil"
	PlaylistNotFoundError            = "playlist not found"
	PlaylistModificationError        = "playlist modification error"
)

const (
	UserIDRequiredError        = "user_id is required"
	SongConditionRequiredError = "song_condition is required"
	PlaylistDataRequiredError  = "playlist_data is required"
	PlaylistNameRequiredError  = "playlist_name is required"
)

const (
	InvalidActionChoice = "invalid action choice"
	ServiceFailedToCreatePlaylist = "Service failed to create playlist"
)
