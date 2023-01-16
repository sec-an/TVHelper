package common

import (
	"encoding/json"

	"github.com/spf13/cast"
)

type Result struct {
	Page      int                 `json:"page,omitempty"`
	PageCount int                 `json:"pagecount,omitempty"`
	Limit     int                 `json:"limit,omitempty"`
	Total     int                 `json:"total,omitempty"`
	List      []Vod               `json:"list,omitempty"`
	Class     []Class             `json:"class,omitempty"`
	Filters   map[string][]Filter `json:"filters,omitempty"`
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
	_ = json.Unmarshal(data, &rr)
	r.Page = cast.ToInt(rr.Page)
	r.PageCount = cast.ToInt(rr.PageCount)
	r.Limit = cast.ToInt(rr.Limit)
	r.Total = cast.ToInt(rr.Total)
	return nil
}
