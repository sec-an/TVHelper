package common

type Sub struct {
	Url    string `json:"url,omitempty"`
	Name   string `json:"name,omitempty"`
	Lang   string `json:"lang,omitempty"`
	Format string `json:"format,omitempty"`
	Flag   *int   `json:"flag,omitempty"`
}
