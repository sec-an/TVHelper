package common

type Core struct {
	Auth   string `json:"auth,omitempty"`
	Name   string `json:"name,omitempty"`
	Pass   string `json:"pass,omitempty"`
	Broker string `json:"broker,omitempty"`
}
