package dto

import "time"

type RoadMapDTO struct {
	Name         string    `json:"name"`
	Start        time.Time `json:"start_time"`
	ExpectedTime time.Time `json:"expected_end_time"`
	Active       bool      `json:"active"`
	LeadBy       string    `json:"lead_by"`
}
