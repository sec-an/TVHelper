package common

type Rule struct {
	Name    string   `json:"name,omitempty"`
	Hosts   []string `json:"hosts,omitempty"`
	Regex   []string `json:"regex,omitempty"`
	Script  []string `json:"script,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}
