//go:generate skygen
package models

// M struct.
// sky:some_table_name
type M struct {
	ID        string `sky:"id,pk"`
	Value     string `sky:"string_value"`
	UintValue uint64 `sky:"uint_value"`
}
