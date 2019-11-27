//+build cgo

package all

import (
	// backends requiring cgo
	_ "github.com/apollison/cayley/graph/sql/sqlite"
)
