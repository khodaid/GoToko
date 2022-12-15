package Helper

type respone struct {
	Meta meta
	// penggunaan interface untuk pola data bisa berubah
	Data interface{}
}

type meta struct {
	Message string
	Code    int
	Status  string
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
