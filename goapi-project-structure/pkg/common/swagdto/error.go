package swagdto

type ErrorDetail struct {
	// Code    string `json:"code" example:"Required"`
	Target  string `json:"target" example:"Name"`
	Message string `json:"message" example:"Name field is required"`
}

type ErrorData struct {
	Code    string        `json:"code" example:"BAD_REQUEST"`
	Message string        `json:"message" example:"Bad Request"`
	Details []ErrorDetail `json:"details"`
}

type Error400 struct {
	Status    uint      `json:"status" example:"400"`
	Error     ErrorData `json:"error"`
	RequestId string    `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
type Error401 struct {
	Status    uint      `json:"status" example:"401"`
	Error     ErrorData `json:"error"`
	RequestId string    `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
type Error403 struct {
	Status    uint      `json:"status" example:"403"`
	Error     ErrorData `json:"error"`
	RequestId string    `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
type Error404 struct {
	Status    uint      `json:"status" example:"404"`
	Error     ErrorData `json:"error"`
	RequestId string    `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
type Error500 struct {
	Status    uint      `json:"status" example:"500"`
	Error     ErrorData `json:"error"`
	RequestId string    `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
