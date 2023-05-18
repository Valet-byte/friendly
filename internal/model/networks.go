package model

type SocialNetworks struct {
	VkLink        string `json:"vk_link"`
	TelegramLink  string `json:"telegram_link"`
	InstagramLink string `json:"instagram_link"`
	TiktokLink    string `json:"tiktok_link"`
	YoutubeLink   string `json:"youtube_link"`
}

func (n *SocialNetworks) IsEmpty() bool {
	if n.VkLink != "" &&
		n.TelegramLink != "" &&
		n.InstagramLink != "" &&
		n.TiktokLink != "" &&
		n.YoutubeLink != "" {
		return false
	} else {
		return true
	}
}
