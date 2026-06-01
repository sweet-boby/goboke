package password

import "golang.org/x/crypto/bcrypt"

// TODO: Implement password strength validation
func IsStrongPassword(password string) bool {
	// TODO: Validate password strength:
	// - At least 8 characters
	if len(password) < 8 {
		return false
	}
	// - Contains uppercase letter
	hasUpper := false
	// - Contains lowercase letter
	hasLower := false
	// - Contains number
	hasNum := false
	// - Contains special character
	hasSpecial := false
	for _, item := range password {
		switch {
		case 'a' <= item && item <= 'z':
			hasLower = true
		case 'A' <= item && item <= 'Z':
			hasUpper = true
		case '0' <= item && item <= '9':
			hasNum = true
		default:
			hasSpecial = true
		}
	}
	return hasLower && hasUpper && hasNum && hasSpecial
}

// TODO: Implement password hashing
func HashPassword(password string) (string, error) {
	// TODO: Use bcrypt to hash the password with cost 12
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// TODO: Implement password verification
func VerifyPassword(password, hash string) bool {
	// TODO: Use bcrypt to compare password with hash
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return error == nil
}
