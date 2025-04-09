package domain

type UserInfo struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	Country      string `json:"country"`
	CountryState string `json:"countryState"`
	City         string `json:"city"`
	ZipCode      string `json:"zipCode"`
	Address      string `json:"address"`
	Lang         string `json:"lang"`
}
