package models

type Status struct {
	IDStatus int    `json:"id_status,omitempty"`
	Name     string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
}
