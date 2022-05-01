package app

import (
	"crypto/subtle"
	"net/http"
)

func authorizationHandler(apiKey string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKeyHeader := []byte(r.Header.Get("X-API-KEY"))
			apiKeyBytes := []byte(apiKey)

			if subtle.ConstantTimeEq(int32(len(apiKeyHeader)), int32(len(apiKeyBytes))) == 0 ||
				subtle.ConstantTimeCompare(apiKeyHeader, apiKeyBytes) == 0 {
				writeJSONResponse(w, 401, map[string]string{"message": "unauthorized"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
