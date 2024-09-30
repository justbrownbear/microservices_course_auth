package password_management

// Пример частично взял здесь:
// https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var hashParams = &params{
	memory:      19 * 1024,
	iterations:  2,
	parallelism: 1,
	saltLength:  32,
	keyLength:   32,
}

// PasswordManagementService defines the interface for password management operations.
// It includes methods for hashing passwords.
type PasswordManagementService interface {
	HashPassword(password string) (string, error)
}
