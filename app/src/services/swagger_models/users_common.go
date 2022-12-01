package swagger_models

// swagger:model User
// User is the swagger / RESTful API model of the User
type User struct {
	// User information related to general public data
	// in: PetMaster
	PetMaster PetMaster `json:"petMaster" validate:"required"`
	// Username of the user
	// in: string
	Username string `json:"username" validate:"required"`
	// Description of the user profile
	// in: string
	Description *string `json:"description"`
}

// swagger:model PetMaster
// PetMaster is the swagger / RESTful API model of the PetMaster
type PetMaster struct {
	// Name of the user
	// in: string
	Name string `json:"name" validate:"required"`
	// Contact number of the user
	// in: string
	ContactNumber *string `json:"contactNumber"`
	// Location of the user
	// in: Location
	Location Location `json:"location" validate:"required"`
}

// swagger:model Location
// Location is the swagger / RESTful API model of the Location
type Location struct {
	// Country of the user
	// in: string
	Country string `json:"country" validate:"required"`
	// State or province of the user
	// in: string
	StateOrProvince *string `json:"stateOrProvince"`
	// City or municipality of the user
	// in: string
	CityOrMunicipality *string `json:"cityOrMunicipality"`
	// Neighborhood of the user
	// in: string
	Neighborhood *string `json:"neighborhood"`
	// ZipCode of the user
	// in: string
	ZipCode *string `json:"zipCode"`
}
