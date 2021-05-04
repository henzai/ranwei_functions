package model

import "time"

type Payload struct {
	ContentType string `json:"content_type"`
	CreatedAt   int64  `json:"created_at"`
	FileName    string `json:"file_name"`
	ProxyURL    string `json:"proxy_url"`
	URL         string `json:"url"`
	ChannelID   string `json:"channel_id"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
}

type Item struct {
	ContentType string
	CreatedAt   time.Time
	FileName    string
	ProxyURL    string
	URL         string
	ChannelID   string
	UserID      string
	UserName    string
}

func PayloadToItem(p *Payload) *Item {
	t := time.Unix(p.CreatedAt, 0)
	return &Item{
		ContentType: p.ContentType,
		CreatedAt:   t,
		FileName:    p.FileName,
		ProxyURL:    p.ProxyURL,
		URL:         p.URL,
		ChannelID:   p.ChannelID,
		UserID:      p.UserID,
		UserName:    p.UserName,
	}
}
