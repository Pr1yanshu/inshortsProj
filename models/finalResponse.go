package models

type Finalresponse struct {
	Confirmed   int64  `json:"confirmed"`
	Deceased    int64  `json:"deceased"`
	Recovered   int64  `json:"recovered"`
	LastUpdated string `json:"last_updated"`
}
