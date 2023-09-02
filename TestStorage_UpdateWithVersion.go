package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	storage_pkg "github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_UpdateWithVersion(t *testing.T, storage storage_pkg.Storage) {

	assert.NotNilf(t, storage, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := storage.Init(ctx)
	assert.Nilf(t, err, "Storage %s initialization error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	TestEnsureLockNotExists(t, storage, TestLockId)

	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	err = storage.CreateWithVersion(ctx2, TestLockId, TestLockVersion, BuildTestLockInformation(t))
	assert.Nilf(t, err, "Storage %s insert error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	newVersion := storage_pkg.Version(TestLockVersion + 1)
	ctx3, cancelFunc3 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc3()
	err = storage.UpdateWithVersion(ctx3, TestLockId, TestLockVersion, newVersion, BuildTestLockInformation(t, newVersion))
	assert.Nilf(t, err, "Storage %s update error: %s", storage.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	t.Log(fmt.Sprintf("Storage %s test UpdateWithVersion done", storage.GetName()))
}
