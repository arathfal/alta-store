package helper

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pass string) string {
	pwd := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

func CheckHashAndPass(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, err
	}
	return true, nil
}
