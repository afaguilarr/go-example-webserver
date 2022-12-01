package swagger_models

// swagger:model CommonSuccess
// CommonSuccess is the swagger / RESTful API model of the common succeses
type CommonSuccess struct {
	// Status of the success
	// in: int64
	Status int64 `json:"status"`
	// Message of the success
	// in: string
	Message string `json:"message"`
}

// swagger:model CommonError
// CommonError is the swagger / RESTful API model of the errors
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}
