package entities

type VideoReq struct {
	Name string
	File string
}

type VideoRes struct {
	VideoURL          string `json:"video_url"`
	VideoPreviewURL   string `json:"video_preview_url"`
	VideoThumbnailURL string `json:"video_thumbnail_url"`
}
