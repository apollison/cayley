package all

import (
	// supported backends
	_ "github.com/apollison/cayley/graph/kv/all"
	_ "github.com/apollison/cayley/graph/memstore"
	_ "github.com/apollison/cayley/graph/nosql/all"
	_ "github.com/apollison/cayley/graph/sql/cockroach"
	_ "github.com/apollison/cayley/graph/sql/mysql"
	_ "github.com/apollison/cayley/graph/sql/postgres"
)
