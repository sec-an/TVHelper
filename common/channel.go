package common

type Channel struct {
	Urls   []string `json:"urls,omitempty"`
	Number string   `json:"number,omitempty"`
	Logo   string   `json:"logo,omitempty"`
	Epg    string   `json:"epg,omitempty"`
	Name   string   `json:"name,omitempty"`
	Ua     string   `json:"ua,omitempty"`
}
