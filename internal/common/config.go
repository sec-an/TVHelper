package common

type Config struct {
	Ads       []string    `json:"ads,omitempty"`
	Flags     []string    `json:"flags,omitempty"`
	Parses    []Parse     `json:"parses,omitempty"`
	Sites     []Site      `json:"sites,omitempty"`
	Ijk       interface{} `json:"ijk,omitempty"`
	Lives     []Live      `json:"lives,omitempty"`
	Spider    string      `json:"spider,omitempty"`
	Wallpaper string      `json:"wallpaper,omitempty"`
}
