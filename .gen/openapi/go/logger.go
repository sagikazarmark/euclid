/*
 * Euclid API
 *
 * Sequential ID generator
 *
 * API version: 1.0.0
 * Contact: mark.sagikazar+euclid@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
