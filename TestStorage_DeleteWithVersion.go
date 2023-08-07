package storage_test_helper

import (
	"context"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_DeleteWithVersion(t *testing.T, s storage.Storage) {
	assert.NotNilf(t, s, "storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "storage %s init error: %#v", s.GetName(), err)

	TestEnsureLockNotExists(t, s)

	// 先插入一条
	lockInformation := BuildTestLockInformation(t)
	err = s.CreateWithVersion(ctx, TestLockId, TestLockVersion, lockInformation)
	assert.Nil(t, err)

	// 确认能够查询得到
	lockInformationJsonString, err := s.Get(ctx, TestLockId)
	assert.Nil(t, err)
	assert.NotEmpty(t, lockInformationJsonString)

	// 再尝试将这一条删除
	err = s.DeleteWithVersion(ctx, TestLockId, TestLockVersion, lockInformation)
	assert.Nil(t, err)

	// 然后再查询，应该就查不到了
	lockInformationJsonString, err = s.Get(ctx, TestLockId)
	assert.ErrorIs(t, err, storage_lock.ErrLockNotFound)
	assert.Empty(t, lockInformationJsonString)

}
