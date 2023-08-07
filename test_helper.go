package storage_test_helper

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// ------------------------------------------------- --------------------------------------------------------------------

// TestStorage 用于测试Storage的实现是否OK
func TestStorage(t *testing.T, storage storage.Storage) {

	TestStorage_GetName(t, storage)

	TestStorage_Init(t, storage)

	TestStorage_Get(t, storage)

	TestStorage_GetTime(t, storage)

	TestStorage_UpdateWithVersion(t, storage)

	TestStorage_CreateWithVersion(t, storage)

	TestStorage_DeleteWithVersion(t, storage)

	TestStorage_Close(t, storage)
}

// ------------------------------------------------- --------------------------------------------------------------------

// TestEnsureLockNotExists 确保给定的锁在数据库中不存在，如果存在的话则将其删除
func TestEnsureLockNotExists(t *testing.T, s storage.Storage, lockId ...string) {

	if len(lockId) == 0 {
		lockId = append(lockId, TestLockId)
	}

	lockInformationJsonString, err := s.Get(context.Background(), lockId[0])
	if errors.Is(err, storage_lock.ErrLockNotFound) {
		return
	} else {
		assert.Nil(t, err)
	}

	information := &storage.LockInformation{}
	err = json.Unmarshal([]byte(lockInformationJsonString), &information)
	assert.Nil(t, err)
	err = s.DeleteWithVersion(context.Background(), lockId[0], information.Version, information)
	assert.Nil(t, err)
}

// BuildTestLockInformation 创建一个单元测试中使用的锁的信息
func BuildTestLockInformation(t *testing.T, version ...storage.Version) *storage.LockInformation {
	if len(version) == 0 {
		version = append(version, TestLockVersion)
	}
	information := &storage.LockInformation{
		OwnerId:         TestOwnerIdA,
		Version:         version[0],
		LockCount:       1,
		LockBeginTime:   time.Now(),
		LeaseExpireTime: time.Now().Add(time.Second * 30),
	}
	return information
}

// ------------------------------------------------- --------------------------------------------------------------------
