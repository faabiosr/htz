package checker

import (
	"bufio"
	"github.com/fabiorphp/htz"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

// Redis returns the status of Redis server
func Redis(client *redis.Client, optional bool) htz.Checker {
	return func() *htz.Check {
		checker := &htz.Check{
			Name:     "redis",
			Type:     Datastore,
			Status:   true,
			Optional: optional,
			Details:  map[string]interface{}{},
		}

		start := time.Now()
		err := client.Ping().Err()
		checker.ResponseTime = time.Since(start)

		if err != nil {
			checker.Status = false
			checker.Details["error"] = err.Error()
		}

		info := client.Info("stats").Val()

		if info == "" {
			return checker
		}

		scanner := bufio.NewScanner(strings.NewReader(info))

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if strings.HasPrefix(line, "#") {
				continue
			}

			splited := strings.Split(line, ":")

			checker.Details[splited[0]] = splited[1]
		}

		return checker
	}
}
