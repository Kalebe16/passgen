package main

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateOptionsAcceptsValidOptions(t *testing.T) {
	opts := Options{
		Length:    16,
		Uppercase: true,
	}

	err := validateOptions(opts)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestValidateOptionsSkipsValidationWhenVersionIsEnabled(t *testing.T) {
	opts := Options{
		Version: true,
		Length:  0,
	}

	err := validateOptions(opts)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestValidateOptionsRejectsLengthBelowMinimum(t *testing.T) {
	opts := Options{
		Length:    0,
		Uppercase: true,
	}

	err := validateOptions(opts)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestValidateOptionsRejectsLengthAboveMaximum(t *testing.T) {
	opts := Options{
		Length:    1001,
		Uppercase: true,
	}

	err := validateOptions(opts)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestValidateOptionsRejectsMissingCharsetOption(t *testing.T) {
	opts := Options{
		Length: 16,
	}

	err := validateOptions(opts)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestValidateOptionsRejectsEmptyOptions(t *testing.T) {
	opts := Options{}

	err := validateOptions(opts)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGeneratePasswordReturnsPasswordWithExpectedLength(t *testing.T) {
	password := generatePassword(32, true, false, false, false)

	if len(password) != 32 {
		t.Fatalf("expected length 32, got %d", len(password))
	}
}

func TestGeneratePasswordUsesOnlyUppercaseChars(t *testing.T) {
	password := generatePassword(64, true, false, false, false)

	for _, char := range password {
		if !strings.ContainsRune(uppercaseChars, char) {
			t.Fatalf("expected only uppercase chars, got %q in %q", char, password)
		}
	}
}

func TestGeneratePasswordUsesOnlyLowercaseChars(t *testing.T) {
	password := generatePassword(64, false, true, false, false)

	for _, char := range password {
		if !strings.ContainsRune(lowercaseChars, char) {
			t.Fatalf("expected only lowercase chars, got %q in %q", char, password)
		}
	}
}

func TestGeneratePasswordUsesOnlyNumberChars(t *testing.T) {
	password := generatePassword(64, false, false, true, false)

	for _, char := range password {
		if !strings.ContainsRune(numberChars, char) {
			t.Fatalf("expected only number chars, got %q in %q", char, password)
		}
	}
}

func TestGeneratePasswordUsesOnlySymbolChars(t *testing.T) {
	password := generatePassword(64, false, false, false, true)

	for _, char := range password {
		if !strings.ContainsRune(symbolChars, char) {
			t.Fatalf("expected only symbol chars, got %q in %q", char, password)
		}
	}
}

func TestGeneratePasswordUsesAllEnabledCharsetsAsAllowedChars(t *testing.T) {
	password := generatePassword(128, true, true, true, true)
	allowedChars := uppercaseChars + lowercaseChars + numberChars + symbolChars

	for _, char := range password {
		if !strings.ContainsRune(allowedChars, char) {
			t.Fatalf("unexpected char %q in %q", char, password)
		}
	}
}

func TestRunCopiesGeneratedPassword(t *testing.T) {
	opts := Options{
		Length:    16,
		Uppercase: true,
	}

	var copiedPassword string

	err := run(opts, func(password string) error {
		copiedPassword = password
		return nil
	})

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if len(copiedPassword) != 16 {
		t.Fatalf("expected copied password length 16, got %d", len(copiedPassword))
	}
}

func TestRunReturnsCopyError(t *testing.T) {
	expectedErr := errors.New("clipboard error")

	opts := Options{
		Length:    16,
		Uppercase: true,
	}

	err := run(opts, func(password string) error {
		return expectedErr
	})

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestRunDoesNotCopyPasswordWhenVersionIsEnabled(t *testing.T) {
	opts := Options{
		Version: true,
	}

	copyWasCalled := false

	err := run(opts, func(password string) error {
		copyWasCalled = true
		return nil
	})

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if copyWasCalled {
		t.Fatal("expected copy not to be called")
	}
}
