package common

type Filter struct {
	Key   string  `json:"key,omitempty"`
	Name  string  `json:"name,omitempty"`
	Value []Value `json:"value,omitempty"`
}

type Value struct {
	N string `json:"n,omitempty"`
	V string `json:"v,omitempty"`
}
