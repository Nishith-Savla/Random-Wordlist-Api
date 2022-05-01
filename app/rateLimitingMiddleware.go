package app

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Make a map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.RWMutex

// Run a background goroutine to remove old entries from the visitors map.
func init() {
	go cleanupVisitors()
}

func getVisitor(ip string) *rate.Limiter {
	mu.RLock()

	v, exists := visitors[ip]
	if !exists {
		mu.RUnlock()
		mu.Lock()

		limiter := rate.NewLimiter(1, 2)
		// Include the current time when creating a new visitor.
		visitors[ip] = &visitor{limiter, time.Now()}

		mu.Unlock()
		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	mu.RUnlock()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 5 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}

		if limiter := getVisitor(ip); limiter.Allow() == false {
			writeJSONResponse(w, http.StatusTooManyRequests, map[string]string{"message": "received too many requests, please slow down..."})
			return
		}

		next.ServeHTTP(w, r)
	})
}
