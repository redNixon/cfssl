// +build postgresql

package sql

import (
	"testing"

	"github.com/redNixon/cfssl/certdb/testdb"
)

func TestPostgreSQL(t *testing.T) {
	db := testdb.PostgreSQLDB()
	ta := TestAccessor{
		Accessor: NewAccessor(db),
		DB:       db,
	}
	testEverything(ta, t)
}
