package models

type UrlShortener_Payload struct{
	OriginalURL string `json:"original_url" form:"original_url"`
	ShortenedURL string `json:"shortened_url" form:"shortened_url"`
}
