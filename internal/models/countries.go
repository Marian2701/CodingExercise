package models

type Countries string

const (
	NotACountry           = ""
	Argentina   Countries = "Argentina"
	Australia   Countries = "Australia"
	Brazil      Countries = "Brazil"
	Canada      Countries = "Canada"
	China       Countries = "China"
	Denmark     Countries = "Denmark"
	Egypt       Countries = "Egypt"
	France      Countries = "France"
	Germany     Countries = "Germany"
	India       Countries = "India"
	Indonesia   Countries = "Indonesia"
	Italy       Countries = "Italy"
	Japan       Countries = "Japan"
	Morocco     Countries = "Morocco"
	Nigeria     Countries = "Nigeria"
	Poland      Countries = "Poland"
	SouthAfrica Countries = "South Africa"
	Spain       Countries = "Spain"
	UK          Countries = "UK"
	USA         Countries = "USA"
)

func GetCountryFromString(countryName string) Countries {
	for _, country := range AllCountries {
		if string(country) == countryName {
			return country
		}
	}
	return NotACountry
}

var AllCountries = []Countries{
	Argentina,
	Australia,
	Brazil,
	Canada,
	China,
	Denmark,
	Egypt,
	France,
	Germany,
	India,
	Indonesia,
	Italy,
	Japan,
	Morocco,
	Nigeria,
	Poland,
	SouthAfrica,
	Spain,
	UK,
	USA,
}

func (c Countries) String() string {
	return string(c)
}
