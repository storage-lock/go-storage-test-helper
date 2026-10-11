package storage_test_helper

import (
	"errors"
	"testing"
)

func TestErrText(t *testing.T) {
	if errText(nil) != "" {
		t.Fatal("nil")
	}
	if errText(errors.New("boom")) != "boom" {
		t.Fatal("error")
	}
}
