package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"golang.org/x/time/rate"
	"sync"
)

// Rate Limiting - 2 requests per 5 seconds
var limiter = rate.NewLimiter(rate.Every(2*time.Second), 2) // Max 2 requests per 5 seconds

// Blacklist and Whitelist configuration
var blacklistedIPs = map[string]bool{
	"192.168.1.10": true, // Example blacklisted IP
}

var whitelistedIPs = map[string]bool{
	"127.0.0.1": true, // Localhost whitelisted
}

// Mutex for synchronization (rate limiting & IP tracking)
var mu sync.Mutex

// Simulated CAPTCHA challenge (you can replace this with real CAPTCHA service)
func captchaChallenge(w http.ResponseWriter, r *http.Request) {
	// Simulate CAPTCHA verification (Replace with a real solution)
	fmt.Fprintf(w, "Please complete the CAPTCHA challenge.")
}

// Check if IP is blacklisted
func checkIPBlacklist(r *http.Request) bool {
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	if blacklistedIPs[clientIP] {
		return true
	}
	return false
}

// Check if IP is whitelisted
func checkIPWhitelist(r *http.Request) bool {
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	if whitelistedIPs[clientIP] {
		return true
	}
	return false
}

// Rate Limiting Handler - 2 requests per 5 seconds
func rateLimitHandler(w http.ResponseWriter, r *http.Request) {
	if limiter.Allow() == false {
		http.Error(w, "Too Many Requests. Please try again later.", http.StatusTooManyRequests)
		return
	}
	fmt.Fprintf(w, "Request Accepted!")
}

// Geo-blocking for allowed regions (can be extended with real geolocation API)
func checkGeofence(r *http.Request) bool {
	// Placeholder: detect IP region (use external geolocation API)
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	region := "US" // Example region (could be fetched dynamically)

	// Restrict access to US region only (can be extended)
	if region != "US" {
		return true
	}
	return false
}

// Main Handler with DDoS Protection Logic
func handler(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limiting
	rateLimitHandler(w, r)

	// Check IP Blacklist
	if checkIPBlacklist(r) {
		http.Error(w, "Forbidden: Your IP is blacklisted", http.StatusForbidden)
		return
	}

	// Check IP Whitelist (Only allow traffic from whitelisted IPs)
	if !checkIPWhitelist(r) {
		http.Error(w, "Forbidden: Your IP is not whitelisted", http.StatusForbidden)
		return
	}

	// Check Geofencing (if the request comes from allowed region)
	if checkGeofence(r) {
		http.Error(w, "Forbidden: Your region is blocked", http.StatusForbidden)
		return
	}

	// CAPTCHA Challenge for suspicious traffic (every POST or suspicious request)
	if r.Method == "POST" { // Simulate DDoS detection
		captchaChallenge(w, r)
		return
	}

	// If everything is fine, proceed with normal processing
	fmt.Fprintf(w, "Request Accepted!")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
