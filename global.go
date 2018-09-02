package twitterdownloader

const (
	URL_PREFIX = "https://twitter.com/i/videos/tweet/"
)

type VideoConfig struct {
	Track Track `json:"track"`
}

type Track struct {
	ContentId    string `json:"contentId"`
	ContentType  string `json:"contentType"`
	Duration     int    `json:"durationMs"`
	ExpandedUrl  string `json:"expandedUrl"`
	PlaybackType string `json:"playbackType"`
	PlaybackUrl  string `json:"playbackUrl"`
}

type GuestToken struct {
	GuestToken string `json:"guest_token"`
}
