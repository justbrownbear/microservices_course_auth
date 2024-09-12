package password_management

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)


func HashPassword(password string) ( string, error ) {
	salt := make([]byte, hashParams.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, hashParams.iterations, hashParams.memory, hashParams.parallelism, hashParams.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, hashParams.memory, hashParams.iterations, hashParams.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}
