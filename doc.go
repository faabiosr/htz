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
//		"github.com/fabiorphp/htz"
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
package htz
