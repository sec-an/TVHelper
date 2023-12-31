package common

type Parse struct {
	Name string `json:"name,omitempty"`
	Type *int   `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
	Ext  *Ext   `json:"ext,omitempty"`
}

type Ext struct {
	Flag   []string    `json:"flag,omitempty"`
	Header interface{} `json:"header,omitempty"`
}

//func (p *Parse) UnmarshalJSON(data []byte) error {
//	type TempParse Parse
//	pr := struct {
//		*TempParse
//		Type interface{} `json:"type,omitempty"`
//	}{
//		TempParse: (*TempParse)(p),
//	}
//	if err := json.Unmarshal(data, &pr); err != nil {
//		global.Logger.Error(string(data), zap.Error(err))
//	}
//	p.Type = cast.ToInt(pr.Type)
//	return nil
//}
