package Helper

import "github.com/go-playground/validator/v10"

type respone struct {
	Meta meta `json:"meta"`
	// penggunaan interface untuk pola data bisa berubah
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIRespone(message string, code int, status string, data interface{}) respone {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	respone := respone{
		Meta: meta,
		Data: data,
	}

	return respone
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
