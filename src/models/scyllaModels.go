package models

import "github.com/scylladb/gocqlx/v2/table"

// Table models.
var (
	// metadata specifies table name and columns it must be in sync with schema.
	HoldingsTable = table.Metadata{
		Name: "holdings",
		Columns: []string{
			"userId",
			"email",
			"symbol",
			"exchange",
			"qty",
			"avg_price",
		},
		PartKey: []string{
			"userId",
			"symbol",
			"exchange",
		},
		SortKey: []string{
			"qty",
		},
	}
)

// Holding represents a row in Holdings table.
// Field names are converted to camel case by default, no need to add special tags.
// A field will not be persisted by adding the `db:"-"` tag or making it unexported.
type Holding struct {
	UserId   int
	Symbol   string
	Exchange string
	Email    string
	Qty      int
	AvgPrice int
}
