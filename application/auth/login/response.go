package login

type LoginResponse struct {
	ValidationErrors validationErrors `json:"errors,omitempty"`
	AccessToken      string           `json:"access_token,omitempty"`
	RefreshToken     string           `json:"refresh_token,omitempty"`
}
