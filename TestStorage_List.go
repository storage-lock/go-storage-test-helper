package storage_test_helper

import (
	"context"
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_List(t *testing.T, s storage.Storage) {

	assert.NotNilf(t, s, "Storage is nil")

	// 确保锁不存在，避免等会儿插入的时候失败
	TestEnsureLockNotExists(t, s, TestLockId)

	// 先插入一条锁的信息
	ctx, cancelFunc := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc()
	err := s.CreateWithVersion(ctx, TestLockId, TestLockVersion, BuildTestLockInformation(t))
	assert.Nilf(t, err, "Storage %s create error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))

	// 然后列出所有的锁
	ctx2, cancelFunc2 := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancelFunc2()
	iterator, err := s.List(ctx2)
	assert.Nilf(t, err, "Storage %s List error: %s", s.GetName(), if_expression.ReturnByFunc(err != nil, func() string {
		return err.Error()
	}, func() string {
		return ""
	}))
	assert.NotNilf(t, iterator, "Storage %s List error, iterator is nil", s.GetName())
	if iterator != nil {
		slice := make([]*storage.LockInformation, 0)
		for iterator.Next() {
			slice = append(slice, iterator.Value())
		}
		assert.Truef(t, len(slice) != 0, "Storage %s list failed", s.GetName())
	}

	t.Log(fmt.Sprintf("Storage %s test List done", s.GetName()))
}
