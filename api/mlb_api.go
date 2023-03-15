package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const statsAPIBaseUrl = "https://statsapi.mlb.com/api"

// GetMLBEventsByDate - Get a list of MLB games by date
func (c *client) GetMLBEventsByDate(date string) GameStats {
	resp, err := c.SendMLBStatsRequest(fmt.Sprintf("%v/v1/schedule?date=%s&sportId=1&language=en", statsAPIBaseUrl, date))

	if err != nil {
		c.logger.Println("no response from request")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result GameStats

	if err := json.Unmarshal(body, &result); err != nil {
		c.logger.Fatalln("GetMLBEventsByDate: cannot unmarshal JSON")
	}
	return result
}

// SendMLBStatsRequest - send the request
func (c *client) SendMLBStatsRequest(url string) (*http.Response, error) {
	c.logger.Println("request sent to mlb statsAPI")
	return http.Get(url)
}
