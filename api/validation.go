package api

import (
	"strconv"
	"time"
)

type StatsGetRequest struct {
	Errors map[string]string
}

// ValidateContentGetRequest - validate get request parameters, teamId and date
func ValidateContentGetRequest(teamIdString string, dateString string) (int, *StatsGetRequest, bool) {
	msg := &StatsGetRequest{}
	msg.Errors = make(map[string]string)
	teamId, err := strconv.Atoi(teamIdString)

	if err != nil {
		msg.Errors["teamId"] = "Please enter a valid integer for the team-id parameter"
		return 0, msg, true
	}

	_, dateErr := time.Parse("2006-01-02", dateString)
	if dateErr != nil {
		msg.Errors["date"] = "Please enter a valid date for the date parameter"
		return 0, msg, true
	}

	return teamId, msg, false
}
