package swagger_models

// swagger:model LogInRequest
// LogInRequest is the swagger / RESTful API model of the LogIn request
type LogInRequest struct {
	// Username of the user logging in
	// in: string
	Username string `json:"username" validate:"required"`
	// Password of the user logging in
	// in: string
	Password string `json:"password" validate:"required"`
}

// swagger:model LogInResponse
// LogInResponse is the swagger / RESTful API model of the LogIn response
type LogInResponse struct {
	// Status of the success
	// in: int64
	Status int64 `json:"status"`
	// AccessToken for the user that logged in
	// in: string
	AccessToken string `json:"accessToken" validate:"required"`
	// RefreshToken for the user that logged in
	// in: string
	RefreshToken string `json:"refreshToken" validate:"required"`
}
