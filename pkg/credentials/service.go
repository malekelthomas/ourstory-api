package credentials

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/pbkdf2"
)

const TOKEN_LEN = 24

func GenerateToken(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateSecurePassword(salt, plainText string) string {
	iter := 100
	keyLen := 256
	k := pbkdf2.Key([]byte(plainText), []byte(salt), iter, keyLen, sha256.New)
	return string(k)
}

func Verify(salt, attempedPassword, securePassword string) bool {
	return GenerateSecurePassword(salt, attempedPassword) == securePassword
}
