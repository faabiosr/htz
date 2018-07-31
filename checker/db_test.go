package checker

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	DBTestSuite struct {
		CheckerTestSuite
	}
)

func (s *DBTestSuite) TestChecker() {
	sql.Register("mock", new(MockDriver))

	db, _ := sql.Open("mock", "user:pass@dbname")

	res := DB(db, true)()

	s.assert.Equal("db", res.Name)
	s.assert.Equal(Datastore, res.Type)
	s.assert.False(res.Status)
	s.assert.NotZero(res.ResponseTime)
	s.assert.True(res.Optional)
	s.assert.Equal("Connection error!", res.Details["error"])
	s.assert.Equal("0", res.Details["open_connections"])
	s.assert.Equal("mock", res.Details["driver_name"])
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
