package common

import (
	"TVHelper/global"
	"encoding/json"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type Result struct {
	Page      int                 `json:"page,omitempty"`
	PageCount int                 `json:"pagecount,omitempty"`
	Limit     int                 `json:"limit,omitempty"`
	Total     int                 `json:"total,omitempty"`
	Parse     *int                `json:"parse,omitempty"`
	Code      *int                `json:"code,omitempty"`
	Jx        *int                `json:"jx,omitempty"`
	Class     []Class             `json:"class,omitempty"`
	List      []Vod               `json:"list,omitempty"`
	Filters   map[string][]Filter `json:"filters,omitempty"`
	Url       interface{}         `json:"url,omitempty"`
	Msg       interface{}         `json:"msg,omitempty"`
	Header    interface{}         `json:"header,omitempty"`
	PlayUrl   string              `json:"playUrl,omitempty"`
	JxFrom    string              `json:"jxFrom,omitempty"`
	Flag      string              `json:"flag,omitempty"`
	Danmaku   string              `json:"danmaku,omitempty"`
	Format    string              `json:"format,omitempty"`
	Click     string              `json:"click,omitempty"`
	Key       string              `json:"key,omitempty"`
	Subs      []Sub               `json:"subs,omitempty"`
}

func (r *Result) UnmarshalJSON(data []byte) error {
	type TempResult Result
	rr := struct {
		*TempResult
		Page      interface{} `json:"page,omitempty"`
		PageCount interface{} `json:"pagecount,omitempty"`
		Limit     interface{} `json:"limit,omitempty"`
		Total     interface{} `json:"total,omitempty"`
	}{
		TempResult: (*TempResult)(r),
	}
	if err := json.Unmarshal(data, &rr); err != nil {
		global.Logger.Error(string(data), zap.Error(err))
	}
	r.Page = cast.ToInt(rr.Page)
	r.PageCount = cast.ToInt(rr.PageCount)
	r.Limit = cast.ToInt(rr.Limit)
	r.Total = cast.ToInt(rr.Total)
	return nil
}
