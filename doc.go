// Package htz show the health of your app.
//
// Example Usage
//
// The following is a simple example using the library:
//	package main
//
//	import (
//		"encoding/json"
//		"fmt"
//		"github.com/faabiosr/htz"
//		"time"
//	)
//
//	func main() {
//		checkers := []htz.Checker{
//			func() *htz.Check {
//				return &htz.Check{
//					Name:         "some-api",
//					Type:         "internal-service",
//					Status:       false,
//					ResponseTime: 6 * time.Second,
//					Optional:     false,
//					Details:      map[string]interface{}{
//					    "url": "internal-service.api",
//				    },
//				}
//			},
//		}
//
//		h := htz.New("my-app", "0.0.1", checkers...)
//		res, _ := json.MarshalIndent(h.Check(), "", "  ")
//
//		fmt.Println(string(res))
//	}
//
// Using with a HTTP Server:
//	package main
//
//	import (
//		"github.com/faabiosr/htz"
//		"net/http"
//		"time"
//	)
//
//	func main() {
//		checkers := []htz.Checker{
//			func() *htz.Check {
//				return &htz.Check{
//					Name:         "some-api",
//					Status:       false,
//					ResponseTime: 6 * time.Second,
//					Optional:     false,
//				}
//			},
//		}
//
//		h := htz.New("my-app", "0.0.1", checkers...)
//
//		http.Handle("/htz", h)
//
//		http.ListenAndServe(":8080", nil)
//	}
package htz
