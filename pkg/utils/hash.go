package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword plain text password ko secure hash mein convert karta hai
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash verify karta hai ke plain password hash se match karta hai ya nahi
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}