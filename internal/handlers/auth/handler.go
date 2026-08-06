package auth


import authsvc "chat-app/internal/service/auth"

// AuthHandler HTTP requests ko AuthService tak forward karta hai
type AuthHandler struct {
    service authsvc.AuthService
}

// NewAuthHandler ek naya AuthHandler banata hai
func NewAuthHandler(s authsvc.AuthService) *AuthHandler {
    return &AuthHandler{service: s}
}