package examples

import (
	"context"
	"github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	"time"
)

type FooStorage struct {
}

var _ storage.Storage = &FooStorage{}

func (x *FooStorage) GetName() string {
	return ""
}

func (x *FooStorage) Init(ctx context.Context) error {
	return nil
}

func (x *FooStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *FooStorage) InsertWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *FooStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *FooStorage) Get(ctx context.Context, lockId string) (string, error) {
	return "", nil
}

func (x *FooStorage) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}

func (x *FooStorage) Close(ctx context.Context) error {
	return nil
}

func (x *FooStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	return nil, nil
}
