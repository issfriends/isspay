package crypto

import (
	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"
)

// Password bcrypt-encrypted password
type Password string

// Verify verify if plain password is correct
func (pwd Password) Verify(plainPwd string) error {
	rawHash, err := Base64Decode(string(pwd))
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(rawHash, []byte(plainPwd))
}

// Value implement sql Valuer interface
func (pwd Password) Value() (driver.Value, error) {
	rawHash, err := pwd.encrypted()
	if err != nil {
		return nil, err
	}
	return Base64Encode(rawHash), nil
}

func (pwd Password) encrypted() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}
