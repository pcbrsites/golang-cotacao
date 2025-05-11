package http

type ResponseCotacao struct {
	Bid string `json:"bid"`
}

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}
