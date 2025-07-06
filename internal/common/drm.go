package common

type Drm struct {
	Key      string      `json:"key,omitempty"`
	Type     string      `json:"type,omitempty"`
	ForceKey bool        `json:"forceKey,omitempty"`
	Header   interface{} `json:"header,omitempty"`
}
