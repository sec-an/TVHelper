package common

type Site struct {
	Key        string      `json:"key,omitempty"`
	Name       string      `json:"name,omitempty"`
	Api        string      `json:"api,omitempty"`
	Ext        interface{} `json:"ext,omitempty"`
	Jar        string      `json:"jar,omitempty"`
	Click      string      `json:"click,omitempty"`
	PlayUrl    string      `json:"playUrl,omitempty"`
	Type       *int        `json:"type,omitempty"`
	Timeout    *int        `json:"timeout,omitempty"`
	PlayerType *int        `json:"playerType,omitempty"`
	Searchable *int        `json:"searchable,omitempty"`
	Changeable *int        `json:"changeable,omitempty"`
	Recordable *int        `json:"recordable,omitempty"`
	Categories []string    `json:"categories,omitempty"`
	Header     interface{} `json:"header,omitempty"`
	Style      *Style      `json:"style,omitempty"`
	ViewType   interface{} `json:"viewType,omitempty"`
}

//func (s *Site) UnmarshalJSON(data []byte) error {
//	type TempSite Site
//	sr := struct {
//		*TempSite
//		Type       interface{} `json:"type,omitempty"`
//		Timeout    interface{} `json:"timeout,omitempty"`
//		PlayerType interface{} `json:"playerType,omitempty"`
//		Searchable interface{} `json:"searchable,omitempty"`
//		Changeable interface{} `json:"changeable,omitempty"`
//		Recordable interface{} `json:"recordable,omitempty"`
//	}{
//		TempSite: (*TempSite)(s),
//	}
//	if err := json.Unmarshal(data, &sr); err != nil {
//		global.Logger.Error(string(data), zap.Error(err))
//	}
//	s.Type = cast.ToInt(sr.Type)
//	s.Timeout = cast.ToInt(sr.Timeout)
//	s.PlayerType = cast.ToInt(sr.PlayerType)
//	s.Searchable = cast.ToInt(sr.Searchable)
//	s.Changeable = cast.ToInt(sr.Changeable)
//	s.Recordable = cast.ToInt(sr.Recordable)
//	return nil
//}
