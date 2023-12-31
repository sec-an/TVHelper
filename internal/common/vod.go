package common

type Vod struct {
	VodId       string   `json:"vod_id,omitempty"`
	VodName     string   `json:"vod_name,omitempty"`
	TypeName    string   `json:"type_name,omitempty"`
	VodPic      string   `json:"vod_pic,omitempty"`
	VodRemarks  string   `json:"vod_remarks,omitempty"`
	VodYear     string   `json:"vod_year,omitempty"`
	VodArea     string   `json:"vod_area,omitempty"`
	VodDirector string   `json:"vod_director,omitempty"`
	VodActor    string   `json:"vod_actor,omitempty"`
	VodContent  string   `json:"vod_content,omitempty"`
	VodPlayFrom string   `json:"vod_play_from,omitempty"`
	VodPlayUrl  string   `json:"vod_play_url,omitempty"`
	VodTag      string   `json:"vod_tag,omitempty"`
	Cate        *Cate    `json:"cate,omitempty"`
	Style       *Style   `json:"style,omitempty"`
	Land        *int     `json:"land,omitempty"`
	Circle      *int     `json:"circle,omitempty"`
	Ratio       *float32 `json:"ratio,omitempty"`
}
