package auth

// loginRequest login ke waqt client se aane wala JSON body
type loginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}