package utils

import (
	"errors"
	"log"
	"regexp"
	"unicode"
)

type Email string
type Password string

func (e *Email) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(string(*e)) {
		return errors.New("invalid email")
	}
	return nil
}

func (p *Password) Validate() error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	log.Println(len(*p))
	if len(string(*p)) >= 7 {
		hasMinLen = true
	}
	for _, char := range *p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if !hasMinLen && !hasUpper && !hasLower && !hasNumber && !hasSpecial {
		return errors.New("invalid password: Should be at least 8 characters, at least 1 letter, at least 1 number and at least 1 special character")
	}
	return nil
}

func (p *Password) Hash() (Password, error) {
	bytes, err := HashPassword(string(*p))
	return Password(bytes), err
}

func (p *Password) Check(password string) bool {
	return CheckPasswordHash(string(*p), password)
}
