package swagger_models

// swagger:model RegisterRequest
// RegisterRequest is the swagger / RESTful API model of the Register request
type RegisterRequest struct {
	// User information of the user being registered
	// in: User
	User User `json:"user" validate:"required"`
	// Password of the user being registered
	// in: string
	Password string `json:"password" validate:"required"`
}

// swagger:model RegisterResponse
// RegisterResponse is the swagger / RESTful API model of the Register response
type RegisterResponse struct {
	// User information of the user being registered
	// in: User
	User User `json:"user" validate:"required"`
}
