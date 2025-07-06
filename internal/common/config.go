package common

type Config struct {
	Doh       []Doh    `json:"doh,omitempty"`
	Rules     []Rule   `json:"rules,omitempty"`
	Sites     []Site   `json:"sites,omitempty"`
	Parses    []Parse  `json:"parses,omitempty"`
	Flags     []string `json:"flags,omitempty"`
	Hosts     []string `json:"hosts,omitempty"`
	Ads       []string `json:"ads,omitempty"`
	Lives     []Live   `json:"lives,omitempty"`
	Spider    string   `json:"spider,omitempty"`
	Logo      string   `json:"logo,omitempty"`
	Wallpaper string   `json:"wallpaper,omitempty"`
}
