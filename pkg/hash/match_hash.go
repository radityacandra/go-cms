package hash

import "golang.org/x/crypto/bcrypt"

func MatchHash(plain, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)); err != nil {
		return false
	}

	return true
}
