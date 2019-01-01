package main

import (
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

// Map for holding informations about visitiors
var visitors = make(map[string]*visitor)
var mtx sync.Mutex

func addVisitor(ip string) *rate.Limiter {
	// 1 request per second, max 5 burst size
	limiter := rate.NewLimiter(1, 5)
	mtx.Lock()
	visitors[ip] = &visitor{limiter, time.Now()}
	mtx.Unlock()
	return limiter
}

func getVisitor(ip string) *rate.Limiter {
	mtx.Lock()
	v, exists := visitors[ip]
	if !exists {
		mtx.Unlock()
		return addVisitor(ip)
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	mtx.Unlock()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mtx.Lock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mtx.Unlock()
	}
}
