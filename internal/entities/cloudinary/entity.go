package cloudinary

type Response struct {
	//AssetId      string `json:"asset_id"`
	PublicId string `json:"public_id"`
	Version  int    `json:"version"`
	//Signature    string `json:"signature"`
	//Width        string `json:"width"`
	//Height       string `json:"height"`
	//Format       string `json:"format"`
	//ResourceType string `json:"resource_type"`
	//CreatedAt    string `json:"created_at"`
	//Pages        string `json:"pages"`
	//Bytes        int    `json:"bytes"`
	//Type         string `json:"String"`
	//Etag         string `json:"etag"`
	//PlaceHolder  string `json:"place_holder"`
	Url          string `json:"url"`
	SecureUrl    string `json:"secure_url"`
	PreviewUrl   string `json:"preview_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	//PlaybackUrl  string `json:"playback_url"`

	// i'll add more fields later on

}
