package checker

import (
	"database/sql"
	"github.com/fabiorphp/htz"
	"reflect"
	"strconv"
	"time"
)

// DB returns the status of SQL driver connection.
func DB(db *sql.DB, optional bool) htz.Checker {
	return func() *htz.Check {
		status := true
		stats := db.Stats()

		details := map[string]interface{}{
			"open_connections": strconv.Itoa(stats.OpenConnections),
		}

		start := time.Now()
		err := db.Ping()
		responseTime := time.Since(start)

		if err != nil {
			status = false
			details["error"] = err.Error()
		}

		for _, driver := range sql.Drivers() {
			dr, _ := sql.Open(driver, "")

			if reflect.TypeOf(dr) == reflect.TypeOf(db) {
				details["driver_name"] = driver
			}
		}

		return &htz.Check{
			Name:         "db",
			Type:         Datastore,
			Status:       status,
			ResponseTime: responseTime,
			Optional:     optional,
			Details:      details,
		}
	}
}
