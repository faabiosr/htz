package checker

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
	"net"
	"testing"
)

type (
	RedisTestSuite struct {
		CheckerTestSuite
	}
)

func (s *RedisTestSuite) TestCheckerWhenPingFailed() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	res := Redis(client, true)()

	s.assert.Equal("redis", res.Name)
	s.assert.Equal(Datastore, res.Type)
	s.assert.False(res.Status)
	s.assert.NotZero(res.ResponseTime)
	s.assert.True(res.Optional)
	s.assert.Contains(res.Details["error"], "connection refused")
}

func (s *RedisTestSuite) TestCheckerWithSuccessfulStatus() {

	if _, err := net.Dial("tcp", "localhost:6379"); err != nil {
		s.T().Skip()
	}

	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	res := Redis(client, true)()

	s.assert.Equal("redis", res.Name)
	s.assert.Equal(Datastore, res.Type)
	s.assert.True(res.Status)
	s.assert.NotZero(res.ResponseTime)
	s.assert.True(res.Optional)
	s.assert.NotZero(res.Details)
}

func TestRedisTestSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}
