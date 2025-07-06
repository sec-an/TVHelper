package common

type Channel struct {
	Urls    []string    `json:"urls,omitempty"`
	Number  string      `json:"number,omitempty"`
	Logo    string      `json:"logo,omitempty"`
	Epg     string      `json:"epg,omitempty"`
	Name    string      `json:"name,omitempty"`
	Ua      string      `json:"ua,omitempty"`
	Click   string      `json:"click,omitempty"`
	Format  string      `json:"format,omitempty"`
	Origin  string      `json:"origin,omitempty"`
	Referer string      `json:"referer,omitempty"`
	TvgId   string      `json:"tvgId,omitempty"`
	TvgName string      `json:"tvgName,omitempty"`
	Catchup *Catchup    `json:"catchup,omitempty"`
	Header  interface{} `json:"header,omitempty"`
	Parse   *int        `json:"parse,omitempty"`
	Drm     *Drm        `json:"drm,omitempty"`
}
