package models

type RegionData struct {
	Meta struct {
		LastUpdated string `json:"last_updated" bson:"last_updated"`
	} `json:"meta" bson:"meta"`
	Total struct {
		Confirmed int64 `json:"confirmed" bson:"confirmed"`
		Deceased  int64 `json:"deceased" bson:"deceased"`
		Recovered int64 `json:"recovered" bson:"recovered"`
	} `json:"total" bson:"total"`
}
