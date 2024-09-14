package password_management

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// HashPassword takes a plain text password and returns a hashed password using the Argon2id algorithm.
// It generates a random salt and combines it with the password to create a secure hash.
// The resulting hash is encoded in a specific format that includes the algorithm version, memory cost,
// iterations, parallelism, salt, and hash.
//
// Parameters:
//   - password: The plain text password to be hashed.
//
// Returns:
//   - A string containing the encoded hash in the format: $argon2id$v=Version$m=Memory,t=Iterations,p=Parallelism$Salt$Hash
//   - An error if any issue occurs during the hashing process.
func HashPassword(password string) (string, error) {
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
