package common

type Core struct {
	Auth   string `json:"auth,omitempty"`
	Name   string `json:"name,omitempty"`
	Pass   string `json:"pass,omitempty"`
	Broker string `json:"broker,omitempty"`
	Domain string `json:"domain,omitempty"`
	Resp   string `json:"resp,omitempty"`
	Sign   string `json:"sign,omitempty"`
	Pkg    string `json:"pkg,omitempty"`
	So     string `json:"so,omitempty"`
}
