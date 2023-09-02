package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_Init(t *testing.T, storage storage.Storage) {

	assert.NotNilf(t, storage, "Storage is nil")

	// 需要能够在五分钟内初始化成功
	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := storage.Init(ctx)

	// 初始化必须成功
	assert.Nilf(t, err, "Storage %s initialization error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	t.Log(fmt.Sprintf("Storage %s test Init done", storage.GetName()))
}
