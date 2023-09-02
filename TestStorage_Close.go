package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStorage_Close 用于测试Storage的Close实现是否正确
func TestStorage_Close(t *testing.T, storage storage.Storage) {

	assert.NotNilf(t, storage, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := storage.Init(ctx)
	assert.Nilf(t, err, "Storage %s initialization error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	err = storage.Close(ctx2)
	assert.Nilf(t, err, "Storage %s close error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	t.Log(fmt.Sprintf("Storage %s test Close done", storage.GetName()))
}
