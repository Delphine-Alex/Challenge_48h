package models

type Event struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
	UUID_organizator string `json:"uuid_organizator"`
}

type CreateEvent struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
	UUID_organizator string `json:"uuid_organizator"`
}

type PostEvent struct {
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Location      string `json:"location"`
	Start         string `json:"start"`
	End           string `json:"end"`
	UUID_organizator string `json:"uuid_organizator"`
}