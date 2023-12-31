package common

type Channel struct {
	Urls       []string    `json:"urls,omitempty"`
	Number     string      `json:"number,omitempty"`
	Logo       string      `json:"logo,omitempty"`
	Epg        string      `json:"epg,omitempty"`
	Name       string      `json:"name,omitempty"`
	Ua         string      `json:"ua,omitempty"`
	Click      string      `json:"click,omitempty"`
	Referer    string      `json:"referer,omitempty"`
	Header     interface{} `json:"header,omitempty"`
	PlayerType *int        `json:"playerType,omitempty"`
	Parse      *int        `json:"parse,omitempty"`
	Drm        *Drm        `json:"drm,omitempty"`
}
