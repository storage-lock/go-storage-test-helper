package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStorage_Init(t *testing.T, storage storage.Storage) {
	assert.NotNilf(t, storage, "storage is nil")

	// 需要能够在一分钟内初始化成功
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute)
	defer cancelFunc()
	err := storage.Init(ctx)

	// 初始化必须成功
	assert.Nilf(t, err, "storage %s initialization error: %#v", storage.GetName(), err)
}
