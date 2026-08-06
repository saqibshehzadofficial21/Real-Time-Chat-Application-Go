package user

// registerRequest signup ke waqt client se aane wala JSON body
type registerRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}