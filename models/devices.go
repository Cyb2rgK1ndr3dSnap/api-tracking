package models

type RegisterToken struct {
	IDUser int    `json:"id_user,omitempty"`
	Token  string `json:"token,omitempty"`
}
