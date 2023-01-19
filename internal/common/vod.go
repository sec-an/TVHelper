package common

type Vod struct {
	VodId      string `json:"vod_id,omitempty"`
	VodName    string `json:"vod_name,omitempty"`
	VodPic     string `json:"vod_pic,omitempty"`
	VodRemarks string `json:"vod_remarks,omitempty"`
}
