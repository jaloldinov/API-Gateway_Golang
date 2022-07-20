package helper

import (
	"errors"
	"math/rand"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

const POOL = "abcdefghijklmnopqrstuwxvyzABcDEFGHIJKLMNOPQRSTUYVWXYZ"

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 5 || len(password) > 30 {
		return errors.New("password length should be 8 to 30 characters")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("^[A-Za-z0-9$_@.#]+$"))) != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[0-9]"))) != nil {
		return errors.New("password should contain at least one number")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[A-Za-z]"))) != nil {
		return errors.New("password should contain at least one alphabetic character")
	}
	return nil
}

func ValidateLogin(login string) error {
	if login == "" {
		return errors.New("login cannot be blank")
	}
	if len(login) < 5 || len(login) > 15 {
		return errors.New("login length should be 5 to 15 characters")
	}
	if validation.Validate(login, validation.Match(regexp.MustCompile("^[A-Za-z0-9$@_.#]+$"))) != nil {
		return errors.New("login should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}

func ValidateUserType(userType string) error {
	if userType == "" {
		return errors.New("user-type cannot be blank")
	}
	return nil
}

func ValidateDate(date string) error {
	if date == "" {
		return errors.New("date is blank")
	}

	if validation.Validate(date, validation.Date("02-01-2006")) != nil {
		return errors.New("date must be DD-MM-YYYY format")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return errors.New("phone_number is blank")
	}

	if validation.Validate(phoneNumber, validation.Match(regexp.MustCompile("998(75|90|91|93|94|97|99)[0-9]{7}$"))) != nil {
		return errors.New("phone_number must be 998(XX)XXXXXXX")
	}
	return nil
}

func ValidateIp(ip string) error {
	if validation.Validate(ip, is.IPv4) != nil {
		return errors.New("ip must be in IPv4 Form")
	}

	return nil
}

func ValidatePort(port string) error {
	return validation.Validate(port, validation.Required, validation.Length(1, 5), is.Digit)
}

func ValidateOrderNo(orderNo int32) error {
	if orderNo < 0 {
		return errors.New("Order Number should be positive")
	}

	return nil
}

func GenerateRandomString(n int) string {
	l := byte(len(POOL))

	b, err := GenerateRandomBytes(n)
	if err != nil {
		return ""
	}

	for i := 0; i < n; i++ {
		b[i] = POOL[(b[i])%l]
	}

	return string(b)
}

// GenerateRandomBytes returns securely generated random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
