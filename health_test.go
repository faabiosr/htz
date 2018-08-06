package htz

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type (
	HealthTestSuite struct {
		suite.Suite

		assert *assert.Assertions
		mux    *http.ServeMux
		server *httptest.Server
	}
)

const (
	TestService CheckType = "test-service"
)

func (s *HealthTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.mux = http.NewServeMux()
	s.server = httptest.NewServer(s.mux)
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

func (s *HealthTestSuite) TestHttpHandlerWithWrongMethod() {
	s.mux.Handle("/health", New("app", "0.0.1"))

	req, _ := http.NewRequest(http.MethodPost, s.server.URL+"/health", nil)
	res, _ := http.DefaultClient.Do(req)

	s.assert.Equal(http.StatusNotFound, res.StatusCode)
}

func (s *HealthTestSuite) TestHttpHandler() {
	checkers := []Checker{
		func() *Check {
			return &Check{
				Name:         "test",
				Status:       true,
				ResponseTime: 2 * time.Second,
				Optional:     true,
			}
		},
	}

	s.mux.Handle("/health", New("app", "0.0.1", checkers...))

	req, _ := http.NewRequest(http.MethodGet, s.server.URL+"/health", nil)
	res, _ := http.DefaultClient.Do(req)

	var status Status

	err := json.NewDecoder(res.Body).Decode(&status)

	s.assert.NoError(err)

	s.assert.Equal(http.StatusOK, res.StatusCode)
	s.assert.Equal("app", status.Name)
	s.assert.Equal("0.0.1", status.Version)
	s.assert.True(status.Status)
	s.assert.NotEmpty(status.Date)
	s.assert.Len(status.Checks, 1)
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}
