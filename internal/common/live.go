package common

type Live struct {
	Name       string      `json:"name,omitempty"`
	Type       *int        `json:"type"`
	Boot       *bool       `json:"boot,omitempty"`
	Pass       *bool       `json:"pass,omitempty"`
	Group      string      `json:"group,omitempty"`
	Url        string      `json:"url,omitempty"`
	Api        string      `json:"api,omitempty"`
	Ext        interface{} `json:"ext,omitempty"`
	Jar        string      `json:"jar,omitempty"`
	Logo       string      `json:"logo,omitempty"`
	Epg        string      `json:"epg,omitempty"`
	Ua         string      `json:"ua,omitempty"`
	Click      string      `json:"click,omitempty"`
	Origin     string      `json:"origin,omitempty"`
	Referer    string      `json:"referer,omitempty"`
	TimeZone   string      `json:"timeZone,omitempty"`
	Keep       string      `json:"keep,omitempty"`
	Catchup    *Catchup    `json:"catchup,omitempty"`
	Timeout    *int        `json:"timeout,omitempty"`
	Header     interface{} `json:"header,omitempty"`
	PlayerType *int        `json:"playerType,omitempty"`
	Channels   []Channel   `json:"channels,omitempty"`
	Groups     []Group     `json:"groups,omitempty"`
	Core       *Core       `json:"core,omitempty"`
}

//func (l *Live) UnmarshalJSON(data []byte) error {
//	type TempLive Live
//	lr := struct {
//		*TempLive
//		Type       interface{} `json:"type"`
//		PlayerType interface{} `json:"playerType,omitempty"`
//		Boot       interface{} `json:"boot,omitempty"`
//	}{
//		TempLive: (*TempLive)(l),
//	}
//	if err := json.Unmarshal(data, &lr); err != nil {
//		global.Logger.Error(string(data), zap.Error(err))
//	}
//	l.Type = cast.ToInt(lr.Type)
//	l.PlayerType = cast.ToInt(lr.PlayerType)
//	switch lr.Boot.(type) {
//	case bool:
//		l.Boot = lr.Boot.(bool)
//	case string:
//		boolValue, err := strconv.ParseBool(lr.Boot.(string))
//		if err != nil {
//			global.Logger.Error(lr.Boot.(string), zap.Error(err))
//		}
//		l.Boot = boolValue
//	default:
//		l.Boot = false
//		if cast.ToInt(lr.Boot) > 0 {
//			l.Boot = true
//		}
//	}
//	return nil
//}
