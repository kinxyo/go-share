package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := simplifyUserAgent(r.UserAgent())
		ip := getIP(r)

		// Try extracting file name
		fileName := extractFileName(r.URL.Path)

		log.Printf("%s %s (%s) from %s %s",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			fileName,
			ua,
			ip,
		)

		next.ServeHTTP(w, r)
	})
}

func simplifyUserAgent(ua string) string {
	switch {
	case strings.Contains(ua, "Android"):
		return "Android"
	case strings.Contains(ua, "iPhone"):
		return "iPhone"
	case strings.Contains(ua, "iPad"):
		return "iPad"
	case strings.Contains(ua, "Windows"):
		return "Windows"
	case strings.Contains(ua, "Macintosh"):
		return "macOS"
	default:
		return "Other"
	}
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr
	// If behind a proxy, trust X-Forwarded-For
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		return fwd
	}
	return ip
}

func extractFileName(path string) string {
	if idx := strings.LastIndex(path, "/"); idx != -1 && idx < len(path)-1 {
		return path[idx+1:]
	}
	return path
}
