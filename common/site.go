package common

import (
	"encoding/json"

	"github.com/spf13/cast"
)

type Site struct {
	Key        string   `json:"key,omitempty"`
	Name       string   `json:"name,omitempty"`
	Type       int      `json:"type,omitempty"`
	Api        string   `json:"api,omitempty"`
	PlayUrl    string   `json:"playUrl,omitempty"`
	PlayerType int      `json:"playerType,omitempty"`
	Searchable int      `json:"searchable,omitempty"`
	Filterable int      `json:"filterable,omitempty"`
	Switchable int      `json:"switchable,omitempty"`
	Ext        string   `json:"ext,omitempty"`
	Jar        string   `json:"jar,omitempty"`
	Categories []string `json:"categories,omitempty"`
}

func (s *Site) UnmarshalJSON(data []byte) error {
	type TempSite Site
	sr := struct {
		*TempSite
		Type       interface{} `json:"type,omitempty"`
		PlayerType interface{} `json:"playerType,omitempty"`
		Searchable interface{} `json:"searchable,omitempty"`
		Filterable interface{} `json:"filterable,omitempty"`
		Switchable interface{} `json:"switchable,omitempty"`
	}{
		TempSite: (*TempSite)(s),
	}
	_ = json.Unmarshal(data, &sr)
	s.Type = cast.ToInt(sr.Type)
	s.PlayerType = cast.ToInt(sr.PlayerType)
	s.Searchable = cast.ToInt(sr.Searchable)
	s.Filterable = cast.ToInt(sr.Filterable)
	s.Switchable = cast.ToInt(sr.Switchable)
	return nil
}
