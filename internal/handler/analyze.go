package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Localisation string

const (
	OnSite Localisation = "On-site"
	Remote Localisation = "Remote"
	Hybrid Localisation = "Hybrid"
)

type userData struct {
	Skills            []string       `json:"skills"`
	YearsOfExperience json.Number    `json:"yearsOfExperience"`
	Role              string         `json:"role"`
	SoftSkills        []string       `json:"softSkillsKeywords"`
	Localisation      []Localisation `json:"localisation"`
}

type Data struct {
	User userData `json:"userData"`
	Jd   string   `json:"jobDescription"`
}

func analyze(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Wrong content type", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			msg := fmt.Sprintf("No body included %s", err)
			logger.Error(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			msg := fmt.Sprintf("Parsing error: %s", err)
			logger.Error(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		logger.Info("userData", "role", data.User.Role, "Job description", data.Jd)
	}
}
