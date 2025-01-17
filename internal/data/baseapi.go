package data

type AuthLoginResponse struct {
	Data struct {
		SID string `json:"sid"`
	} `json:"data"`
	Error   Error `json:"error"`
	Success bool  `json:"success"`
}

type GenericResponse struct {
	Error   Error `json:"error"`
	Success bool  `json:"success"`
}
