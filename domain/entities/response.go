package entities

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}
