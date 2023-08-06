package examples

import (
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"testing"
)

func TestFooStorage(t *testing.T) {
	storage_test_helper.TestStorage(t, &FooStorage{})
}
