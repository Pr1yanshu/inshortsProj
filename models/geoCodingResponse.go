package models

type GeoCodingResponse struct {
	Data []data `json:"data"`
}

type data struct {
	AdministrativeArea interface{} `json:"administrative_area"`
	Confidence         float64     `json:"confidence"`
	Continent          string      `json:"continent"`
	Country            string      `json:"country"`
	CountryCode        string      `json:"country_code"`
	County             string      `json:"county"`
	Distance           float64     `json:"distance"`
	Label              string      `json:"label"`
	Latitude           float64     `json:"latitude"`
	Locality           string      `json:"locality"`
	Longitude          float64     `json:"longitude"`
	Name               string      `json:"name"`
	Neighbourhood      interface{} `json:"neighbourhood"`
	Number             string      `json:"number"`
	PostalCode         string      `json:"postal_code"`
	Region             string      `json:"region"`
	RegionCode         string      `json:"region_code"`
	Street             string      `json:"street"`
	Type               string      `json:"type"`
}
