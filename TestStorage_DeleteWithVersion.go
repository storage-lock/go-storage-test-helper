package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_DeleteWithVersion(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "Storage is nil")

	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.Init(ctx)
	assert.Nilf(t, err, "Storage %s init error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	TestEnsureLockNotExists(t, s)

	// 先插入一条
	lockInformation := BuildTestLockInformation(t)
	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	err = s.CreateWithVersion(ctx2, TestLockId, TestLockVersion, lockInformation)
	assert.Nil(t, err)

	// 确认能够查询得到
	ctx3, cancelFunc3 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc3()
	lockInformationJsonString, err := s.Get(ctx3, TestLockId)
	assert.Nil(t, err)
	assert.NotEmpty(t, lockInformationJsonString)

	// 再尝试将这一条删除
	ctx4, cancelFunc4 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc4()
	err = s.DeleteWithVersion(ctx4, TestLockId, TestLockVersion, lockInformation)
	assert.Nil(t, err)

	// 然后再查询，应该就查不到了
	ctx5, cancelFunc5 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc5()
	lockInformationJsonString, err = s.Get(ctx5, TestLockId)
	assert.ErrorIs(t, err, storage_lock.ErrLockNotFound)
	assert.Empty(t, lockInformationJsonString)

	t.Log(fmt.Sprintf("Storage %s test DeleteWithVersion done", s.GetName()))

}
