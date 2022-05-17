package models

type Event struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
	UUID_wishlist string `json:"uuid_wishlist"`
}

type CreateEvent struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
}

type PostEvent struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
}