package checker

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	CheckerTestSuite struct {
		suite.Suite

		assert *assert.Assertions
	}
)

func (s *CheckerTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
}

func TestCheckerTestSuite(t *testing.T) {
	suite.Run(t, new(CheckerTestSuite))
}
