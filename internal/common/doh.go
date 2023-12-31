package common

type Doh struct {
	Name string   `json:"name,omitempty"`
	Url  string   `json:"url,omitempty"`
	Ips  []string `json:"ips,omitempty"`
}
