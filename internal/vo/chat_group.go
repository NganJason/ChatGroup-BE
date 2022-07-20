package vo

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	DebugMsg *string `json:"debug_msg"`
	Message  *string `json:"message"`
}
