package httpresponse

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Status  int         `json:"status"`
}
