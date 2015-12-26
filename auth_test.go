package main

import "testing"

func TestToken(t *testing.T) {
	string1, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	if string1 == "" {
		t.Error("Empty token!")
	}

	string2, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	if string2 == "" {
		t.Error("Empty token!")
	}

	if string1 == string2 {
		t.Error("Tokens should be different!")
	}
}
