package model

type HealthResponse struct {
	Status string `json:"status"`
}

type PingResponse struct {
	Message string `json:"message"`
}
