package hash_password

import "golang.org/x/crypto/bcrypt"

func Hash(s string) []byte {
	bsp, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return bsp
}
