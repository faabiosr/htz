package checker

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	RuntimeTestSuite struct {
		CheckerTestSuite
	}
)

func (s *RuntimeTestSuite) TestChecker() {
	res := Runtime(true)()

	s.assert.Equal("runtime", res.Name)
	s.assert.Equal(App, res.Type)
	s.assert.True(res.Status)
	s.assert.Zero(res.ResponseTime)
	s.assert.True(res.Optional)
	s.assert.NotZero(res.Details["goroutines"])
	s.assert.NotZero(res.Details["memory_total_alloc"])
	s.assert.NotZero(res.Details["memory_heap_alloc"])
}

func TestRuntimeTestSuite(t *testing.T) {
	suite.Run(t, new(RuntimeTestSuite))
}
