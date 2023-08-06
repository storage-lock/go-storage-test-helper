package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStorage_Close 用于测试Storage的Close实现是否正确
func TestStorage_Close(t *testing.T, storage storage.Storage) {

	assert.NotNilf(t, storage, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := storage.Init(ctx)
	assert.Nilf(t, err, "storage %s initialization error: %#v", storage.GetName(), err)

	err = storage.Close(ctx)
	assert.Nilf(t, err, "storage %s close error: %#v", storage.GetName(), err)
}
