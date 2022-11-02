package dao

type Location struct {
	Country            string
	StateOrProvince    *string
	CityOrMunicipality *string
	Neighborhood       *string
	ZipCode            *string
}
