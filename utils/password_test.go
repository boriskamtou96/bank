package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "myPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}
	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "myPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}
	err = CheckPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Failed to check password: %v", err)
	}
}
