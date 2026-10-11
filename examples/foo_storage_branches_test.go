package examples

import (
	"context"
	"errors"
	"testing"

	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

func TestFooStorageConflictBranches(t *testing.T) {
	ctx := context.Background()
	store := NewFooStorage()
	info := &storage.LockInformation{LockId: "lock", OwnerId: "owner", Version: 1, LockCount: 1}
	if err := store.UpdateWithVersion(ctx, "missing", 1, 2, info); !errors.Is(err, storage_lock.ErrVersionMiss) {
		t.Fatal(err)
	}
	if err := store.DeleteWithVersion(ctx, "missing", 1, info); !errors.Is(err, storage_lock.ErrLockNotFound) {
		t.Fatal(err)
	}
	if err := store.CreateWithVersion(ctx, "lock", 1, info); err != nil {
		t.Fatal(err)
	}
	if err := store.CreateWithVersion(ctx, "lock", 1, info); !errors.Is(err, storage_lock.ErrLockAlreadyExists) {
		t.Fatal(err)
	}
	if err := store.UpdateWithVersion(ctx, "lock", 9, 2, info); !errors.Is(err, storage_lock.ErrVersionMiss) {
		t.Fatal(err)
	}
	if err := store.DeleteWithVersion(ctx, "lock", 9, info); !errors.Is(err, storage_lock.ErrVersionMiss) {
		t.Fatal(err)
	}
}
