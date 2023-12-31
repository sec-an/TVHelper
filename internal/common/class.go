package common

type Class struct {
	TypeId   string   `json:"type_id,omitempty"`
	TypeName string   `json:"type_name,omitempty"`
	TypeFlag string   `json:"type_flag,omitempty"`
	Filters  []Filter `json:"filters,omitempty"`
	Land     *int     `json:"land,omitempty"`
	Circle   *int     `json:"circle,omitempty"`
	Ratio    *float32 `json:"ratio,omitempty"`
}
