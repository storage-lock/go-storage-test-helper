package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// FooStorage 演示如何实现 storage.Storage 并接入共享测试套件。
// 此前这里是一个空壳实现（GetName 返回空、Get 永远返回空字符串、Create/Update/Delete
// 都是空操作），导致自带的 TestFooStorage 全程失败——示例代码自己跑不通，起不到演示作用。
// 现在改为一个最小可用的内存实现（与 go-memory-storage 的核心语义一致：mutex 内
// check+write、版本号校验、UpdateWithVersion 记录不存在时返回 ErrVersionMiss），
// 让 TestFooStorage 真正通过，读者可以直接照着这个骨架替换成自己的存储后端。
type FooStorage struct {
	mu   sync.RWMutex
	data map[string]*storedRecord
}

type storedRecord struct {
	version storage.Version
	json    string
}

var _ storage.Storage = &FooStorage{}

func NewFooStorage() *FooStorage {
	return &FooStorage{
		data: make(map[string]*storedRecord),
	}
}

func (x *FooStorage) GetName() string {
	return "foo-storage"
}

func (x *FooStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{
		storage.CapabilityCAS,
		storage.CapabilityReliableTime,
	}
}

func (x *FooStorage) Init(ctx context.Context) error {
	return nil
}

func (x *FooStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	record, ok := x.data[lockId]
	if !ok {
		// 记录不存在与版本不匹配对调用方是同一回事：CAS 前提不成立，可重试。
		return fmt.Errorf("%w: lock %s not found", storage_lock.ErrVersionMiss, lockId)
	}
	if record.version != exceptedVersion {
		return fmt.Errorf("%w: excepted %d but got %d", storage_lock.ErrVersionMiss, exceptedVersion, record.version)
	}
	lockInformation.Version = newVersion
	jsonBytes, _ := json.Marshal(lockInformation)
	x.data[lockId] = &storedRecord{version: newVersion, json: string(jsonBytes)}
	return nil
}

func (x *FooStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	if _, ok := x.data[lockId]; ok {
		return fmt.Errorf("%w: lock %s already exists", storage_lock.ErrLockAlreadyExists, lockId)
	}
	lockInformation.Version = version
	jsonBytes, _ := json.Marshal(lockInformation)
	x.data[lockId] = &storedRecord{version: version, json: string(jsonBytes)}
	return nil
}

func (x *FooStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	record, ok := x.data[lockId]
	if !ok {
		return storage_lock.ErrLockNotFound
	}
	if record.version != exceptedVersion {
		return storage_lock.ErrVersionMiss
	}
	delete(x.data, lockId)
	return nil
}

func (x *FooStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	record, ok := x.data[lockId]
	if !ok {
		return "", storage_lock.ErrLockNotFound
	}
	return record.json, nil
}

func (x *FooStorage) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}

func (x *FooStorage) Close(ctx context.Context) error {
	return nil
}

func (x *FooStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	infos := make([]*storage.LockInformation, 0, len(x.data))
	for _, record := range x.data {
		var info storage.LockInformation
		_ = json.Unmarshal([]byte(record.json), &info)
		infos = append(infos, &info)
	}
	return iterator.FromSlice(infos), nil
}
