package model

type Link struct {
	Rel      string `json:"rel"`
	Href     string `json:"href"`
	Type     string `json:"type"`
	Template string `json:"template"`
}
