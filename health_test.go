package htz

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type (
	HealthTestSuite struct {
		suite.Suite

		assert *assert.Assertions
	}
)

const (
	TestService CheckType = "test-service"
)

func (s *HealthTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
}

func (s *HealthTestSuite) TestFactory() {
	s.assert.IsType(new(Health), New("app", "0.0.1"))
}

func (s *HealthTestSuite) TestWithoutCheckers() {
	status := New("app", "0.1.0").Check()

	s.assert.Equal("app", status.Name)
	s.assert.Equal("0.1.0", status.Version)
	s.assert.True(status.Status)
	s.assert.NotEmpty(status.Date)
	s.assert.Len(status.Checks, 0)
}

func (s *HealthTestSuite) TestWithCheckers() {
	checkers := []Checker{
		func() *Check {
			return &Check{
				Name:         "test-1",
				Type:         TestService,
				Status:       false,
				ResponseTime: 60 * time.Second,
				Optional:     false,
			}
		},
		func() *Check {
			return &Check{
				Name:         "test-2",
				Status:       false,
				ResponseTime: 60 * time.Second,
				Optional:     false,
			}
		},
	}

	status := New("app", "0.1.0", checkers...).Check()

	s.assert.Equal("app", status.Name)
	s.assert.False(status.Status)
	s.assert.NotEmpty(status.Date)
	s.assert.Len(status.Checks, 2)

	res, _ := json.Marshal(status)

	s.assert.Contains(string(res), "test-service")
	s.assert.Contains(string(res), "unknown")
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}
