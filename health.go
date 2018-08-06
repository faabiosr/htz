package htz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	// CheckType type.
	CheckType string

	// Checker it's type function responsible for checking the service health.
	Checker func() *Check

	// Health defines the healthcheck configuration.
	Health struct {
		Name     string
		Version  string
		checkers []Checker
	}

	// Status retrieves the status of health.
	Status struct {
		Name    string    `json:"name"`
		Version string    `json:"version"`
		Status  bool      `json:"status"`
		Date    time.Time `json:"date"`
		Checks  []*Check  `json:"checks"`
	}

	// Check retrieves the health of service.
	Check struct {
		Name         string                 `json:"name"`
		Type         CheckType              `json:"type"`
		Status       bool                   `json:"status"`
		ResponseTime time.Duration          `json:"response_time"`
		Optional     bool                   `json:"optional"`
		Details      map[string]interface{} `json:"details"`
	}
)

// New retrieves the new instance of Health.
func New(name, version string, checkers ...Checker) *Health {
	return &Health{
		Name:     name,
		Version:  version,
		checkers: checkers,
	}
}

// Check checks the status of services.
func (h *Health) Check() *Status {
	status := &Status{
		Name:    h.Name,
		Version: h.Version,
		Status:  true,
		Date:    time.Now(),
	}

	numCheckers := len(h.checkers)

	if numCheckers == 0 {
		return status
	}

	ch := make(chan *Check, numCheckers)

	for _, ckr := range h.checkers {

		go func(ch chan<- *Check, ckr Checker) {
			status := ckr()

			ch <- status
		}(ch, ckr)
	}

	for i := 0; i < numCheckers; i++ {
		check := <-ch
		status.Checks = append(status.Checks, check)

		if check.Status == false && check.Optional == false {
			status.Status = false
		}

		if check.Details == nil {
			check.Details = make(map[string]interface{})
		}
	}

	return status
}

// MarshalJSON marshals CheckType instance into string.
func (ct CheckType) MarshalJSON() ([]byte, error) {
	if len(ct) > 0 {
		return []byte(fmt.Sprintf(`"%s"`, ct)), nil
	}

	return []byte(`"unknown"`), nil
}

// ServerHTTP retrieves the health status.
func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(h.Check())
}
