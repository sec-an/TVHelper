package vod

type Vod struct {
	VodId         int    `json:"vod_id"`
	Version       int    `json:"version"`
	Episode       int    `json:"episode"`
	SpPath        string `json:"sp_path"`
	CmeMaterialId string `json:"cme_material_id,omitempty"`
	CmeSignature  string `json:"cme_signature,omitempty"`
	Size          uint64 `json:"size,omitempty"`
	Width         int    `json:"width,omitempty"`
	Height        int    `json:"height,omitempty"`
	Fps           int    `json:"fps,omitempty"`
}
