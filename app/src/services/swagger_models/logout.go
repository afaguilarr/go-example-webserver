package swagger_models

// swagger:model LogOutRequest
// LogOutRequest is the swagger / RESTful API model of the LogOut request body
type LogOutRequest struct {
	// Username of the user being logged out
	// in: string
	Username string `json:"username" validate:"required"`
}
