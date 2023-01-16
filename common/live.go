package common

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/spf13/cast"
)

type Live struct {
	Type       int       `json:"type,omitempty"`
	Boot       bool      `json:"boot,omitempty"`
	Name       string    `json:"name,omitempty"`
	Group      string    `json:"group,omitempty"`
	Url        string    `json:"url,omitempty"`
	Logo       string    `json:"logo,omitempty"`
	Epg        string    `json:"epg,omitempty"`
	Ua         string    `json:"ua,omitempty"`
	PlayerType int       `json:"playerType,omitempty"`
	Channels   []Channel `json:"channels,omitempty"`
	Groups     []Group   `json:"groups,omitempty"`
	Core       Core      `json:"core,omitempty"`
}

func (l *Live) UnmarshalJSON(data []byte) error {
	type TempLive Live
	lr := struct {
		*TempLive
		Type       interface{} `json:"type,omitempty"`
		PlayerType interface{} `json:"playerType,omitempty"`
		Boot       interface{} `json:"boot,omitempty"`
	}{
		TempLive: (*TempLive)(l),
	}
	_ = json.Unmarshal(data, &lr)
	l.Type = cast.ToInt(lr.Type)
	l.PlayerType = cast.ToInt(lr.PlayerType)
	switch lr.Boot.(type) {
	case bool:
		l.Boot = lr.Boot.(bool)
	case string:
		boolValue, err := strconv.ParseBool(lr.Boot.(string))
		if err != nil {
			log.Fatal(err)
		}
		l.Boot = boolValue
	default:
		l.Boot = false
		if cast.ToInt(lr.Boot) > 0 {
			l.Boot = true
		}
	}
	return nil
}
