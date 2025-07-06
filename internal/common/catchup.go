package common

type Catchup struct {
	Type    string `json:"type,omitempty"`
	Days    string `json:"days,omitempty"`
	Regex   string `json:"regex,omitempty"`
	Source  string `json:"source,omitempty"`
	Replace string `json:"replace,omitempty"`
}
