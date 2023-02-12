package common

import (
	"TVHelper/global"
	"encoding/json"

	"github.com/spf13/cast"
	"go.uber.org/zap"
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
	Changeable int      `json:"changeable,omitempty"`
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
		Changeable interface{} `json:"changeable,omitempty"`
	}{
		TempSite: (*TempSite)(s),
	}
	if err := json.Unmarshal(data, &sr); err != nil {
		global.Logger.Error(string(data), zap.Error(err))
	}
	s.Type = cast.ToInt(sr.Type)
	s.PlayerType = cast.ToInt(sr.PlayerType)
	s.Searchable = cast.ToInt(sr.Searchable)
	s.Filterable = cast.ToInt(sr.Filterable)
	s.Changeable = cast.ToInt(sr.Changeable)
	return nil
}
