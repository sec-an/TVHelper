package common

type Group struct {
	Channel []Channel `json:"channel,omitempty"`
	Logo    string    `json:"logo,omitempty"`
	Name    string    `json:"name,omitempty"`
	Pass    string    `json:"pass,omitempty"`
}
