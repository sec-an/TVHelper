package common

type Group struct {
	Channel []Channel `json:"channel,omitempty"`
	Name    string    `json:"name,omitempty"`
	Pass    string    `json:"pass,omitempty"`
}
