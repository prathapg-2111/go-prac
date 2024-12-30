package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time,omitempty"`
}

type MultiTimeResponse map[string]string

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Get timezone parameter from URL
	tzParam := r.URL.Query().Get("tz")

	// If no timezone specified, return UTC time
	if tzParam == "" {
		currentTime := time.Now().UTC().Format("2006-01-02 15:04:05 -0700 MST")
		response := TimeResponse{CurrentTime: currentTime}
		jsonResponse(w, response, http.StatusOK)
		return
	}

	// Check if multiple timezones are requested
	timeZones := strings.Split(tzParam, ",")
	if len(timeZones) > 1 {
		handleMultipleTimezones(w, timeZones)
		return
	}

	// Handle single timezone
	loc, err := time.LoadLocation(tzParam)
	if err != nil {
		errorResponse(w, "invalid timezone", http.StatusNotFound)
		return
	}

	currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05 -0700 MST")
	response := TimeResponse{CurrentTime: currentTime}
	jsonResponse(w, response, http.StatusOK)
}

func handleMultipleTimezones(w http.ResponseWriter, timeZones []string) {
	response := make(MultiTimeResponse)

	for _, tz := range timeZones {
		loc, err := time.LoadLocation(strings.TrimSpace(tz))
		if err != nil {
			errorResponse(w, "invalid timezone", http.StatusNotFound)
			return
		}
		response[tz] = time.Now().In(loc).Format("2006-01-02 15:04:05 -0700 MST")
	}

	jsonResponse(w, response, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, message string, status int) {
	response := map[string]string{"error": message}
	jsonResponse(w, response, status)
}
