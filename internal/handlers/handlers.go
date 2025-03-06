package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/amjadfqs/go-htmx-starter/templates/pages"
)

// HomeHandler renders the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	pages.Home().Render(r.Context(), w)
}

// AboutHandler renders the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	pages.About().Render(r.Context(), w)
}

// ExampleHtmxHandler demonstrates a simple HTMX response
func ExampleHtmxHandler(w http.ResponseWriter, r *http.Request) {
	// Add HTMX-specific headers
	w.Header().Set("HX-Trigger", `{"showToast": "Content loaded successfully"}`)

	// Simulate processing time (optional)
	time.Sleep(200 * time.Millisecond)

	// Check if the request wants JSON instead of HTML
	if r.Header.Get("Accept") == "application/json" {
		response := map[string]interface{}{
			"success":   true,
			"message":   "Data loaded successfully",
			"timestamp": time.Now().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// For regular HTMX requests, render the HTML component
	pages.HtmxExample().Render(r.Context(), w)
}

// ClientInfoHandler provides detailed information about the client
func ClientInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Get client IP address
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = r.RemoteAddr
	}

	// Get user agent
	userAgent := r.Header.Get("User-Agent")

	// Get accept language
	acceptLanguage := r.Header.Get("Accept-Language")

	// Get request method and path
	requestMethod := r.Method
	requestPath := r.URL.Path

	// Get referer if available
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "Direct visit (no referer)"
	}

	// Simple device type detection (very basic)
	deviceType := "Unknown"
	if containsAny(userAgent, []string{"Mobile", "Android", "iPhone", "iPad"}) {
		deviceType = "Mobile"
		if containsAny(userAgent, []string{"iPad", "Tablet"}) {
			deviceType = "Tablet"
		}
	} else if containsAny(userAgent, []string{"Windows", "Macintosh", "Linux", "X11"}) {
		deviceType = "Desktop"
	}

	// Create response data
	clientInfo := map[string]string{
		"ip_address":      clientIP,
		"user_agent":      userAgent,
		"device_type":     deviceType,
		"accept_language": acceptLanguage,
		"request_method":  requestMethod,
		"request_path":    requestPath,
		"referer":         referer,
		"timestamp":       time.Now().Format(time.RFC3339),
	}

	// Determine response format based on Accept header
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clientInfo)
		return
	}

	// For HTMX requests or regular HTML requests
	w.Header().Set("Content-Type", "text/html")

	// Build HTML response
	html := "<div class='client-info-container'>"
	html += "<h3>Client Information</h3>"
	html += "<table class='client-info-table'>"
	html += "<tr><th>Property</th><th>Value</th></tr>"

	for key, value := range clientInfo {
		html += "<tr>"
		html += "<td>" + formatLabel(key) + "</td>"
		html += "<td>" + value + "</td>"
		html += "</tr>"
	}

	html += "</table></div>"

	w.Write([]byte(html))
}

// Helper function to format label text
func formatLabel(s string) string {
	// Replace underscores with spaces and capitalize each word
	words := strings.Split(s, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// Helper function to check if string contains any of the given substrings
func containsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}
