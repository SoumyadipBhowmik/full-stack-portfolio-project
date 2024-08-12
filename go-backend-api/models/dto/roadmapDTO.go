package dto

type RoadMapDTO struct {
	Name         string `json:"name"`
	Start        string `json:"start"`
	ExpectedTime string `json:"expected_time"`
	Active       bool   `json:"active"`
	LeadBy       string `json:"lead_by"`
}
