package models

// SuccessMessage representa la respuesta de éxito
type SuccessMessage struct {
	Message string `json:"message"`
}

// ErrorMessage representa la respuesta de error
type ErrorMessage struct {
	Error string `json:"error"`
}
