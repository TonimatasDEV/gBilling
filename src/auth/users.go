package auth

import "golang.org/x/crypto/bcrypt"

var users = map[string]string{
	"admin@admin.com": "$2a$10$POnxwr26JnnEtcazSyniHOLqVARw1F4JvATKUtUocexooRwC0ITK2", // email: password
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(email, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(users[email]), []byte(password))
	return err == nil
}

func EmailExist(email string) bool {
	_, ok := users[email]
	return ok
}
