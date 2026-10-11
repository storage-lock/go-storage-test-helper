package storage_test_helper_test

import (
	"context"
	"testing"
	"time"

	iterator "github.com/golang-infrastructure/go-iterator"
	memory_storage "github.com/storage-lock/go-memory-storage"
	"github.com/storage-lock/go-storage"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
)

// newTestStorage 每个用例使用独立实例，避免 TestStorage_* 共用 TestLockId 相互干扰。
func newTestStorage() storage.Storage {
	return memory_storage.NewMemoryStorage()
}

func TestSmoke_TestStorage_Full(t *testing.T) {
	storage_test_helper.TestStorage(t, newTestStorage())
}

func TestSmoke_TestStorage_GetName(t *testing.T) {
	storage_test_helper.TestStorage_GetName(t, newTestStorage())
}

func TestSmoke_TestStorage_Capabilities(t *testing.T) {
	storage_test_helper.TestStorage_Capabilities(t, newTestStorage())
}

func TestSmoke_TestStorage_Init(t *testing.T) {
	storage_test_helper.TestStorage_Init(t, newTestStorage())
}

func TestSmoke_TestStorage_Get(t *testing.T) {
	storage_test_helper.TestStorage_Get(t, newTestStorage())
}

func TestSmoke_TestStorage_GetTime(t *testing.T) {
	storage_test_helper.TestStorage_GetTime(t, newTestStorage())
}

func TestSmoke_TestStorage_UpdateWithVersion(t *testing.T) {
	storage_test_helper.TestStorage_UpdateWithVersion(t, newTestStorage())
}

func TestSmoke_TestStorage_CreateWithVersion(t *testing.T) {
	storage_test_helper.TestStorage_CreateWithVersion(t, newTestStorage())
}

func TestSmoke_TestStorage_DeleteWithVersion(t *testing.T) {
	storage_test_helper.TestStorage_DeleteWithVersion(t, newTestStorage())
}

func TestSmoke_TestStorage_List(t *testing.T) {
	storage_test_helper.TestStorage_List(t, newTestStorage())
}

func TestSmoke_TestEnsureLockNotExists_DefaultLockId(t *testing.T) {
	s := newTestStorage()
	// 默认 lockId（无参数）：锁不存在时直接返回
	storage_test_helper.TestEnsureLockNotExists(t, s)
	// 插入一条后再调用，应将其删除
	storage_test_helper.TestStorage_CreateWithVersion(t, s)
	storage_test_helper.TestEnsureLockNotExists(t, s)
}

func TestSmoke_TestEnsureLockNotExists_CustomLockId(t *testing.T) {
	s := newTestStorage()
	storage_test_helper.TestEnsureLockNotExists(t, s, "custom-lock-id")
}

func TestSmoke_BuildTestLockInformation(t *testing.T) {
	info := storage_test_helper.BuildTestLockInformation(t)
	if info == nil {
		t.Fatal("BuildTestLockInformation 返回 nil")
	}
	if info.Version != storage.Version(storage_test_helper.TestLockVersion) {
		t.Fatalf("默认版本应为 %d，实际 %d", storage_test_helper.TestLockVersion, info.Version)
	}
	custom := storage_test_helper.BuildTestLockInformation(t, 42)
	if custom.Version != 42 {
		t.Fatalf("自定义版本应为 42，实际 %d", custom.Version)
	}
}

func TestSmoke_Constants(t *testing.T) {
	if storage_test_helper.TestLockId == "" {
		t.Fatal("TestLockId 不应为空")
	}
	if storage_test_helper.DefaultContextTimeout <= 0 {
		t.Fatal("DefaultContextTimeout 应为正数")
	}
}

type partialStorage struct{}

func (partialStorage) GetName() string { return "partial" }
func (partialStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{storage.CapabilityAtomicDelete}
}
func (partialStorage) Init(context.Context) error { return nil }
func (partialStorage) UpdateWithVersion(context.Context, string, storage.Version, storage.Version, *storage.LockInformation) error {
	return nil
}
func (partialStorage) CreateWithVersion(context.Context, string, storage.Version, *storage.LockInformation) error {
	return nil
}
func (partialStorage) DeleteWithVersion(context.Context, string, storage.Version, *storage.LockInformation) error {
	return nil
}
func (partialStorage) Get(context.Context, string) (string, error) { return "", nil }
func (partialStorage) GetTime(context.Context) (time.Time, error)  { return time.Now(), nil }
func (partialStorage) Close(context.Context) error                 { return nil }
func (partialStorage) List(context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	return nil, nil
}

func TestSmoke_MissingCapabilities(t *testing.T) {
	storage_test_helper.TestStorage_Capabilities(t, partialStorage{})
}
